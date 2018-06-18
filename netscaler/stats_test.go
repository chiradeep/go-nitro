package netscaler

import "testing"

func TestNitroClient_FindStat(t *testing.T) {
	_, err := client.FindAllStats("lbvserver")
	if err != nil {
		t.Error("Did not find statistics of type ", err, Lbvserver.Type())
	}
}
