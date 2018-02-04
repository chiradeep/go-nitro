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
	"strconv"
)

var SSLVerify bool = true

//NitroClient has methods to configure the NetScaler
//It abstracts the REST operations of the NITRO API
type NitroClient struct {
	url       string
	username  string
	password  string
	proxiedNs string
	client    *http.Client
}

// DisableSSLVerify sets our SSL verification flag to disabled
func DisableSSLVerify(){
	SSLVerify = false
}

// GetClient returns a standard client or one with SSL checking disabled
func GetClient() *http.Client {
	if SSLVerify == true {
		return &http.Client{}
	} else {

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		return &http.Client{Transport: tr}
	}
}

//NewNitroClient returns a usable NitroClient. Does not check validity of supplied parameters
func NewNitroClient(url string, username string, password string) *NitroClient {
	c := new(NitroClient)
	c.url = strings.Trim(url, " /") + "/nitro/v1/config/"
	c.username = username
	c.password = password
	c.client = GetClient()

	return c
}

func NewProxyingNitroClient(url string, username string, password string, proxiedNsIp string) *NitroClient {
	c := new(NitroClient)
	c.url = strings.Trim(url, " /") + "/nitro/v1/config/"
	c.username = username
	c.password = password
	c.proxiedNs = proxiedNsIp
	c.client = GetClient()

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
	sslverify, err := strconv.ParseBool(os.Getenv("NS_SSLVERIFY"))
	if err != nil  {
		if argslen >= 4 {
			url = args[3]
		} else {
			return nil, fmt.Errorf("Could not convert SSL vefification type to true or false: NS_SSLVERIFY not set properly or passed in as fourth parameter")
		}
	}
	proxiedNs := os.Getenv("_MPS_API_PROXY_MANAGED_INSTANCE_IP")

	if sslverify == false {
		DisableSSLVerify()
	}

	if proxiedNs == "" {
		return NewNitroClient(url, username, password), nil
	} else {
		return NewProxyingNitroClient(url, username, password, proxiedNs), nil
	}
}
