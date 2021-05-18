package aaa

type Aaastats struct {
	Aaaauthfail                int    `json:"aaaauthfail,omitempty"`
	Aaaauthfailrate            int    `json:"aaaauthfailrate,omitempty"`
	Aaaauthnonhttpfail         int    `json:"aaaauthnonhttpfail,omitempty"`
	Aaaauthnonhttpfailrate     int    `json:"aaaauthnonhttpfailrate,omitempty"`
	Aaaauthnonhttpsuccess      int    `json:"aaaauthnonhttpsuccess,omitempty"`
	Aaaauthnonhttpsuccessrate  int    `json:"aaaauthnonhttpsuccessrate,omitempty"`
	Aaaauthonlyhttpfail        int    `json:"aaaauthonlyhttpfail,omitempty"`
	Aaaauthonlyhttpfailrate    int    `json:"aaaauthonlyhttpfailrate,omitempty"`
	Aaaauthonlyhttpsuccess     int    `json:"aaaauthonlyhttpsuccess,omitempty"`
	Aaaauthonlyhttpsuccessrate int    `json:"aaaauthonlyhttpsuccessrate,omitempty"`
	Aaaauthsuccess             int    `json:"aaaauthsuccess,omitempty"`
	Aaaauthsuccessrate         int    `json:"aaaauthsuccessrate,omitempty"`
	Aaacuricaconn              int    `json:"aaacuricaconn,omitempty"`
	Aaacuricaconnrate          int    `json:"aaacuricaconnrate,omitempty"`
	Aaacuricaonlyconn          int    `json:"aaacuricaonlyconn,omitempty"`
	Aaacuricaonlyconnrate      int    `json:"aaacuricaonlyconnrate,omitempty"`
	Aaacuricasessions          int    `json:"aaacuricasessions,omitempty"`
	Aaacuricasessionsrate      int    `json:"aaacuricasessionsrate,omitempty"`
	Aaacursessions             int    `json:"aaacursessions,omitempty"`
	Aaacursessionsrate         int    `json:"aaacursessionsrate,omitempty"`
	Aaacurtmsessions           int    `json:"aaacurtmsessions,omitempty"`
	Aaacurtmsessionsrate       int    `json:"aaacurtmsessionsrate,omitempty"`
	Aaasessionsrate            int    `json:"aaasessionsrate,omitempty"`
	Aaasessiontimeoutrate      int    `json:"aaasessiontimeoutrate,omitempty"`
	Aaatmsessionsrate          int    `json:"aaatmsessionsrate,omitempty"`
	Aaatotsessions             int    `json:"aaatotsessions,omitempty"`
	Aaatotsessiontimeout       int    `json:"aaatotsessiontimeout,omitempty"`
	Aaatottmsessions           int    `json:"aaatottmsessions,omitempty"`
	Clearstats                 string `json:"clearstats,omitempty"`
}
