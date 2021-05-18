package lb

type Lbprofile struct {
	Adccookieattributewarningmsg  string `json:"adccookieattributewarningmsg,omitempty"`
	Computedadccookieattribute    string `json:"computedadccookieattribute,omitempty"`
	Cookiepassphrase              string `json:"cookiepassphrase,omitempty"`
	Dbslb                         string `json:"dbslb,omitempty"`
	Httponlycookieflag            string `json:"httponlycookieflag,omitempty"`
	Lbprofilename                 string `json:"lbprofilename,omitempty"`
	Literaladccookieattribute     string `json:"literaladccookieattribute,omitempty"`
	Processlocal                  string `json:"processlocal,omitempty"`
	Storemqttclientidandusername  string `json:"storemqttclientidandusername,omitempty"`
	Useencryptedpersistencecookie string `json:"useencryptedpersistencecookie,omitempty"`
	Usesecuredpersistencecookie   string `json:"usesecuredpersistencecookie,omitempty"`
	Vsvrcount                     int    `json:"vsvrcount,omitempty"`
}
