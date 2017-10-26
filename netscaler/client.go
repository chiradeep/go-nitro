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
	"fmt"
	"net/http"
	"crypto/tls"
	"os"
	"strings"
)

//NitroClient has methods to configure the NetScaler
//It abstracts the REST operations of the NITRO API
type NitroClient struct {
	url       string
	username  string
	password  string
	proxiedNs string
	client    *http.Client
}

//NewNitroClient returns a usable NitroClient. Does not check validity of supplied parameters
func NewNitroClient(url string, username string, password string, bool sslverify) *NitroClient {
	c := new(NitroClient)
	c.url = strings.Trim(url, " /") + "/nitro/v1/config/"
	c.username = username
	c.password = password

	if sslverify == true {
		c.client = &http.Client{}
	} else {

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		c.client = &http.Client{Transport: tr}
	}

	return c
}

func NewProxyingNitroClient(url string, username string, password string, proxiedNsIp string, bool sslverify) *NitroClient {
	c := new(NitroClient)
	c.url = strings.Trim(url, " /") + "/nitro/v1/config/"
	c.username = username
	c.password = password
	c.proxiedNs = proxiedNsIp

	if sslverify == true {
		c.client = &http.Client{}
	} else {

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		c.client = &http.Client{Transport: tr}
	}

	return c
}

//NewNitroClientFromEnv returns a usable NitroClient. Parameters url, username and password can be passed in
//as the first three positional parameters. Otherwise, it tries to read these values from
//environment variable NS_URL, NS_LOGIN and NS_PASSWORD
func NewNitroClientFromEnv(args ...string) (*NitroClient, error) {
	url := os.Getenv("NS_URL")
	argslen := len(args)
	if url == "" {
		if argslen >= 1 {
			url = args[0]
		} else {
			return nil, fmt.Errorf("Could not determine NetScaler URL: NS_URL not set or passed in as first parameter")
		}
	}
	username := os.Getenv("NS_LOGIN")
	if username == "" {
		if argslen >= 2 {
			url = args[1]
		} else {
			return nil, fmt.Errorf("Could not determine NetScaler login: NS_LOGIN not set or passed in as second parameter")
		}
	}
	password := os.Getenv("NS_PASSWORD")
	if password == "" {
		if argslen >= 3 {
			url = args[2]
		} else {
			return nil, fmt.Errorf("Could not determine NetScaler password: NS_PASSWORD not set or passed in as third parameter")
		}
	}
	password := os.Getenv("NS_SSLVERIFY")
	if sslverify == "" {
		if argslen >= 4 {
			url = args[3]
		} else {
			return nil, fmt.Errorf("Could not determine ssl verification: NS_SSLVERIFY not set or passed in as third parameter")
		}
	}
	proxiedNs := os.Getenv("_MPS_API_PROXY_MANAGED_INSTANCE_IP")
	if proxiedNs == "" {
		return NewNitroClient(url, username, password, sslverify), nil
	} else {
		return NewProxyingNitroClient(url, username, password, proxiedNs, sslverify), nil
	}
}
