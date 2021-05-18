package appqoe

type Appqoeaction struct {
	Altcontentpath    string `json:"altcontentpath,omitempty"`
	Altcontentsvcname string `json:"altcontentsvcname,omitempty"`
	Customfile        string `json:"customfile,omitempty"`
	Delay             int    `json:"delay,omitempty"`
	Dosaction         string `json:"dosaction,omitempty"`
	Dostrigexpression string `json:"dostrigexpression,omitempty"`
	Hits              int    `json:"hits,omitempty"`
	Maxconn           int    `json:"maxconn,omitempty"`
	Name              string `json:"name,omitempty"`
	Numretries        int    `json:"numretries,omitempty"`
	Polqdepth         int    `json:"polqdepth,omitempty"`
	Priority          string `json:"priority,omitempty"`
	Priqdepth         int    `json:"priqdepth,omitempty"`
	Respondwith       string `json:"respondwith,omitempty"`
	Retryonreset      string `json:"retryonreset,omitempty"`
	Retryontimeout    int    `json:"retryontimeout,omitempty"`
	Tcpprofile        string `json:"tcpprofile,omitempty"`
}
