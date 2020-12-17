/*
Copyright 2016 Citrix Systems, Inc, All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package netscaler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Idempotent flag can't be added for these resources
var idempotentInvalidResources = []string{"login", "logout", "reboot", "shutdown", "ping", "ping6", "traceroute", "traceroute6", "install", "appfwjsoncontenttype", "appfwxmlcontenttype"}

const (
	nsErrSessionExpired = 444
	nsErrAuthTimeout    = 1027
)

func contains(slice []string, val string) bool {
	for _, item := range slice {
		if strings.EqualFold(item, val) {
			return true
		}
	}
	return false
}

type responseHandlerFunc func(resp *http.Response) ([]byte, error)

func createResponseHandler(resp *http.Response) ([]byte, error) {
	switch resp.Status {
	case "201 Created", "200 OK":
		body, _ := ioutil.ReadAll(resp.Body)
		return body, nil
	case "409 Conflict":
		body, _ := ioutil.ReadAll(resp.Body)
		return body, errors.New("failed: " + resp.Status + " (" + string(body) + ")")

	case "207 Multi Status":
		//This happens in case of Bulk operations, which we do not support yet
		body, _ := ioutil.ReadAll(resp.Body)
		return body, nil
	case "400 Bad Request", "401 Unauthorized", "403 Forbidden",
		"404 Not Found", "405 Method Not Allowed", "406 Not Acceptable",
		"503 Service Unavailable", "599 Netscaler specific error":
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println("[INFO] go-nitro: error = " + string(body))
		return body, errors.New("failed: " + resp.Status + " (" + string(body) + ")")
	default:
		body, err := ioutil.ReadAll(resp.Body)
		return body, err

	}
}

func deleteResponseHandler(resp *http.Response) ([]byte, error) {
	switch resp.Status {
	case "200 OK", "404 Not Found":
		body, _ := ioutil.ReadAll(resp.Body)
		return body, nil

	case "400 Bad Request", "401 Unauthorized", "403 Forbidden",
		"405 Method Not Allowed", "406 Not Acceptable",
		"409 Conflict", "503 Service Unavailable", "599 Netscaler specific error":
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println("[INFO] go-nitro: delete: error = " + string(body))
		return body, errors.New("[INFO] delete failed: " + resp.Status + " (" + string(body) + ")")
	default:
		body, err := ioutil.ReadAll(resp.Body)
		return body, err

	}
}

func readResponseHandler(resp *http.Response) ([]byte, error) {
	switch resp.Status {
	case "200 OK":
		body, _ := ioutil.ReadAll(resp.Body)
		return body, nil
	case "404 Not Found":
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println("[DEBUG] go-nitro: read: 404 not found")
		return body, errors.New("go-nitro: read: 404 not found: ")
	case "400 Bad Request", "401 Unauthorized", "403 Forbidden",
		"405 Method Not Allowed", "406 Not Acceptable",
		"409 Conflict", "503 Service Unavailable", "599 Netscaler specific error":
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println("[INFO] go-nitro: read: error = " + string(body))
		return body, errors.New("[INFO] go-nitro: failed read: " + resp.Status + " (" + string(body) + ")")
	default:
		body, err := ioutil.ReadAll(resp.Body)
		log.Println("[INFO] go-nitro: read error = " + string(body))
		return body, err

	}
}

func (c *NitroClient) createHTTPRequest(method string, urlstr string, buff *bytes.Buffer) (*http.Request, error) {
	req, err := http.NewRequest(method, urlstr, buff)
	if err != nil {
		return nil, err
	}
	// Get resourceType from url
	u, err := url.Parse(urlstr)
	if err != nil {
		return nil, err
	}
	splitStrings := strings.Split(u.Path, "/")
	resourceType := splitStrings[len(splitStrings)-1]

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	if c.proxiedNs == "" {
		if c.IsLoggedIn() {
			req.Header.Set("Set-Cookie", "NITRO_AUTH_TOKEN="+c.getSessionid())
		} else {
			if resourceType != "login" {
				req.Header.Set("X-NITRO-USER", c.username)
				req.Header.Set("X-NITRO-PASS", c.password)
			}
		}
	} else {
		req.SetBasicAuth(c.username, c.password)
		req.Header.Set("_MPS_API_PROXY_MANAGED_INSTANCE_IP", c.proxiedNs)
	}

	// User defined headers may overwrite previous headers
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}
	return req, nil
}

func (c *NitroClient) doHTTPRequest(method string, urlstr string, bytes *bytes.Buffer, respHandler responseHandlerFunc) ([]byte, error) {
	req, err := c.createHTTPRequest(method, urlstr, bytes)
	log.Printf("[TRACE] go-nitro: doHTTPRequest HTTP method: %v", method)
	log.Printf("[TRACE] go-nitro: doHTTPRequest HTTP url: %v", urlstr)
	log.Printf("[TRACE] go-nitro: doHTTPRequest HTTP body: %v", bytes.String())
	log.Printf("[TRACE] go-nitro: doHTTPRequest HTTP request headers: %v", req.Header)

	resp, err := c.client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return []byte{}, err
	}
	log.Println("[DEBUG] go-nitro: response Status:", resp.Status)
	body, err := respHandler(resp)
	// Clear sessionid in case of session-expiry
	if resp.Status == "401 Unauthorized" {
		var data map[string]interface{}
		err2 := json.Unmarshal(body, &data)
		if err2 == nil {
			errorcode, ok := data["errorcode"]
			if ok {
				errorcode = int(errorcode.(float64))
				if errorcode == nsErrSessionExpired || errorcode == nsErrAuthTimeout {
					c.clearSessionid()
				}
			}
		}
	}
	return body, err
}

func (c *NitroClient) createResource(resourceType string, resourceJSON []byte) ([]byte, error) {
	log.Println("[DEBUG] go-nitro: Creating resource of type ", resourceType)

	url := c.url + resourceType

	if !strings.HasSuffix(resourceType, "_binding") && !contains(idempotentInvalidResources, resourceType) {
		url = url + "?idempotent=yes"
	}
	log.Println("[TRACE] go-nitro: url is ", url)

	return c.doHTTPRequest("POST", url, bytes.NewBuffer(resourceJSON), createResponseHandler)

}

func (c *NitroClient) applyResource(resourceType string, resourceJSON []byte) ([]byte, error) {
	log.Println("[DEBUG] go-nitro: Applying resource of type ", resourceType)

	url := c.url + resourceType + "?action=apply"
	log.Println("[TRACE] go-nitro: url is ", url)

	return c.doHTTPRequest("POST", url, bytes.NewBuffer(resourceJSON), createResponseHandler)

}

func (c *NitroClient) actOnResource(resourceType string, resourceJSON []byte, action string) ([]byte, error) {
	log.Println("[DEBUG] go-nitro: changing resource of type ", resourceType)

	var url string
	if action == "" {
		url = c.url + fmt.Sprintf("%s", resourceType)
	} else {
		url = c.url + fmt.Sprintf("%s?action=%s", resourceType, action)
	}
	log.Println("[TRACE] go-nitro: url is ", url)

	return c.doHTTPRequest("POST", url, bytes.NewBuffer(resourceJSON), createResponseHandler)

}

func (c *NitroClient) changeResource(resourceType string, resourceName string, resourceJSON []byte) ([]byte, error) {
	log.Println("[DEBUG] go-nitro: changing resource of type ", resourceType)

	url := c.url + resourceType + "/" + resourceName + "?action=update"
	log.Println("[TRACE] go-nitro: url is ", url)

	return c.doHTTPRequest("POST", url, bytes.NewBuffer(resourceJSON), createResponseHandler)

}

func (c *NitroClient) updateResource(resourceType string, resourceName string, resourceJSON []byte) ([]byte, error) {
	log.Println("[DEBUG] go-nitro: Updating resource of type ", resourceType)

	url := c.url + resourceType + "/" + resourceName
	log.Println("[TRACE] go-nitro: url is ", url)

	return c.doHTTPRequest("PUT", url, bytes.NewBuffer(resourceJSON), createResponseHandler)

}

func (c *NitroClient) updateUnnamedResource(resourceType string, resourceJSON []byte) ([]byte, error) {
	log.Println("[DEBUG] go-nitro: Updating resource of type ", resourceType)

	url := c.url + resourceType
	log.Println("[TRACE] go-nitro: url is ", url)

	return c.doHTTPRequest("PUT", url, bytes.NewBuffer(resourceJSON), createResponseHandler)

}

func (c *NitroClient) deleteResource(resourceType string, resourceName string) ([]byte, error) {
	log.Println("[DEBUG] go-nitro: Deleting resource of type ", resourceType)
	var url string
	if resourceName != "" {
		url = c.url + fmt.Sprintf("%s/%s", resourceType, resourceName)
	} else {
		url = c.url + fmt.Sprintf("%s", resourceType)
	}
	log.Println("[TRACE] go-nitro: url is ", url)

	return c.doHTTPRequest("DELETE", url, bytes.NewBuffer([]byte{}), deleteResponseHandler)

}

func (c *NitroClient) deleteResourceWithArgs(resourceType string, resourceName string, args []string) ([]byte, error) {
	log.Println("[DEBUG] go-nitro: Deleting resource of type ", resourceType, "with args ", args)
	var url string
	if resourceName != "" {
		url = c.url + fmt.Sprintf("%s/%s?args=", resourceType, resourceName)
	} else {
		url = c.url + fmt.Sprintf("%s?args=", resourceType)
	}
	url = url + strings.Join(args, ",")
	log.Println("[TRACE] go-nitro: url is ", url)

	return c.doHTTPRequest("DELETE", url, bytes.NewBuffer([]byte{}), deleteResponseHandler)

}

func (c *NitroClient) deleteResourceWithArgsMap(resourceType string, resourceName string, argsMap map[string]string) ([]byte, error) {
	args := make([]string, len(argsMap))
	i := 0
	for key, value := range argsMap {
		args[i] = fmt.Sprintf("%s:%s", key, value)
		i++
	}
	return c.deleteResourceWithArgs(resourceType, resourceName, args)

}

func (c *NitroClient) unbindResource(resourceType string, resourceName string, boundResourceType string, boundResource string, bindingFilterName string) ([]byte, error) {
	log.Println("[DEBUG] go-nitro: Unbinding resource of type ", resourceType, " name", resourceName)
	bindingName := resourceType + "_" + boundResourceType + "_binding"

	url := c.url + "/" + bindingName + "/" + resourceName + "?args=" + bindingFilterName + ":" + boundResource

	return c.doHTTPRequest("DELETE", url, bytes.NewBuffer([]byte{}), deleteResponseHandler)

}

func (c *NitroClient) listBoundResources(resourceName string, resourceType string, boundResourceType string, boundResourceFilterName string, boundResourceFilterValue string) ([]byte, error) {
	log.Println("[DEBUG] go-nitro: listing bound resources of type ", resourceType, ": ", resourceName)
	var url string
	if boundResourceFilterName == "" {
		url = c.url + fmt.Sprintf("%s_%s_binding/%s", resourceType, boundResourceType, resourceName)
	} else {
		url = c.url + fmt.Sprintf("%s_%s_binding/%s?filter=%s:%s", resourceType, boundResourceType, resourceName, boundResourceFilterName, boundResourceFilterValue)
	}

	return c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}), readResponseHandler)

}

func (c *NitroClient) listFilteredResource(resourceType string, filter map[string]string) ([]byte, error) {
	log.Println("[DEBUG] go-nitro: listing resource of type ", resourceType, ", filter: ", filter)

	var filter_strings []string
	for key, value := range filter {
		filter_strings = append(filter_strings, fmt.Sprintf("%s:%s", key, value))
	}

	filter_string := strings.Join(filter_strings, ",")

	url := c.url + fmt.Sprintf("%s?filter=%s", resourceType, filter_string)

	return c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}), readResponseHandler)

}

func (c *NitroClient) listResource(resourceType string, resourceName string) ([]byte, error) {
	log.Println("[DEBUG] go-nitro: listing resource of type ", resourceType, ", name: ", resourceName)
	url := c.url + resourceType

	if resourceName != "" {
		url = c.url + fmt.Sprintf("%s/%s", resourceType, resourceName)
	}
	log.Println("[TRACE] go-nitro: url is ", url)

	return c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}), readResponseHandler)

}

func (c *NitroClient) listResourceWithArgs(resourceType string, resourceName string, args []string) ([]byte, error) {
	log.Println("[DEBUG] go-nitro: listing resource of type ", resourceType, ", name: ", resourceName, ", args:", args)
	var url string

	if resourceName != "" {
		url = c.url + fmt.Sprintf("%s/%s", resourceType, resourceName)
	} else {
		url = c.url + fmt.Sprintf("%s", resourceType)
	}
	strArgs := strings.Join(args, ",")
	url2 := url + "?args=" + strArgs
	log.Println("[TRACE] go-nitro: url is ", url)

	data, err := c.doHTTPRequest("GET", url2, bytes.NewBuffer([]byte{}), readResponseHandler)
	if err != nil {
		log.Println("[DEBUG] go-nitro: error listing with args, trying filter")
		url2 = url + "?filter=" + strArgs
		return c.doHTTPRequest("GET", url2, bytes.NewBuffer([]byte{}), readResponseHandler)
	}
	return data, err

}

func (c *NitroClient) listResourceWithArgsMap(resourceType string, resourceName string, argsMap map[string]string) ([]byte, error) {
	args := make([]string, len(argsMap))
	i := 0
	for key, value := range argsMap {
		args[i] = fmt.Sprintf("%s:%s", key, value)
		i++
	}
	return c.listResourceWithArgs(resourceType, resourceName, args)

}

func (c *NitroClient) enableFeatures(featureJSON []byte) ([]byte, error) {
	log.Println("[DEBUG] go-nitro Enabling features")
	url := c.url + "nsfeature?action=enable"

	return c.doHTTPRequest("POST", url, bytes.NewBuffer(featureJSON), createResponseHandler)

}

func (c *NitroClient) disableFeatures(featureJSON []byte) ([]byte, error) {
	log.Println("[DEBUG] go-nitro Disabling features")
	url := c.url + "nsfeature?action=disable"

	return c.doHTTPRequest("POST", url, bytes.NewBuffer(featureJSON), createResponseHandler)

}

func (c *NitroClient) listEnabledFeatures() ([]byte, error) {
	log.Println("[DEBUG] go-nitro: listing features")
	url := c.url + "nsfeature"

	return c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}), readResponseHandler)

}

func (c *NitroClient) enableModes(modeJSON []byte) ([]byte, error) {
	log.Println("[DEBUG] go-nitro Enabling modes")
	url := c.url + "nsmode?action=enable"

	return c.doHTTPRequest("POST", url, bytes.NewBuffer(modeJSON), createResponseHandler)

}

func (c *NitroClient) listEnabledModes() ([]byte, error) {
	log.Println("[DEBUG] go-nitro: listing modes")
	url := c.url + "nsmode"

	return c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}), readResponseHandler)

}

func (c *NitroClient) saveConfig(saveJSON []byte) error {
	log.Println("[DEBUG] go-nitro: Saving config")
	url := c.url + "nsconfig?action=save"

	_, err := c.doHTTPRequest("POST", url, bytes.NewBuffer(saveJSON), createResponseHandler)
	return err

}

func (c *NitroClient) clearConfig(clearJSON []byte) error {
	log.Println("[DEBUG] go-nitro: Clearing config")
	url := c.url + "nsconfig?action=clear"

	_, err := c.doHTTPRequest("POST", url, bytes.NewBuffer(clearJSON), createResponseHandler)
	return err
}

func (c *NitroClient) listStat(resourceType, resourceName string) ([]byte, error) {
	log.Println("[DEBUG] go-nitro: listing resource of type ", resourceType, ", name: ", resourceName)
	url := c.statsURL + resourceType

	if resourceName != "" {
		url = c.statsURL + fmt.Sprintf("%s/%s", resourceType, resourceName)
	}
	log.Println("[TRACE] go-nitro: url is ", url)

	return c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}), readResponseHandler)

}

func (c *NitroClient) listStatWithArgs(resourceType string, resourceName string, args []string) ([]byte, error) {
	log.Println("[DEBUG] go-nitro: listing stat of type ", resourceType, ", name: ", resourceName, ", args:", args)
	var url string

	if len(resourceName) > 0 {
		url = c.statsURL + fmt.Sprintf("%s/%s", resourceType, resourceName)
	} else {
		url = c.statsURL + fmt.Sprintf("%s", resourceType)
	}
	strArgs := strings.Join(args, ",")
	url = url + "?args=" + strArgs
	log.Println("[TRACE] go-nitro: url is ", url)

	return c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}), readResponseHandler)
}
