package dns

type Dnsrecordsstats struct {
	Clearstats                  string `json:"clearstats,omitempty"`
	Dnscurentries               int    `json:"dnscurentries,omitempty"`
	Dnscurrecords               int    `json:"dnscurrecords,omitempty"`
	Dnsrecordtype               string `json:"dnsrecordtype,omitempty"`
	Dnstotentries               int    `json:"dnstotentries,omitempty"`
	Dnstoterraliasex            int    `json:"dnstoterraliasex,omitempty"`
	Dnstoterrlimits             int    `json:"dnstoterrlimits,omitempty"`
	Dnstoterrnodomains          int    `json:"dnstoterrnodomains,omitempty"`
	Dnstoterrrespform           int    `json:"dnstoterrrespform,omitempty"`
	Dnstotqueryforexpiredrecord int    `json:"dnstotqueryforexpiredrecord,omitempty"`
	Dnstotrequests              int    `json:"dnstotrequests,omitempty"`
	Dnstotresponses             int    `json:"dnstotresponses,omitempty"`
	Dnstotupdates               int    `json:"dnstotupdates,omitempty"`
}
