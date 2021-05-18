package lldp

type Lldpstats struct {
	Clearstats                  string `json:"clearstats,omitempty"`
	Ifnum                       string `json:"ifnum,omitempty"`
	Rxportbytesrate             int    `json:"rxportbytesrate,omitempty"`
	Rxportbytestotal            int    `json:"rxportbytestotal,omitempty"`
	Rxportframeserrors          int    `json:"rxportframeserrors,omitempty"`
	Rxportframeserrorsrate      int    `json:"rxportframeserrorsrate,omitempty"`
	Rxportframesrate            int    `json:"rxportframesrate,omitempty"`
	Rxportframestotal           int    `json:"rxportframestotal,omitempty"`
	Rxporttlvsdiscardedrate     int    `json:"rxporttlvsdiscardedrate,omitempty"`
	Rxporttlvsdiscardedtotal    int    `json:"rxporttlvsdiscardedtotal,omitempty"`
	Rxporttlvsunrecognizedrate  int    `json:"rxporttlvsunrecognizedrate,omitempty"`
	Rxporttlvsunrecognizedtotal int    `json:"rxporttlvsunrecognizedtotal,omitempty"`
	Txportbytesrate             int    `json:"txportbytesrate,omitempty"`
	Txportbytestotal            int    `json:"txportbytestotal,omitempty"`
	Txportframesrate            int    `json:"txportframesrate,omitempty"`
	Txportframestotal           int    `json:"txportframestotal,omitempty"`
}
