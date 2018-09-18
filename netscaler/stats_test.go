package netscaler

import "testing"

func TestNitroClient_FindAllStats(t *testing.T) {
	_, err := client.FindAllStats(Lbvserver.Type())
	if err != nil {
		t.Error("Did not find statistics of type ", err, Lbvserver.Type())
	}
}

func TestNitroClient_FindStats(t *testing.T) {
	for _, resourceType := range []string{Lbvserver.Type(), Service.Type(), Gslbvserver.Type(), Gslbservice.Type()} {
		rsrc, err := client.FindAllResources(resourceType)
		if err != nil {
			// Ignore the erratic resource type
			continue
		}
		for _, availableItem := range rsrc {
			_, err = client.FindStat(resourceType, availableItem["name"].(string))
			if err != nil {
				t.Fatal(err)
			}
			// only check one
			break
		}
	}
}
