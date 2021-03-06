package authentication

type Authenticationldapaction struct {
	Alternateemailattr         string `json:"alternateemailattr,omitempty"`
	Attribute1                 string `json:"attribute1,omitempty"`
	Attribute10                string `json:"attribute10,omitempty"`
	Attribute11                string `json:"attribute11,omitempty"`
	Attribute12                string `json:"attribute12,omitempty"`
	Attribute13                string `json:"attribute13,omitempty"`
	Attribute14                string `json:"attribute14,omitempty"`
	Attribute15                string `json:"attribute15,omitempty"`
	Attribute16                string `json:"attribute16,omitempty"`
	Attribute2                 string `json:"attribute2,omitempty"`
	Attribute3                 string `json:"attribute3,omitempty"`
	Attribute4                 string `json:"attribute4,omitempty"`
	Attribute5                 string `json:"attribute5,omitempty"`
	Attribute6                 string `json:"attribute6,omitempty"`
	Attribute7                 string `json:"attribute7,omitempty"`
	Attribute8                 string `json:"attribute8,omitempty"`
	Attribute9                 string `json:"attribute9,omitempty"`
	Attributes                 string `json:"attributes,omitempty"`
	Authentication             string `json:"authentication,omitempty"`
	Authtimeout                int    `json:"authtimeout,omitempty"`
	Cloudattributes            string `json:"cloudattributes,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Email                      string `json:"email,omitempty"`
	Failure                    int    `json:"failure,omitempty"`
	Followreferrals            string `json:"followreferrals,omitempty"`
	Groupattrname              string `json:"groupattrname,omitempty"`
	Groupnameidentifier        string `json:"groupnameidentifier,omitempty"`
	Groupsearchattribute       string `json:"groupsearchattribute,omitempty"`
	Groupsearchfilter          string `json:"groupsearchfilter,omitempty"`
	Groupsearchsubattribute    string `json:"groupsearchsubattribute,omitempty"`
	Kbattribute                string `json:"kbattribute,omitempty"`
	Ldapbase                   string `json:"ldapbase,omitempty"`
	Ldapbinddn                 string `json:"ldapbinddn,omitempty"`
	Ldapbinddnpassword         string `json:"ldapbinddnpassword,omitempty"`
	Ldaphostname               string `json:"ldaphostname,omitempty"`
	Ldaploginname              string `json:"ldaploginname,omitempty"`
	Maxldapreferrals           int    `json:"maxldapreferrals,omitempty"`
	Maxnestinglevel            int    `json:"maxnestinglevel,omitempty"`
	Mssrvrecordlocation        string `json:"mssrvrecordlocation,omitempty"`
	Name                       string `json:"name,omitempty"`
	Nestedgroupextraction      string `json:"nestedgroupextraction,omitempty"`
	Otpsecret                  string `json:"otpsecret,omitempty"`
	Passwdchange               string `json:"passwdchange,omitempty"`
	Pushservice                string `json:"pushservice,omitempty"`
	Referraldnslookup          string `json:"referraldnslookup,omitempty"`
	Requireuser                string `json:"requireuser,omitempty"`
	Searchfilter               string `json:"searchfilter,omitempty"`
	Sectype                    string `json:"sectype,omitempty"`
	Serverip                   string `json:"serverip,omitempty"`
	Servername                 string `json:"servername,omitempty"`
	Serverport                 int    `json:"serverport,omitempty"`
	Sshpublickey               string `json:"sshpublickey,omitempty"`
	Ssonameattribute           string `json:"ssonameattribute,omitempty"`
	Subattributename           string `json:"subattributename,omitempty"`
	Success                    int    `json:"success,omitempty"`
	Svrtype                    string `json:"svrtype,omitempty"`
	Validateservercert         string `json:"validateservercert,omitempty"`
}
