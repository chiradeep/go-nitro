package gslb

type Gslbsitestats struct {
	Clearstats             string `json:"clearstats,omitempty"`
	Nwmetricexchange       string `json:"nwmetricexchange,omitempty"`
	Nwmetricmepstatus      string `json:"nwmetricmepstatus,omitempty"`
	Persexchange           string `json:"persexchange,omitempty"`
	Sitecurclntconnections int    `json:"sitecurclntconnections,omitempty"`
	Sitecursrvrconnections int    `json:"sitecursrvrconnections,omitempty"`
	Siteip                 string `json:"siteip,omitempty"`
	Siteipstr              string `json:"siteipstr,omitempty"`
	Sitemepstatus          string `json:"sitemepstatus,omitempty"`
	Sitemetricexchange     string `json:"sitemetricexchange,omitempty"`
	Sitemetricmepstatus    string `json:"sitemetricmepstatus,omitempty"`
	Sitename               string `json:"sitename,omitempty"`
	Sitepublicip           string `json:"sitepublicip,omitempty"`
	Sitepublicipstr        string `json:"sitepublicipstr,omitempty"`
	Siterequestbytesrate   int    `json:"siterequestbytesrate,omitempty"`
	Siterequestsrate       int    `json:"siterequestsrate,omitempty"`
	Siteresponsebytesrate  int    `json:"siteresponsebytesrate,omitempty"`
	Siteresponsesrate      int    `json:"siteresponsesrate,omitempty"`
	Sitetotalrequestbytes  int    `json:"sitetotalrequestbytes,omitempty"`
	Sitetotalrequests      int    `json:"sitetotalrequests,omitempty"`
	Sitetotalresponsebytes int    `json:"sitetotalresponsebytes,omitempty"`
	Sitetotalresponses     int    `json:"sitetotalresponses,omitempty"`
	Sitetype               string `json:"sitetype,omitempty"`
}
