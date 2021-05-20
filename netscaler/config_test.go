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
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"os"

	"github.com/chiradeep/go-nitro/config/basic"
	"github.com/chiradeep/go-nitro/config/lb"
	"github.com/chiradeep/go-nitro/config/network"
	"github.com/chiradeep/go-nitro/config/ns"
)

var client *NitroClient

//Used to generate random config object names
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(125)+1, rand.Intn(252)+1, rand.Intn(252)+1, rand.Intn(252)+1)
}

//init random and client
func init() {
	rand.Seed(time.Now().UnixNano())
	var err error
	client, err = NewNitroClientFromEnv()
	if err != nil {
		log.Fatal("Could not create a client: ", err)
	}
	client.SetLogLevel("TRACE")

}

func TestMain(m *testing.M) {
	r := m.Run()
	client.ClearConfig()
	os.Exit(r)
}

// Functional tests
func TestClearConfig(t *testing.T) {
	err := client.ClearConfig()
	if err != nil {
		t.Error("Could not clear config: ", err)
	}
}

func TestAdd(t *testing.T) {

	rndIP := randomIP()
	lbName := "test_lb_" + randomString(5)

	lb1 := lb.Lbvserver{
		Name:        lbName,
		Ipv46:       rndIP,
		Lbmethod:    "ROUNDROBIN",
		Servicetype: "HTTP",
		Port:        8000,
	}
	_, err := client.AddResource(Lbvserver.Type(), lbName, &lb1)
	if err != nil {
		t.Error("Could not add Lbvserver: ", err)
		log.Println("Not continuing test")
		return
	}

	rsrc, err := client.FindResource(Lbvserver.Type(), lbName)
	if err != nil {
		t.Error("Did not find resource of type ", err, Lbvserver.Type(), ":", lbName)
	}
	val, ok := rsrc["ipv46"]
	if ok {
		if val != rndIP {
			t.Error("Wrong ipv46 for lb ", lbName, ": ", val)
		}
		val, ok = rsrc["lbmethod"]
		if val != "ROUNDROBIN" {
			t.Error("Wrong lbmethod for lb ", lbName, ": ", val)
		}
		val, ok = rsrc["servicetype"]
		if val != "HTTP" {
			t.Error("Wrong servicetype for lb ", err, lbName, ": ", val)
		}
	}
	if !ok {
		t.Error("Non existent property in retrieved lb ", lbName)
	}

	svcName := randomString(5)
	rndIP2 := randomIP()

	service1 := basic.Service{
		Name:        svcName,
		Ip:          rndIP2,
		Port:        80,
		Servicetype: "HTTP",
	}

	client.AddResource(Service.Type(), svcName, &service1)

	_, err = client.FindResource(Service.Type(), svcName)
	if err != nil {
		t.Error("Did not find resource of type ", err, Service.Type(), ":", svcName)
	}
}

func TestApply(t *testing.T) {
	aclName := "test_acl_" + randomString(5)
	acl1 := ns.Nsacl{
		Aclname:   aclName,
		Aclaction: "ALLOW",
		Srcip:     true,
		Srcipval:  "192.168.11.10",
		Destip:    true,
		Destipval: "192.183.83.11",
		Priority:  1100,
	}

	_, err := client.AddResource(Nsacl.Type(), aclName, &acl1)
	if err != nil {
		t.Error("Could not add resource Nsacl", err)
		log.Println("Cannot continue")
		return
	}

	acls := ns.Nsacls{}
	client.ApplyResource(Nsacls.Type(), &acls)

	readAcls, err := client.FindResourceArray(Nsacl.Type(), aclName)
	if err != nil {
		t.Error("Did not find resource of type ", Nsacl.Type(), err, ":", aclName)
	}
	if err == nil {
		acl2 := readAcls[0]
		log.Println("Found acl, kernelstate= ", acl2["kernelstate"])
		if acl2["kernelstate"].(string) != "APPLIED" {
			t.Error("ACL created but not APPLIED ", Nsacl.Type(), ":", aclName)
		}
	}
}

func TestUpdate(t *testing.T) {
	rndIP := randomIP()
	lbName := "test_lb_" + randomString(5)

	lb1 := lb.Lbvserver{
		Name:        lbName,
		Ipv46:       rndIP,
		Lbmethod:    "ROUNDROBIN",
		Servicetype: "HTTP",
		Port:        8000,
	}
	_, err := client.AddResource(Lbvserver.Type(), lbName, &lb1)
	if err != nil {
		t.Error("Could not create LB", err)
		log.Println("Cannot continue")
		return
	}

	lb1 = lb.Lbvserver{
		Name:     lbName,
		Lbmethod: "LEASTCONNECTION",
	}
	_, err = client.UpdateResource(Lbvserver.Type(), lbName, &lb1)
	if err != nil {
		t.Error("Could not update LB")
		log.Println("Cannot continue")
		return
	}
	rsrc, err := client.FindResource(Lbvserver.Type(), lbName)
	if err != nil {
		t.Error("Did not find resource of type ", Lbvserver.Type(), ":", lbName, err)
		log.Println("Cannot continue")
		return
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

func TestBindUnBind(t *testing.T) {
	rndIP := randomIP()
	lbName := "test_lb_" + randomString(5)
	rndIP2 := randomIP()
	svcName := "test_svc_" + randomString(5)

	lb1 := lb.Lbvserver{
		Name:        lbName,
		Ipv46:       rndIP,
		Lbmethod:    "ROUNDROBIN",
		Servicetype: "HTTP",
		Port:        8000,
	}
	_, err := client.AddResource(Lbvserver.Type(), lbName, &lb1)
	if err != nil {
		t.Error("Could not create LB", err)
		log.Println("Cannot continue")
		return
	}
	service1 := basic.Service{
		Name:        svcName,
		Ip:          rndIP2,
		Port:        80,
		Servicetype: "HTTP",
	}

	_, err = client.AddResource(Service.Type(), svcName, &service1)
	if err != nil {
		t.Error("Could not create service", err)
		log.Println("Cannot continue")
		return
	}

	binding := lb.Lbvserverservicebinding{
		Name:        lbName,
		Servicename: svcName,
	}

	err = client.BindResource(Lbvserver.Type(), lbName, Service.Type(), svcName, &binding)
	if err != nil {
		t.Error("Could not bind LB to svc", err)
		log.Println("Cannot continue")
		return
	}
	exists := client.ResourceBindingExists(Lbvserver.Type(), lbName, Service.Type(), "servicename", svcName)
	if !exists {
		t.Error("Failed to bind service to lb vserver")
		log.Println("Cannot continue")
		return
	}
	err = client.UnbindResource(Lbvserver.Type(), lbName, Service.Type(), svcName, "servicename")
	if err != nil {
		t.Error("Could not unbind LB to svc", err)
		log.Println("Cannot continue")
		return
	}
	exists = client.ResourceBindingExists(Lbvserver.Type(), lbName, Service.Type(), "servicename", svcName)
	if exists {
		t.Error("Failed to unbind service to lb vserver")
	}

}

func TestFindBoundResource(t *testing.T) {
	lbName := "test_lb_" + randomString(5)
	lb1 := lb.Lbvserver{
		Name:        lbName,
		Ipv46:       randomIP(),
		Lbmethod:    "ROUNDROBIN",
		Servicetype: "HTTP",
		Port:        8000,
	}
	_, err := client.AddResource(Lbvserver.Type(), lbName, &lb1)
	if err != nil {
		t.Error("Failed to add resource of type ", Lbvserver.Type(), ":", "sample_lb_1", err)
		log.Println("Cannot continue")
		return
	}
	svcName := "test_svc_" + randomString(5)
	service1 := basic.Service{
		Name:        svcName,
		Ip:          randomIP(),
		Port:        80,
		Servicetype: "HTTP",
	}

	_, err = client.AddResource(Service.Type(), svcName, &service1)
	if err != nil {
		t.Error("Failed to add resource of type ", Service.Type(), ":", svcName, err)
		log.Println("Cannot continue")
		return

	}
	binding := lb.Lbvserverservicebinding{
		Name:        lbName,
		Servicename: svcName,
	}
	err = client.BindResource(Lbvserver.Type(), lbName, Service.Type(), svcName, &binding)
	if err != nil {
		t.Error("Failed to bind resource of type ", Service.Type(), ":", svcName)
		log.Println("Cannot continue")
		return

	}
	result, err := client.FindBoundResource(Lbvserver.Type(), lbName, Service.Type(), "servicename", svcName)
	if err != nil {
		t.Error("Failed to find bound resource of type ", Service.Type(), ":", svcName)
		log.Println("Cannot continue")
		return

	}
	//log.Println("Found bound resource ", result)
	if result["servicename"] != svcName {
		t.Error("Failed to find bound resource of type ", Service.Type(), ":", svcName)
	}

}

func TestDelete(t *testing.T) {
	rndIP := randomIP()
	lbName := "test_lb_" + randomString(5)

	lb1 := lb.Lbvserver{
		Name:        lbName,
		Ipv46:       rndIP,
		Lbmethod:    "ROUNDROBIN",
		Servicetype: "HTTP",
		Port:        8000,
	}
	_, err := client.AddResource(Lbvserver.Type(), lbName, &lb1)
	if err != nil {
		t.Error("Could not create LB", err)
		log.Println("Cannot continue")
		return
	}

	err = client.DeleteResource(Lbvserver.Type(), lbName)
	if err != nil {
		t.Error("Could not delete LB", lbName, err)
		log.Println("Cannot continue")
		return
	}
	if client.ResourceExists(Lbvserver.Type(), lbName) {
		t.Error("Failed to delete ", lbName)
	}
}

func TestDeleteWithArgs(t *testing.T) {
	monitorName := "test_lb_monitor_" + randomString(5)

	lbmonitor := lb.Lbmonitor{
		Monitorname:    monitorName,
		Type:           "http",
		Retries:        20,
		Failureretries: 10,
		Downtime:       60,
	}
	_, err := client.AddResource(Lbmonitor.Type(), monitorName, &lbmonitor)
	if err != nil {
		t.Error("Could not create monitor", err)
		log.Println("Cannot continue")
		return
	}

	args := map[string]string{"type": "http"}
	err = client.DeleteResourceWithArgsMap(Lbmonitor.Type(), monitorName, args)
	if err != nil {
		t.Error("Could not delete monitor", monitorName, err)
		log.Println("Cannot continue")
		return
	}
}

func TestEnableFeatures(t *testing.T) {
	features := []string{"SSL", "CS"}
	err := client.EnableFeatures(features)
	if err != nil {
		t.Error("Failed to enable features", err)
		log.Println("Cannot continue")
		return
	}
	result, err := client.ListEnabledFeatures()
	if err != nil {
		t.Error("Failed to retrieve features", err)
		log.Println("Cannot continue")
		return
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

func TestEnableModes(t *testing.T) {
	modes := []string{"ULFD", "MBF"}
	err := client.EnableModes(modes)
	if err != nil {
		t.Error("Failed to enable modes", err)
		log.Println("Cannot continue")
		return
	}
	result, err := client.ListEnabledModes()
	if err != nil {
		t.Error("Failed to retrieve modes", err)
		log.Println("Cannot continue")
		return
	}
	found := 0
	for _, m := range modes {
		for _, r := range result {
			if m == r {
				found = found + 1
			}
		}
	}
	if found != len(modes) {
		t.Error("Requested modes do not match enabled modes=", modes, "result=", result)
	}
}

func TestSaveConfig(t *testing.T) {
	err := client.SaveConfig()
	if err != nil {
		t.Error("Failed to save config", err)
	}
}

func TestFindAllResources(t *testing.T) {
	lbName1 := "test_lb_" + randomString(5)
	lbName2 := "test_lb_" + randomString(5)
	lb1 := lb.Lbvserver{
		Name:        lbName1,
		Ipv46:       randomIP(),
		Lbmethod:    "ROUNDROBIN",
		Servicetype: "HTTP",
		Port:        8000,
	}
	lb2 := lb.Lbvserver{
		Name:        lbName2,
		Ipv46:       randomIP(),
		Lbmethod:    "LEASTCONNECTION",
		Servicetype: "HTTP",
		Port:        8000,
	}
	_, err := client.AddResource(Lbvserver.Type(), lbName1, &lb1)
	if err != nil {
		t.Error("Failed to add resource of type ", Lbvserver.Type(), ":", lbName1)
		log.Println("Cannot continue")
		return
	}
	_, err = client.AddResource(Lbvserver.Type(), lbName2, &lb2)
	if err != nil {
		t.Error("Failed to add resource of type ", Lbvserver.Type(), ":", lbName2)
		log.Println("Cannot continue")
		return
	}
	rsrcs, err := client.FindAllResources(Lbvserver.Type())
	if err != nil {
		t.Error("Did not find resources of type ", Lbvserver.Type(), err)
	}
	if len(rsrcs) < 2 {
		t.Error("Found only ", len(rsrcs), " resources of type ", Lbvserver.Type(), " expected at least 2")
	}

	found := 0
	for _, v := range rsrcs {
		name := v["name"].(string)
		if name == lbName1 || name == lbName2 {
			found = found + 1
		}
	}
	if found != 2 {
		t.Error("Did not find all configured lbvservers")
	}

}

func TestFindAllBoundResources(t *testing.T) {
	lbName1 := "test_lb_" + randomString(5)
	svcName1 := "test_svc_" + randomString(5)
	svcName2 := "test_svc_" + randomString(5)
	lb1 := lb.Lbvserver{
		Name:        lbName1,
		Ipv46:       randomIP(),
		Lbmethod:    "ROUNDROBIN",
		Servicetype: "HTTP",
		Port:        8000,
	}
	_, err := client.AddResource(Lbvserver.Type(), lbName1, &lb1)
	if err != nil {
		t.Error("Could not create LB")
	}
	service1 := basic.Service{
		Name:        svcName1,
		Ip:          randomIP(),
		Port:        80,
		Servicetype: "HTTP",
	}
	service2 := basic.Service{
		Name:        svcName2,
		Ip:          randomIP(),
		Port:        80,
		Servicetype: "HTTP",
	}

	_, err = client.AddResource(Service.Type(), svcName1, &service1)
	if err != nil {
		t.Error("Could not create service service1", err)
		log.Println("Cannot continue")
		return
	}
	_, err = client.AddResource(Service.Type(), svcName2, &service2)
	if err != nil {
		t.Error("Could not create service service2", err)
		log.Println("Cannot continue")
		return
	}

	binding1 := lb.Lbvserverservicebinding{
		Name:        lbName1,
		Servicename: svcName1,
	}
	binding2 := lb.Lbvserverservicebinding{
		Name:        lbName1,
		Servicename: svcName2,
	}

	err = client.BindResource(Lbvserver.Type(), lbName1, Service.Type(), svcName1, &binding1)
	if err != nil {
		t.Error("Could not bind service service1")
		log.Println("Cannot continue")
		return
	}

	err = client.BindResource(Lbvserver.Type(), lbName1, Service.Type(), svcName2, &binding2)
	if err != nil {
		t.Error("Could not bind service service2")
		log.Println("Cannot continue")
		return
	}
	rsrcs, err := client.FindAllBoundResources(Lbvserver.Type(), lbName1, Service.Type())
	if err != nil {
		t.Error("Did not find bound resources of type ", Service.Type())
	}
	if len(rsrcs) < 2 {
		t.Error("Found only ", len(rsrcs), " resources of type ", Service.Type(), " expected at least 2")
		log.Println("Cannot continue")
		return
	}

	found := 0
	for _, v := range rsrcs {
		name := v["servicename"].(string)
		if name == svcName1 || name == svcName2 {
			found = found + 1
		}
	}
	if found != 2 {
		t.Error("Did not find all bound services")
	}

}

func TestAction(t *testing.T) {
	svcGrpName := "test_sg_" + randomString(5)
	sg1 := basic.Servicegroup{
		Servicegroupname: svcGrpName,
		Servicetype:      "http",
	}

	_, err := client.AddResource(Servicegroup.Type(), svcGrpName, &sg1)
	if err != nil {
		t.Error("Could not add resource service group", err)
		log.Println("Cannot continue")
		return
	}
	createServer := basic.Server{
		Ipaddress: "192.168.1.101",
		Name:      "test-srvr",
	}

	_, err = client.AddResource(Server.Type(), "test-server", &createServer)
	if err != nil {
		t.Error("Could not add resource server", err)
		log.Println("Cannot continue")
		return
	}

	bindSvcGrpToServer := basic.Servicegroupservicegroupmemberbinding{
		Servicegroupname: svcGrpName,
		Servername:       "test-srvr",
		Port:             22,
	}

	_, err = client.AddResource(Servicegroup_servicegroupmember_binding.Type(), "test-svcgroup", &bindSvcGrpToServer)
	if err != nil {
		t.Error("Could not bind resource server", err)
		log.Println("Cannot continue")
		return
	}

	bindSvcGrpToServer2 := basic.Servicegroupservicegroupmemberbinding{
		Servicegroupname: svcGrpName,
		Ip:               "192.168.1.102",
		Port:             22,
	}

	_, err = client.AddResource(Servicegroup_servicegroupmember_binding.Type(), "test-svcgroup", &bindSvcGrpToServer2)
	if err != nil {
		t.Error("Could not bind resource server", err)
		log.Println("Cannot continue")
		return
	}
	sg2 := basic.Servicegroup{
		Servicegroupname: svcGrpName,
		Servername:       "test-srvr",
		Port:             22,
		Delay:            100,
		Graceful:         "YES",
	}

	err = client.ActOnResource(Servicegroup.Type(), &sg2, "disable")

	if err != nil {
		t.Error("Could not disable server", err)
		log.Println("Cannot continue")
		return
	}
	sg3 := basic.Servicegroup{
		Servicegroupname: svcGrpName,
		Servername:       "test-srvr",
		Port:             22,
	}

	err = client.ActOnResource(Servicegroup.Type(), &sg3, "enable")

	if err != nil {
		t.Error("Could not enable server", err)
		log.Println("Cannot continue")
		return
	}

	sg4 := basic.Servicegroup{
		Servicegroupname: svcGrpName,
		Newname:          svcGrpName + "-NEW",
	}

	err = client.ActOnResource(Servicegroup.Type(), &sg4, "rename")

	if err != nil {
		t.Error("Could not rename servicegroup", err)
		log.Println("Cannot continue")
		return
	}

}

func TestUpdateUnnamedResource(t *testing.T) {
	rnat := network.Rnat{
		Natip:   "172.17.0.2",
		Netmask: "255.255.240.0",
		Network: "192.168.16.0",
	}

	err := client.UpdateUnnamedResource(Rnat.Type(), &rnat)
	if err != nil {
		t.Error("Could not add Rnat", err)
		log.Println("Cannot continue")
		return
	}
}

func TestFindFilteredResource(t *testing.T) {
	rnat := network.Rnat{
		Natip:   "172.17.0.2",
		Netmask: "255.255.240.0",
		Network: "192.168.16.0",
	}

	err := client.UpdateUnnamedResource(Rnat.Type(), &rnat)
	if err != nil {
		t.Error("Could not add Rnat", err)
		log.Println("Cannot continue")
		return
	}
	d, err := client.FindFilteredResourceArray(Rnat.Type(), map[string]string{"network": "192.168.16.0", "netmask": "255.255.240.0", "natip": "172.17.0.2"})
	if err != nil {
		t.Error("Could not find Rnat", err)
		log.Println("Cannot continue")
		return
	}
	if len(d) != 1 {
		t.Error("Error finding Rnat", fmt.Errorf("Wrong number of RNAT discovered: %d", len(d)))
		return
	}
	rnat2 := d[0]
	if rnat2["natip"].(string) == "172.17.0.2" && rnat2["netmask"].(string) == "255.255.240.0" && rnat2["network"].(string) == "192.168.16.0" {
		return
	} else {
		t.Error("Error finding Rnat", fmt.Errorf("Discovered RNAT does not match"))
	}
}

// TestDesiredStateServicegroupAPI tests the servicegroup_servicegroupmemberlist_binding API
// which is used to bind multiple IP-only members to servicegroup in single Nitro call
func TestDesiredStateServicegroupAPI(t *testing.T) {
	svcGrpName := "test_sg_" + randomString(5)
	sg1 := basic.Servicegroup{
		Servicegroupname: svcGrpName,
		Servicetype:      "http",
		Autoscale:        "API",
	}

	_, err := client.AddResource(Servicegroup.Type(), svcGrpName, &sg1)
	if err != nil {
		t.Error("Could not add resource autoscale service group", err)
		log.Println("Cannot continue")
		return
	}

	ipmembers := []basic.Member{
		{
			Ip:   "1.1.1.1",
			Port: 80,
		},
		{
			Ip:   "2.2.2.2",
			Port: 80,
		},
		{
			Ip:   "3.3.3.3",
			Port: 80,
		},
	}

	bindSvcGrpToServer := basic.Servicegroupservicegroupmemberlistbinding{
		Servicegroupname: svcGrpName,
		Members:          ipmembers,
	}

	_, err = client.AddResource(Servicegroup_servicegroupmemberlist_binding.Type(), "test-svcgroup", &bindSvcGrpToServer)
	if err != nil {
		t.Error("Could not bind resource server", err)
		log.Println("Cannot continue")
		return
	}

}

func TestNullAction(t *testing.T) {
	reboot := ns.Reboot{
		Warm: true,
	}

	err := client.ActOnResource("reboot", &reboot, "")
	if err != nil {
		t.Error("Could not make null action reboot", err)
		log.Println("Cannot continue")
		return
	}

	// Add a timeout to wait for instance to be back online
	time.Sleep(60 * time.Second)
}

// TestTokenBasedAuth tests token-based authentication and tests if session-is is cleared in case of session-expiry
func TestTokenBasedAuth(t *testing.T) {
	var err error
	err = client.Login()
	if err != nil {
		t.Error("Login Failed")
		return
	}
	rndIP := randomIP()
	lbName := "test_lb_" + randomString(5)
	lb1 := lb.Lbvserver{
		Name:        lbName,
		Ipv46:       rndIP,
		Lbmethod:    "ROUNDROBIN",
		Servicetype: "HTTP",
		Port:        8000,
	}
	_, err = client.AddResource(Lbvserver.Type(), lbName, &lb1)
	if err != nil {
		t.Error("Could not add Lbvserver: ", err)
		log.Println("Not continuing test")
		return
	}

	rsrc, err := client.FindResource(Lbvserver.Type(), lbName)
	if err != nil {
		t.Error("Did not find resource of type ", err, Lbvserver.Type(), ":", lbName)
	} else {
		log.Println("LB-METHOD: ", rsrc["lbmethod"])
	}
	err = client.DeleteResource(Lbvserver.Type(), lbName)
	if err != nil {
		t.Error("Could not delete LB", lbName, err)
		log.Println("Cannot continue")
		return
	}
	err = client.Logout()
	if err != nil {
		t.Error("Logout Failed")
		return
	}

	// Test if session-id is cleared in case of session-expiry
	client.timeout = 10
	client.Login()
	time.Sleep(15 * time.Second)
	_, err = client.AddResource(Lbvserver.Type(), lbName, &lb1)
	if err != nil {
		if client.IsLoggedIn() {
			t.Error("Sessionid not cleared")
			return
		}
		log.Println("sessionid cleared because of session-expiry")
	} else {
		t.Error("Adding lbvserver should have failed because of session-expiry")
	}
}

func TestConstructQueryString(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	generateTestCase := func(findParams FindParams, expected string) func(t *testing.T) {
		return func(t *testing.T) {
			output := constructQueryString(&findParams)
			if output != expected {
				t.Log(buf.String())
				t.Logf("Expected output \"%s\"", expected)
				t.Logf("Actual output \"%s\"", output)
				t.Fail()
			}
		}
	}
	var argsMap, attrsMap, filterMap map[string]string
	var findParams FindParams

	argsMap = make(map[string]string)
	argsMap["hello"] = "bye"
	findParams = FindParams{
		ArgsMap: argsMap,
	}

	t.Run("CASE=1", generateTestCase(findParams, "?args=hello:bye"))

	argsMap["bye"] = "hello"
	findParams = FindParams{
		ArgsMap: argsMap,
	}
	t.Run("CASE=2", generateTestCase(findParams, "?args=bye:hello,hello:bye"))

	attrsMap = make(map[string]string)
	attrsMap["bye"] = "hello"
	findParams = FindParams{
		AttrsMap: attrsMap,
	}
	t.Run("CASE=3", generateTestCase(findParams, "?attrs=bye:hello"))

	attrsMap["hello"] = "bye"

	t.Run("CASE=4", generateTestCase(findParams, "?attrs=bye:hello,hello:bye"))

	filterMap = make(map[string]string)
	filterMap["bye"] = "hello"
	findParams = FindParams{
		FilterMap: filterMap,
	}
	t.Run("CASE=5", generateTestCase(findParams, "?filter=bye:hello"))

	filterMap["hello"] = "bye"

	t.Run("CASE=6", generateTestCase(findParams, "?filter=bye:hello,hello:bye"))

	filterMap = make(map[string]string)
	attrsMap = make(map[string]string)
	argsMap = make(map[string]string)

	filterMap["bye"] = "hello"
	attrsMap["bye"] = "hello"
	argsMap["bye"] = "hello"

	findParams = FindParams{
		FilterMap: filterMap,
		ArgsMap:   argsMap,
		AttrsMap:  attrsMap,
	}

	t.Run("CASE=7", generateTestCase(findParams, "?args=bye:hello&filter=bye:hello&attrs=bye:hello"))

	filterMap["hello"] = "bye"
	attrsMap["hello"] = "bye"
	argsMap["hello"] = "bye"

	expected := "?args=bye:hello,hello:bye&filter=bye:hello,hello:bye&attrs=bye:hello,hello:bye"
	t.Run("CASE=8", generateTestCase(findParams, expected))
}

func TestConstructUrlPathString(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	generateTestCase := func(findParams FindParams, expected string) func(t *testing.T) {
		return func(t *testing.T) {
			output := constructUrlPathString(&findParams)
			if output != expected {
				t.Log(buf.String())
				t.Logf("Expected output \"%s\"", expected)
				t.Logf("Actual output \"%s\"", output)
				t.Fail()
			}
		}
	}

	var findParams FindParams

	findParams = FindParams{
		ResourceType: "resourcetype",
	}

	t.Run("CASE=1", generateTestCase(findParams, "resourcetype"))

	findParams = FindParams{
		ResourceName: "resourcename",
	}

	t.Run("CASE=2", generateTestCase(findParams, "resourcename"))

	findParams = FindParams{
		ResourceType: "resourcetype",
		ResourceName: "resourcename",
	}

	t.Run("CASE=3", generateTestCase(findParams, "resourcetype/resourcename"))
}

func TestFindResourceArrayWithParams(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	testCase1 := func(t *testing.T) {
		findParams := FindParams{
			ResourceType:             "lbvserver",
			ResourceName:             "definitelynotexists",
			ResourceMissingErrorCode: 258,
		}
		resource, err := client.FindResourceArrayWithParams(findParams)
		hasErrors := false
		if err != nil {
			hasErrors = true
			t.Logf("Error from NITRO request: %s", err.Error())
		}
		if len(resource) > 0 {
			hasErrors = true
			t.Logf("Resource array not empty")
		}
		if hasErrors {
			t.Log(buf.String())
			t.Fail()
		}
	}
	t.Run("CASE=1", testCase1)

	testCase2 := func(t *testing.T) {
		argsMap := make(map[string]string)
		argsMap["filename"] = "ns.conf"
		argsMap["filelocation"] = "%2Fnsconfig"
		findParams := FindParams{
			ResourceType: "systemfile",
			ArgsMap:      argsMap,
		}
		resource, err := client.FindResourceArrayWithParams(findParams)
		hasErrors := false
		if err != nil {
			hasErrors = true
			t.Logf("Error from NITRO request: %s", err.Error())
		}
		if len(resource) != 1 {
			hasErrors = true
			t.Logf("Resource array not exactly 1")
		}
		if hasErrors {
			t.Log(buf.String())
			t.Fail()
		}
	}

	t.Run("CASE=2", testCase2)

	testCase3 := func(t *testing.T) {
		argsMap := make(map[string]string)
		//argsMap["filename"] = "ns.conf"
		argsMap["filelocation"] = "%2Fnsconfig"
		findParams := FindParams{
			ResourceType: "systemfile",
			ArgsMap:      argsMap,
		}
		resource, err := client.FindResourceArrayWithParams(findParams)
		hasErrors := false
		if err != nil {
			hasErrors = true
			t.Logf("Error from NITRO request: %s", err.Error())
		}
		if len(resource) <= 1 {
			hasErrors = true
			t.Logf("Resource array len not > 1")
		}
		if hasErrors {
			t.Log(buf.String())
			t.Fail()
		}
	}

	t.Run("CASE=3", testCase3)

	testCase4 := func(t *testing.T) {
		t.Skipf("TODO: find a resource for which NITRO returns a map instead of an array")
	}
	t.Run("CASE=4", testCase4)
}
