package aaa

type Aaaparameter struct {
	Aaadloglevel               string      `json:"aaadloglevel,omitempty"`
	Aaadnatip                  string      `json:"aaadnatip,omitempty"`
	Aaasessionloglevel         string      `json:"aaasessionloglevel,omitempty"`
	Apitokencache              string      `json:"apitokencache,omitempty"`
	Builtin                    interface{} `json:"builtin,omitempty"`
	Defaultauthtype            string      `json:"defaultauthtype,omitempty"`
	Defaultcspheader           string      `json:"defaultcspheader,omitempty"`
	Dynaddr                    string      `json:"dynaddr,omitempty"`
	Enableenhancedauthfeedback string      `json:"enableenhancedauthfeedback,omitempty"`
	Enablesessionstickiness    string      `json:"enablesessionstickiness,omitempty"`
	Enablestaticpagecaching    string      `json:"enablestaticpagecaching,omitempty"`
	Failedlogintimeout         int         `json:"failedlogintimeout,omitempty"`
	Feature                    string      `json:"feature,omitempty"`
	Ftmode                     string      `json:"ftmode,omitempty"`
	Loginencryption            string      `json:"loginencryption,omitempty"`
	Maxaaausers                int         `json:"maxaaausers,omitempty"`
	Maxkbquestions             int         `json:"maxkbquestions,omitempty"`
	Maxloginattempts           int         `json:"maxloginattempts,omitempty"`
	Maxsamldeflatesize         int         `json:"maxsamldeflatesize,omitempty"`
	Persistentloginattempts    string      `json:"persistentloginattempts,omitempty"`
	Pwdexpirynotificationdays  int         `json:"pwdexpirynotificationdays,omitempty"`
	Samesite                   string      `json:"samesite,omitempty"`
	Tokenintrospectioninterval int         `json:"tokenintrospectioninterval,omitempty"`
}
