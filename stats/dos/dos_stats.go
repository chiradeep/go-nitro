package dos

type Dosstats struct {
	Clearstats                string `json:"clearstats,omitempty"`
	Dosconditiontriggeredrate int    `json:"dosconditiontriggeredrate,omitempty"`
	Dosdospriorityclientsrate int    `json:"dosdospriorityclientsrate,omitempty"`
	Dostotconditiontriggered  int    `json:"dostotconditiontriggered,omitempty"`
	Dostotdospriorityclients  int    `json:"dostotdospriorityclients,omitempty"`
	Dostotvalidcookies        int    `json:"dostotvalidcookies,omitempty"`
	Dosvalidcookiesrate       int    `json:"dosvalidcookiesrate,omitempty"`
}
