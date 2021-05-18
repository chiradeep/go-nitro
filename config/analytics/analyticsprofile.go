package analytics

type Analyticsprofile struct {
	Allhttpheaders             string `json:"allhttpheaders,omitempty"`
	Auditlogs                  string `json:"auditlogs,omitempty"`
	Collectors                 string `json:"collectors,omitempty"`
	Cqareporting               string `json:"cqareporting,omitempty"`
	Events                     string `json:"events,omitempty"`
	Grpcstatus                 string `json:"grpcstatus,omitempty"`
	Httpauthentication         string `json:"httpauthentication,omitempty"`
	Httpclientsidemeasurements string `json:"httpclientsidemeasurements,omitempty"`
	Httpcontenttype            string `json:"httpcontenttype,omitempty"`
	Httpcookie                 string `json:"httpcookie,omitempty"`
	Httpdomainname             string `json:"httpdomainname,omitempty"`
	Httphost                   string `json:"httphost,omitempty"`
	Httplocation               string `json:"httplocation,omitempty"`
	Httpmethod                 string `json:"httpmethod,omitempty"`
	Httppagetracking           string `json:"httppagetracking,omitempty"`
	Httpreferer                string `json:"httpreferer,omitempty"`
	Httpsetcookie              string `json:"httpsetcookie,omitempty"`
	Httpsetcookie2             string `json:"httpsetcookie2,omitempty"`
	Httpurl                    string `json:"httpurl,omitempty"`
	Httpurlquery               string `json:"httpurlquery,omitempty"`
	Httpuseragent              string `json:"httpuseragent,omitempty"`
	Httpvia                    string `json:"httpvia,omitempty"`
	Httpxforwardedforheader    string `json:"httpxforwardedforheader,omitempty"`
	Integratedcache            string `json:"integratedcache,omitempty"`
	Metrics                    string `json:"metrics,omitempty"`
	Name                       string `json:"name,omitempty"`
	Outputmode                 string `json:"outputmode,omitempty"`
	Refcnt                     int    `json:"refcnt,omitempty"`
	Tcpburstreporting          string `json:"tcpburstreporting,omitempty"`
	Type                       string `json:"type,omitempty"`
	Urlcategory                string `json:"urlcategory,omitempty"`
}
