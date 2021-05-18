package lsn

type Lsnstats struct {
	Clearstats             string `json:"clearstats,omitempty"`
	Lsncuricmpsessions     int    `json:"lsncuricmpsessions,omitempty"`
	Lsncuricmpsessionsrate int    `json:"lsncuricmpsessionsrate,omitempty"`
	Lsncursessions         int    `json:"lsncursessions,omitempty"`
	Lsncursessionsrate     int    `json:"lsncursessionsrate,omitempty"`
	Lsncursubscribers      int    `json:"lsncursubscribers,omitempty"`
	Lsncursubscribersrate  int    `json:"lsncursubscribersrate,omitempty"`
	Lsncurtcpsessions      int    `json:"lsncurtcpsessions,omitempty"`
	Lsncurtcpsessionsrate  int    `json:"lsncurtcpsessionsrate,omitempty"`
	Lsncurudpsessions      int    `json:"lsncurudpsessions,omitempty"`
	Lsncurudpsessionsrate  int    `json:"lsncurudpsessionsrate,omitempty"`
	Lsnicmpdrppktsrate     int    `json:"lsnicmpdrppktsrate,omitempty"`
	Lsnicmprxbytesrate     int    `json:"lsnicmprxbytesrate,omitempty"`
	Lsnicmprxpktsrate      int    `json:"lsnicmprxpktsrate,omitempty"`
	Lsnicmptxbytesrate     int    `json:"lsnicmptxbytesrate,omitempty"`
	Lsnicmptxpktsrate      int    `json:"lsnicmptxpktsrate,omitempty"`
	Lsntcpdrppktsrate      int    `json:"lsntcpdrppktsrate,omitempty"`
	Lsntcprxbytesrate      int    `json:"lsntcprxbytesrate,omitempty"`
	Lsntcprxpktsrate       int    `json:"lsntcprxpktsrate,omitempty"`
	Lsntcptxbytesrate      int    `json:"lsntcptxbytesrate,omitempty"`
	Lsntcptxpktsrate       int    `json:"lsntcptxpktsrate,omitempty"`
	Lsntoticmpdrppkts      int    `json:"lsntoticmpdrppkts,omitempty"`
	Lsntoticmprxbytes      int    `json:"lsntoticmprxbytes,omitempty"`
	Lsntoticmprxpkts       int    `json:"lsntoticmprxpkts,omitempty"`
	Lsntoticmptxbytes      int    `json:"lsntoticmptxbytes,omitempty"`
	Lsntoticmptxpkts       int    `json:"lsntoticmptxpkts,omitempty"`
	Lsntottcpdrppkts       int    `json:"lsntottcpdrppkts,omitempty"`
	Lsntottcprxbytes       int    `json:"lsntottcprxbytes,omitempty"`
	Lsntottcprxpkts        int    `json:"lsntottcprxpkts,omitempty"`
	Lsntottcptxbytes       int    `json:"lsntottcptxbytes,omitempty"`
	Lsntottcptxpkts        int    `json:"lsntottcptxpkts,omitempty"`
	Lsntotudpdrppkts       int    `json:"lsntotudpdrppkts,omitempty"`
	Lsntotudprxbytes       int    `json:"lsntotudprxbytes,omitempty"`
	Lsntotudprxpkts        int    `json:"lsntotudprxpkts,omitempty"`
	Lsntotudptxbytes       int    `json:"lsntotudptxbytes,omitempty"`
	Lsntotudptxpkts        int    `json:"lsntotudptxpkts,omitempty"`
	Lsnudpdrppktsrate      int    `json:"lsnudpdrppktsrate,omitempty"`
	Lsnudprxbytesrate      int    `json:"lsnudprxbytesrate,omitempty"`
	Lsnudprxpktsrate       int    `json:"lsnudprxpktsrate,omitempty"`
	Lsnudptxbytesrate      int    `json:"lsnudptxbytesrate,omitempty"`
	Lsnudptxpktsrate       int    `json:"lsnudptxpktsrate,omitempty"`
}
