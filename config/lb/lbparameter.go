package lb

type Lbparameter struct {
	Adccookieattributewarningmsg  string      `json:"adccookieattributewarningmsg,omitempty"`
	Allowboundsvcremoval          string      `json:"allowboundsvcremoval,omitempty"`
	Builtin                       interface{} `json:"builtin,omitempty"`
	Computedadccookieattribute    string      `json:"computedadccookieattribute,omitempty"`
	Consolidatedlconn             string      `json:"consolidatedlconn,omitempty"`
	Cookiepassphrase              string      `json:"cookiepassphrase,omitempty"`
	Dbsttl                        int         `json:"dbsttl,omitempty"`
	Dropmqttjumbomessage          string      `json:"dropmqttjumbomessage,omitempty"`
	Feature                       string      `json:"feature,omitempty"`
	Httponlycookieflag            string      `json:"httponlycookieflag,omitempty"`
	Literaladccookieattribute     string      `json:"literaladccookieattribute,omitempty"`
	Maxpipelinenat                int         `json:"maxpipelinenat,omitempty"`
	Monitorconnectionclose        string      `json:"monitorconnectionclose,omitempty"`
	Monitorskipmaxclient          string      `json:"monitorskipmaxclient,omitempty"`
	Preferdirectroute             string      `json:"preferdirectroute,omitempty"`
	Retainservicestate            string      `json:"retainservicestate,omitempty"`
	Sessionsthreshold             int         `json:"sessionsthreshold,omitempty"`
	Startuprrfactor               int         `json:"startuprrfactor,omitempty"`
	Storemqttclientidandusername  string      `json:"storemqttclientidandusername,omitempty"`
	Useencryptedpersistencecookie string      `json:"useencryptedpersistencecookie,omitempty"`
	Useportforhashlb              string      `json:"useportforhashlb,omitempty"`
	Usesecuredpersistencecookie   string      `json:"usesecuredpersistencecookie,omitempty"`
	Vserverspecificmac            string      `json:"vserverspecificmac,omitempty"`
}
