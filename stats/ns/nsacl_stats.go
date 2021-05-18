package ns

type Nsaclstats struct {
	Aclhitsrate          int    `json:"aclhitsrate,omitempty"`
	Aclmissesrate        int    `json:"aclmissesrate,omitempty"`
	Aclname              string `json:"aclname,omitempty"`
	Aclperhits           int    `json:"aclperhits,omitempty"`
	Aclperhitsrate       int    `json:"aclperhitsrate,omitempty"`
	Aclpktsallowedrate   int    `json:"aclpktsallowedrate,omitempty"`
	Aclpktsbridgedrate   int    `json:"aclpktsbridgedrate,omitempty"`
	Aclpktsdeniedrate    int    `json:"aclpktsdeniedrate,omitempty"`
	Aclpktsnatrate       int    `json:"aclpktsnatrate,omitempty"`
	Acltotcount          int    `json:"acltotcount,omitempty"`
	Acltothits           int    `json:"acltothits,omitempty"`
	Acltotmisses         int    `json:"acltotmisses,omitempty"`
	Acltotpktsallowed    int    `json:"acltotpktsallowed,omitempty"`
	Acltotpktsbridged    int    `json:"acltotpktsbridged,omitempty"`
	Acltotpktsdenied     int    `json:"acltotpktsdenied,omitempty"`
	Acltotpktsnat        int    `json:"acltotpktsnat,omitempty"`
	Clearstats           string `json:"clearstats,omitempty"`
	Dfdaclhitsrate       int    `json:"dfdaclhitsrate,omitempty"`
	Dfdaclmissesrate     int    `json:"dfdaclmissesrate,omitempty"`
	Dfdacltotcount       int    `json:"dfdacltotcount,omitempty"`
	Dfdacltothits        int    `json:"dfdacltothits,omitempty"`
	Dfdacltotmisses      int    `json:"dfdacltotmisses,omitempty"`
	Effectiveacltotcount int    `json:"effectiveacltotcount,omitempty"`
}
