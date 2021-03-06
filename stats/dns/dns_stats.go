package dns

type Dnsstats struct {
	Clearstats                     string `json:"clearstats,omitempty"`
	Dns64aaaabypassrate            int    `json:"dns64aaaabypassrate,omitempty"`
	Dns64activepolicies            int    `json:"dns64activepolicies,omitempty"`
	Dns64answersrate               int    `json:"dns64answersrate,omitempty"`
	Dns64gslbanswersrate           int    `json:"dns64gslbanswersrate,omitempty"`
	Dns64gslbqueriesrate           int    `json:"dns64gslbqueriesrate,omitempty"`
	Dns64nodataresprate            int    `json:"dns64nodataresprate,omitempty"`
	Dns64queriesrate               int    `json:"dns64queriesrate,omitempty"`
	Dns64responsesrate             int    `json:"dns64responsesrate,omitempty"`
	Dns64rwanswersrate             int    `json:"dns64rwanswersrate,omitempty"`
	Dns64svraqueriesrate           int    `json:"dns64svraqueriesrate,omitempty"`
	Dns64tcanswersrate             int    `json:"dns64tcanswersrate,omitempty"`
	Dns64tcpqueriesrate            int    `json:"dns64tcpqueriesrate,omitempty"`
	Dns64totaaaabypass             int    `json:"dns64totaaaabypass,omitempty"`
	Dns64totanswers                int    `json:"dns64totanswers,omitempty"`
	Dns64totgslbanswers            int    `json:"dns64totgslbanswers,omitempty"`
	Dns64totgslbqueries            int    `json:"dns64totgslbqueries,omitempty"`
	Dns64totnodataresp             int    `json:"dns64totnodataresp,omitempty"`
	Dns64totqueries                int    `json:"dns64totqueries,omitempty"`
	Dns64totresponses              int    `json:"dns64totresponses,omitempty"`
	Dns64totrwanswers              int    `json:"dns64totrwanswers,omitempty"`
	Dns64totsvraqueries            int    `json:"dns64totsvraqueries,omitempty"`
	Dns64tottcanswers              int    `json:"dns64tottcanswers,omitempty"`
	Dns64tottcpqueries             int    `json:"dns64tottcpqueries,omitempty"`
	Dnsaaaarecqueriesrate          int    `json:"dnsaaaarecqueriesrate,omitempty"`
	Dnsaaaaresponserate            int    `json:"dnsaaaaresponserate,omitempty"`
	Dnsanswersrate                 int    `json:"dnsanswersrate,omitempty"`
	Dnsanyqueriesrate              int    `json:"dnsanyqueriesrate,omitempty"`
	Dnsanyresponserate             int    `json:"dnsanyresponserate,omitempty"`
	Dnsarecqueriesrate             int    `json:"dnsarecqueriesrate,omitempty"`
	Dnsaresponserate               int    `json:"dnsaresponserate,omitempty"`
	Dnscnamerecqueriesrate         int    `json:"dnscnamerecqueriesrate,omitempty"`
	Dnscnameresponserate           int    `json:"dnscnameresponserate,omitempty"`
	Dnscuraaaarecord               int    `json:"dnscuraaaarecord,omitempty"`
	Dnscurarecord                  int    `json:"dnscurarecord,omitempty"`
	Dnscurauthentries              int    `json:"dnscurauthentries,omitempty"`
	Dnscurcachesize                int    `json:"dnscurcachesize,omitempty"`
	Dnscurcnamerecord              int    `json:"dnscurcnamerecord,omitempty"`
	Dnscurmxrecord                 int    `json:"dnscurmxrecord,omitempty"`
	Dnscurnegcachesize             int    `json:"dnscurnegcachesize,omitempty"`
	Dnscurnoauthentries            int    `json:"dnscurnoauthentries,omitempty"`
	Dnscurnsrecord                 int    `json:"dnscurnsrecord,omitempty"`
	Dnscurptrrecord                int    `json:"dnscurptrrecord,omitempty"`
	Dnscursoarecord                int    `json:"dnscursoarecord,omitempty"`
	Dnscursrvrecord                int    `json:"dnscursrvrecord,omitempty"`
	Dnserrnullattack               int    `json:"dnserrnullattack,omitempty"`
	Dnsjumboanswersrate            int    `json:"dnsjumboanswersrate,omitempty"`
	Dnsjumboqueriesrate            int    `json:"dnsjumboqueriesrate,omitempty"`
	Dnsjumboserverresponsesrate    int    `json:"dnsjumboserverresponsesrate,omitempty"`
	Dnsmxrecqueriesrate            int    `json:"dnsmxrecqueriesrate,omitempty"`
	Dnsmxresponserate              int    `json:"dnsmxresponserate,omitempty"`
	Dnsnsrecqueriesrate            int    `json:"dnsnsrecqueriesrate,omitempty"`
	Dnsnsresponserate              int    `json:"dnsnsresponserate,omitempty"`
	Dnsptrrecqueriesrate           int    `json:"dnsptrrecqueriesrate,omitempty"`
	Dnsptrresponserate             int    `json:"dnsptrresponserate,omitempty"`
	Dnsqueriesrate                 int    `json:"dnsqueriesrate,omitempty"`
	Dnsserverqueryrate             int    `json:"dnsserverqueryrate,omitempty"`
	Dnsserverresponserate          int    `json:"dnsserverresponserate,omitempty"`
	Dnssoarecqueriesrate           int    `json:"dnssoarecqueriesrate,omitempty"`
	Dnssoaresponserate             int    `json:"dnssoaresponserate,omitempty"`
	Dnssrvrecqueriesrate           int    `json:"dnssrvrecqueriesrate,omitempty"`
	Dnssrvresponserate             int    `json:"dnssrvresponserate,omitempty"`
	Dnstotaaaarecfailed            int    `json:"dnstotaaaarecfailed,omitempty"`
	Dnstotaaaarecqueries           int    `json:"dnstotaaaarecqueries,omitempty"`
	Dnstotaaaarecupdate            int    `json:"dnstotaaaarecupdate,omitempty"`
	Dnstotaaaaresponse             int    `json:"dnstotaaaaresponse,omitempty"`
	Dnstotanswers                  int    `json:"dnstotanswers,omitempty"`
	Dnstotanyqueries               int    `json:"dnstotanyqueries,omitempty"`
	Dnstotanyrecfailed             int    `json:"dnstotanyrecfailed,omitempty"`
	Dnstotanyresponse              int    `json:"dnstotanyresponse,omitempty"`
	Dnstotarecfailed               int    `json:"dnstotarecfailed,omitempty"`
	Dnstotarecqueries              int    `json:"dnstotarecqueries,omitempty"`
	Dnstotarecupdate               int    `json:"dnstotarecupdate,omitempty"`
	Dnstotaresponse                int    `json:"dnstotaresponse,omitempty"`
	Dnstotauthans                  int    `json:"dnstotauthans,omitempty"`
	Dnstotauthnonames              int    `json:"dnstotauthnonames,omitempty"`
	Dnstotcacheentriesflush        int    `json:"dnstotcacheentriesflush,omitempty"`
	Dnstotcacheflush               int    `json:"dnstotcacheflush,omitempty"`
	Dnstotcnamerecfailed           int    `json:"dnstotcnamerecfailed,omitempty"`
	Dnstotcnamerecqueries          int    `json:"dnstotcnamerecqueries,omitempty"`
	Dnstotcnamerecupdate           int    `json:"dnstotcnamerecupdate,omitempty"`
	Dnstotcnameresponse            int    `json:"dnstotcnameresponse,omitempty"`
	Dnstotinvalidqueryformat       int    `json:"dnstotinvalidqueryformat,omitempty"`
	Dnstotjumboanswers             int    `json:"dnstotjumboanswers,omitempty"`
	Dnstotjumboqueries             int    `json:"dnstotjumboqueries,omitempty"`
	Dnstotjumboserverresponses     int    `json:"dnstotjumboserverresponses,omitempty"`
	Dnstotmultiquery               int    `json:"dnstotmultiquery,omitempty"`
	Dnstotmultiquerydisableerror   int    `json:"dnstotmultiquerydisableerror,omitempty"`
	Dnstotmxrecfailed              int    `json:"dnstotmxrecfailed,omitempty"`
	Dnstotmxrecqueries             int    `json:"dnstotmxrecqueries,omitempty"`
	Dnstotmxrecupdate              int    `json:"dnstotmxrecupdate,omitempty"`
	Dnstotmxresponse               int    `json:"dnstotmxresponse,omitempty"`
	Dnstotnodataresps              int    `json:"dnstotnodataresps,omitempty"`
	Dnstotnonauthnodatas           int    `json:"dnstotnonauthnodatas,omitempty"`
	Dnstotnsrecfailed              int    `json:"dnstotnsrecfailed,omitempty"`
	Dnstotnsrecqueries             int    `json:"dnstotnsrecqueries,omitempty"`
	Dnstotnsrecupdate              int    `json:"dnstotnsrecupdate,omitempty"`
	Dnstotnsresponse               int    `json:"dnstotnsresponse,omitempty"`
	Dnstotothererrors              int    `json:"dnstotothererrors,omitempty"`
	Dnstotptrrecfailed             int    `json:"dnstotptrrecfailed,omitempty"`
	Dnstotptrrecqueries            int    `json:"dnstotptrrecqueries,omitempty"`
	Dnstotptrrecupdate             int    `json:"dnstotptrrecupdate,omitempty"`
	Dnstotptrresponse              int    `json:"dnstotptrresponse,omitempty"`
	Dnstotqueries                  int    `json:"dnstotqueries,omitempty"`
	Dnstotrecupdate                int    `json:"dnstotrecupdate,omitempty"`
	Dnstotreqrefusals              int    `json:"dnstotreqrefusals,omitempty"`
	Dnstotresponsebadlen           int    `json:"dnstotresponsebadlen,omitempty"`
	Dnstotserverquery              int    `json:"dnstotserverquery,omitempty"`
	Dnstotserverresponse           int    `json:"dnstotserverresponse,omitempty"`
	Dnstotsoarecfailed             int    `json:"dnstotsoarecfailed,omitempty"`
	Dnstotsoarecqueries            int    `json:"dnstotsoarecqueries,omitempty"`
	Dnstotsoarecupdate             int    `json:"dnstotsoarecupdate,omitempty"`
	Dnstotsoaresponse              int    `json:"dnstotsoaresponse,omitempty"`
	Dnstotsrvrecfailed             int    `json:"dnstotsrvrecfailed,omitempty"`
	Dnstotsrvrecqueries            int    `json:"dnstotsrvrecqueries,omitempty"`
	Dnstotsrvrecupdate             int    `json:"dnstotsrvrecupdate,omitempty"`
	Dnstotsrvresponse              int    `json:"dnstotsrvresponse,omitempty"`
	Dnstotstrayanswer              int    `json:"dnstotstrayanswer,omitempty"`
	Dnstotunsupportedqueries       int    `json:"dnstotunsupportedqueries,omitempty"`
	Dnstotunsupportedqueryclass    int    `json:"dnstotunsupportedqueryclass,omitempty"`
	Dnstotunsupportedresponseclass int    `json:"dnstotunsupportedresponseclass,omitempty"`
	Dnstotunsupportedresponsetype  int    `json:"dnstotunsupportedresponsetype,omitempty"`
}
