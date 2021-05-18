package dos

type Dospolicystats struct {
	Clearstats                string  `json:"clearstats,omitempty"`
	Doscurrentcltdetectrate   float64 `json:"doscurrentcltdetectrate,omitempty"`
	Doscurrentqueuesize       int     `json:"doscurrentqueuesize,omitempty"`
	Doscurrentqueuesizerate   int     `json:"doscurrentqueuesizerate,omitempty"`
	Dosjsbytessentrate        int     `json:"dosjsbytessentrate,omitempty"`
	Dosjsrefusedrate          int     `json:"dosjsrefusedrate,omitempty"`
	Dosjssentrate             int     `json:"dosjssentrate,omitempty"`
	Dosnongetpostrequestsrate int     `json:"dosnongetpostrequestsrate,omitempty"`
	Dosphysicalserviceip      string  `json:"dosphysicalserviceip,omitempty"`
	Dosphysicalserviceport    int     `json:"dosphysicalserviceport,omitempty"`
	Dostotjsbytessent         int     `json:"dostotjsbytessent,omitempty"`
	Dostotjsrefused           int     `json:"dostotjsrefused,omitempty"`
	Dostotjssent              int     `json:"dostotjssent,omitempty"`
	Dostotnongetpostrequests  int     `json:"dostotnongetpostrequests,omitempty"`
	Dostotvalidclients        int     `json:"dostotvalidclients,omitempty"`
	Dosvalidclientsrate       int     `json:"dosvalidclientsrate,omitempty"`
	Name                      string  `json:"name,omitempty"`
}
