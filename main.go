/*
Copyright 2016 Citrix Systems, Inc

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
package main

import (
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/chiradeep/go-nitro/netscaler/basic"
	"github.com/chiradeep/go-nitro/netscaler/lb"
)

func main() {
	client := netscaler.NewNitroClient("http://127.0.0.1:32774", "nsroot", "nsroot")
	lb1 := lb.Lbvserver{
		Name:        "sample_lb",
		Ipv46:       "10.71.136.50",
		Lbmethod:    "ROUNDROBIN",
		Servicetype: "HTTP",
		Port:        8000,
	}
	client.AddResource("sample_lb", "lbvserver", &lb1)

	lb1 = lb.Lbvserver{
		Name:     "sample_lb",
		Lbmethod: "LEASTCONNECTION",
	}
	client.UpdateResource("sample_lb", "lbvserver", &lb1)

	service1 := basic.Service{
		Name:        "service_1",
		Ip:          "172.22.33.4",
		Port:        80,
		Servicetype: "HTTP",
	}

	client.AddResource("service_1", "service", &service1)

	binding := lb.Lbvserverservicebinding{
		Name:        "sample_lb",
		Servicename: "service_1",
	}

	client.BindResource("sample_lb", "lbvserver", "service_1", "service", &binding)

	client.UnbindResource("sample_lb", "lbvserver", "service_1", "service", "servicename")

	client.DeleteResource("sample_lb", "lbvserver")
	client.DeleteResource("service_1", "service")
}
