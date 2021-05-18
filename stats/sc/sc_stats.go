package sc

type Scstats struct {
	Clearstats             string `json:"clearstats,omitempty"`
	Scaltconturls          int    `json:"scaltconturls,omitempty"`
	Scaltconturlsrate      int    `json:"scaltconturlsrate,omitempty"`
	Sccondtriggeredrate    int    `json:"sccondtriggeredrate,omitempty"`
	Scfaultycookies        int    `json:"scfaultycookies,omitempty"`
	Scfaultycookiesrate    int    `json:"scfaultycookiesrate,omitempty"`
	Scpolicyurlhits        int    `json:"scpolicyurlhits,omitempty"`
	Scpolicyurlhitsrate    int    `json:"scpolicyurlhitsrate,omitempty"`
	Scpopups               int    `json:"scpopups,omitempty"`
	Scpopupsrate           int    `json:"scpopupsrate,omitempty"`
	Scpostreqs             int    `json:"scpostreqs,omitempty"`
	Scpostreqsrate         int    `json:"scpostreqsrate,omitempty"`
	Screissuedrequestsrate int    `json:"screissuedrequestsrate,omitempty"`
	Scresetstats           int    `json:"scresetstats,omitempty"`
	Scresetstatsrate       int    `json:"scresetstatsrate,omitempty"`
	Scsessionreqs          int    `json:"scsessionreqs,omitempty"`
	Scsessionreqsrate      int    `json:"scsessionreqsrate,omitempty"`
	Sctotcondtriggered     int    `json:"sctotcondtriggered,omitempty"`
	Sctotreissuedrequests  int    `json:"sctotreissuedrequests,omitempty"`
	Scunsupbrow            int    `json:"scunsupbrow,omitempty"`
	Scunsupbrowrate        int    `json:"scunsupbrowrate,omitempty"`
}
