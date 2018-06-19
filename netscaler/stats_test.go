package netscaler

import "testing"

func TestNitroClient_FindAllStats(t *testing.T) {
	_, err := client.FindAllStats("lbvserver")
	if err != nil {
		t.Error("Did not find statistics of type ", err, Lbvserver.Type())
	}
}

func TestNitroClient_FindStats(t *testing.T) {
	for _, resourceType := range []string{"lbvserver", "service", "gslbvserver", "gslbservice"} {
		rsrc, err := client.FindAllResources("i")
		if err != nil {
			// Ignore the erratic resource type
			continue
		}
		_, err = client.FindStat(resourceType, rsrc[0]["name"].(string))
		if err != nil {
			t.Fatal(err)
		}
	}
}
