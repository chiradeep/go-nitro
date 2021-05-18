package protocol

type Protocolicmpv6stats struct {
	Clearstats            string `json:"clearstats,omitempty"`
	Icmpv6badchecksums    int    `json:"icmpv6badchecksums,omitempty"`
	Icmpv6errna           int    `json:"icmpv6errna,omitempty"`
	Icmpv6errns           int    `json:"icmpv6errns,omitempty"`
	Icmpv6errra           int    `json:"icmpv6errra,omitempty"`
	Icmpv6rtthsld         int    `json:"icmpv6rtthsld,omitempty"`
	Icmpv6rxbytesrate     int    `json:"icmpv6rxbytesrate,omitempty"`
	Icmpv6rxechoreplyrate int    `json:"icmpv6rxechoreplyrate,omitempty"`
	Icmpv6rxechoreqrate   int    `json:"icmpv6rxechoreqrate,omitempty"`
	Icmpv6rxnarate        int    `json:"icmpv6rxnarate,omitempty"`
	Icmpv6rxnsrate        int    `json:"icmpv6rxnsrate,omitempty"`
	Icmpv6rxpktsrate      int    `json:"icmpv6rxpktsrate,omitempty"`
	Icmpv6rxrarate        int    `json:"icmpv6rxrarate,omitempty"`
	Icmpv6rxrsrate        int    `json:"icmpv6rxrsrate,omitempty"`
	Icmpv6totrxbytes      int    `json:"icmpv6totrxbytes,omitempty"`
	Icmpv6totrxechoreply  int    `json:"icmpv6totrxechoreply,omitempty"`
	Icmpv6totrxechoreq    int    `json:"icmpv6totrxechoreq,omitempty"`
	Icmpv6totrxna         int    `json:"icmpv6totrxna,omitempty"`
	Icmpv6totrxns         int    `json:"icmpv6totrxns,omitempty"`
	Icmpv6totrxpkts       int    `json:"icmpv6totrxpkts,omitempty"`
	Icmpv6totrxra         int    `json:"icmpv6totrxra,omitempty"`
	Icmpv6totrxrs         int    `json:"icmpv6totrxrs,omitempty"`
	Icmpv6tottxbytes      int    `json:"icmpv6tottxbytes,omitempty"`
	Icmpv6tottxechoreply  int    `json:"icmpv6tottxechoreply,omitempty"`
	Icmpv6tottxechoreq    int    `json:"icmpv6tottxechoreq,omitempty"`
	Icmpv6tottxna         int    `json:"icmpv6tottxna,omitempty"`
	Icmpv6tottxns         int    `json:"icmpv6tottxns,omitempty"`
	Icmpv6tottxpkts       int    `json:"icmpv6tottxpkts,omitempty"`
	Icmpv6tottxra         int    `json:"icmpv6tottxra,omitempty"`
	Icmpv6tottxrs         int    `json:"icmpv6tottxrs,omitempty"`
	Icmpv6txbytesrate     int    `json:"icmpv6txbytesrate,omitempty"`
	Icmpv6txechoreplyrate int    `json:"icmpv6txechoreplyrate,omitempty"`
	Icmpv6txechoreqrate   int    `json:"icmpv6txechoreqrate,omitempty"`
	Icmpv6txnarate        int    `json:"icmpv6txnarate,omitempty"`
	Icmpv6txnsrate        int    `json:"icmpv6txnsrate,omitempty"`
	Icmpv6txpktsrate      int    `json:"icmpv6txpktsrate,omitempty"`
	Icmpv6txrarate        int    `json:"icmpv6txrarate,omitempty"`
	Icmpv6txrsrate        int    `json:"icmpv6txrsrate,omitempty"`
	Icmpv6unspt           int    `json:"icmpv6unspt,omitempty"`
}
