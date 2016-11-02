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
package netscaler

import (
	"github.com/chiradeep/go-nitro/config/basic"
	"github.com/chiradeep/go-nitro/config/lb"
	"log"
	"testing"
)

/*
func TestClientCreate() {
	client, err := netscaler.NewNitroClientFromEnv("http://127.0.0.1:32774")
	if err != nil {
		log.Fatal("Could not create a client: ", err)
	}
	log.Printf("Client is %+v\n", *client)

}
*/

// Functional test
func TestAdd(t *testing.T) {
	client, err := NewNitroClientFromEnv()
	if err != nil {
		log.Fatal("Could not create a client: ", err)
	}

	lb1 := lb.Lbvserver{
		Name:        "sample_lb",
		Ipv46:       "10.71.136.50",
		Lbmethod:    "ROUNDROBIN",
		Servicetype: "HTTP",
		Port:        8000,
	}
	client.AddResource(Lbvserver.Type(), "sample_lb", &lb1)

	rsrc, err := client.FindResource(Lbvserver.Type(), "sample_lb")
	if err != nil {
		t.Error("Did not find resource of type ", Lbvserver.Type(), ":", "sample_lb")
	}
	val, ok := rsrc["ipv46"]
	if ok {
		if val != "10.71.136.50" {
			t.Error("Wrong ipv46 for lb ", "sample_lb", ": ", val)
		}
		val, ok = rsrc["lbmethod"]
		if val != "ROUNDROBIN" {
			t.Error("Wrong lbmethod for lb ", "sample_lb", ": ", val)
		}
		val, ok = rsrc["servicetype"]
		if val != "HTTP" {
			t.Error("Wrong servicetype for lb ", "sample_lb", ": ", val)
		}
	}
	if !ok {
		t.Error("Non existent property in retrieved lb ", "sample_lb")
	}

	service1 := basic.Service{
		Name:        "sample_svc_1",
		Ip:          "172.22.33.4",
		Port:        80,
		Servicetype: "HTTP",
	}

	client.AddResource(Service.Type(), "sample_svc_1", &service1)

	_, err = client.FindResource(Service.Type(), "sample_svc_1")
	if err != nil {
		t.Error("Did not find resource of type ", Service.Type(), ":", "sample_svc_1")
	}
}

func TestUpdate(t *testing.T) {
	client, err := NewNitroClientFromEnv()
	if err != nil {
		log.Fatal("Could not create a client: ", err)
	}

	lb1 := lb.Lbvserver{
		Name:     "sample_lb",
		Lbmethod: "LEASTCONNECTION",
	}
	client.UpdateResource(Lbvserver.Type(), "sample_lb", &lb1)
	rsrc, err := client.FindResource(Lbvserver.Type(), "sample_lb")
	if err != nil {
		t.Error("Did not find resource of type ", Lbvserver.Type(), ":", "sample_lb")
	}
	val, ok := rsrc["lbmethod"]
	if ok {
		if val != "LEASTCONNECTION" {
			t.Error("Did not update lb method to LEASTCONNECTION")
		}
	}
	if !ok {
		t.Error("Failed to retrieve lb vserver object")
	}
}

func TestBind(t *testing.T) {
	client, err := NewNitroClientFromEnv()
	if err != nil {
		log.Fatal("Could not create a client: ", err)
	}

	binding := lb.Lbvserverservicebinding{
		Name:        "sample_lb",
		Servicename: "sample_svc_1",
	}

	client.BindResource(Lbvserver.Type(), "sample_lb", Service.Type(), "sample_svc_1", &binding)
	exists := client.ResourceBindingExists(Lbvserver.Type(), "sample_lb", Service.Type(), "servicename", "sample_svc_1")
	if !exists {
		t.Error("Failed to bind service to lb vserver")
	}

}

func TestUnbind(t *testing.T) {
	client, err := NewNitroClientFromEnv()
	if err != nil {
		log.Fatal("Could not create a client: ", err)
	}

	client.UnbindResource(Lbvserver.Type(), "sample_lb", Service.Type(), "sample_svc_1", "servicename")
	exists := client.ResourceBindingExists(Lbvserver.Type(), "sample_lb", Service.Type(), "servicename", "sample_svc_1")
	if exists {
		t.Error("Failed to unbind service to lb vserver")
	}

}

func TestDelete(t *testing.T) {
	client, err := NewNitroClientFromEnv()
	if err != nil {
		log.Fatal("Could not create a client: ", err)
	}
	client.DeleteResource(Lbvserver.Type(), "sample_lb")
	if client.ResourceExists(Lbvserver.Type(), "sample_lb") {
		t.Error("Failed to delete ", "sample_lb")
	}
	client.DeleteResource(Service.Type(), "sample_svc_1")
	if client.ResourceExists(Service.Type(), "sample_svc_1") {
		t.Error("Failed to delete ", "sample_svc_1")
	}
}

func TestEnableFeatures(t *testing.T) {
	client, err := NewNitroClientFromEnv()
	if err != nil {
		log.Fatal("Could not create a client: ", err)
	}
	features := []string{"SSL", "CS"}
	err = client.EnableFeatures(features)
	if err != nil {
		t.Error("Failed to enable features")
	}
	result, err := client.ListEnabledFeatures()
	if err != nil {
		t.Error("Failed to retrieve features")
	}
	found := 0
	for _, f := range features {
		for _, r := range result {
			if f == r {
				found = found + 1
			}
		}
	}
	if found != len(features) {
		t.Error("Requested features do not match enabled features=", features, "result=", result)
	}
}

func TestSaveConfig(t *testing.T) {
	client, err := NewNitroClientFromEnv()
	if err != nil {
		log.Fatal("Could not create a client: ", err)
	}
	err = client.SaveConfig()
	if err != nil {
		t.Error("Failed to save config")
	}
}

func TestFindAllResources(t *testing.T) {
	client, err := NewNitroClientFromEnv()
	if err != nil {
		log.Fatal("Could not create a client: ", err)
	}
	lb1 := lb.Lbvserver{
		Name:        "sample_lb_1",
		Ipv46:       "10.71.136.50",
		Lbmethod:    "ROUNDROBIN",
		Servicetype: "HTTP",
		Port:        8000,
	}
	lb2 := lb.Lbvserver{
		Name:        "sample_lb_2",
		Ipv46:       "10.71.136.60",
		Lbmethod:    "LEASTCONNECTION",
		Servicetype: "HTTP",
		Port:        8000,
	}
	_, err = client.AddResource(Lbvserver.Type(), "sample_lb_1", &lb1)
	if err != nil {
		t.Error("Failed to add resource of type ", Lbvserver.Type(), ":", "sample_lb_1")
	}
	_, err = client.AddResource(Lbvserver.Type(), "sample_lb_2", &lb2)
	if err != nil {
		t.Error("Failed to add resource of type ", Lbvserver.Type(), ":", "sample_lb_2")
	}
	rsrcs, err := client.FindAllResources(Lbvserver.Type())
	if err != nil {
		t.Error("Did not find resources of type ", Lbvserver.Type())
	}
	if len(rsrcs) < 2 {
		t.Error("Found only ", len(rsrcs), " resources of type ", Lbvserver.Type(), " expected at least 2")
	}

	found := 0
	for _, v := range rsrcs {
		name := v["name"].(string)
		log.Println("Found lb vserver name=", name)
		if name == "sample_lb_1" || name == "sample_lb_2" {
			found = found + 1
		}
	}
	if found != 2 {
		t.Error("Did not find all configured lbvservers")
	}

}
