package audit

type Auditstats struct {
	Auditcontextnotfound                int    `json:"auditcontextnotfound,omitempty"`
	Auditcontextnotfoundrate            int    `json:"auditcontextnotfoundrate,omitempty"`
	Auditlog32errsyslogallocnsbfail     int    `json:"auditlog32errsyslogallocnsbfail,omitempty"`
	Auditlog32errsyslogallocnsbfailrate int    `json:"auditlog32errsyslogallocnsbfailrate,omitempty"`
	Auditmemallocfail                   int    `json:"auditmemallocfail,omitempty"`
	Auditmemallocfailrate               int    `json:"auditmemallocfailrate,omitempty"`
	Auditnsballocfail                   int    `json:"auditnsballocfail,omitempty"`
	Auditnsballocfailrate               int    `json:"auditnsballocfailrate,omitempty"`
	Auditportallocfail                  int    `json:"auditportallocfail,omitempty"`
	Auditportallocfailrate              int    `json:"auditportallocfailrate,omitempty"`
	Auditsyslogmsggen                   int    `json:"auditsyslogmsggen,omitempty"`
	Auditsyslogmsggenrate               int    `json:"auditsyslogmsggenrate,omitempty"`
	Auditsyslogmsgsent                  int    `json:"auditsyslogmsgsent,omitempty"`
	Auditsyslogmsgsentrate              int    `json:"auditsyslogmsgsentrate,omitempty"`
	Auditsyslogmsgsenttcp               int    `json:"auditsyslogmsgsenttcp,omitempty"`
	Auditsyslogmsgsenttcprate           int    `json:"auditsyslogmsgsenttcprate,omitempty"`
	Clearstats                          string `json:"clearstats,omitempty"`
	Clientconnfail                      int    `json:"clientconnfail,omitempty"`
	Clientconnfailrate                  int    `json:"clientconnfailrate,omitempty"`
	Flushcmdcnt                         int    `json:"flushcmdcnt,omitempty"`
	Flushcmdcntrate                     int    `json:"flushcmdcntrate,omitempty"`
	Logsdropped                         int    `json:"logsdropped,omitempty"`
	Logsdroppedrate                     int    `json:"logsdroppedrate,omitempty"`
	Logsdroppedtxminnsbs                int    `json:"logsdroppedtxminnsbs,omitempty"`
	Logsdroppedtxminnsbsrate            int    `json:"logsdroppedtxminnsbsrate,omitempty"`
	Logunsentlbsys                      int    `json:"logunsentlbsys,omitempty"`
	Logunsentlbsysrate                  int    `json:"logunsentlbsysrate,omitempty"`
	Nsbchainallocfail                   int    `json:"nsbchainallocfail,omitempty"`
	Nsbchainallocfailrate               int    `json:"nsbchainallocfailrate,omitempty"`
	Systcpconnfail                      int    `json:"systcpconnfail,omitempty"`
	Systcpconnfailrate                  int    `json:"systcpconnfailrate,omitempty"`
}
