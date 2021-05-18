package lsn

type Lsngroupstats struct {
	Clearstats                string      `json:"clearstats,omitempty"`
	Groupname                 string      `json:"groupname,omitempty"`
	Lsngrpcuricmpsessions     int         `json:"lsngrpcuricmpsessions,omitempty"`
	Lsngrpcuricmpsessionsrate int         `json:"lsngrpcuricmpsessionsrate,omitempty"`
	Lsngrpcursessions         int         `json:"lsngrpcursessions,omitempty"`
	Lsngrpcursessionsrate     int         `json:"lsngrpcursessionsrate,omitempty"`
	Lsngrpcursubscribers      int         `json:"lsngrpcursubscribers,omitempty"`
	Lsngrpcursubscribersrate  int         `json:"lsngrpcursubscribersrate,omitempty"`
	Lsngrpcurtcpsessions      int         `json:"lsngrpcurtcpsessions,omitempty"`
	Lsngrpcurtcpsessionsrate  int         `json:"lsngrpcurtcpsessionsrate,omitempty"`
	Lsngrpcurudpsessions      int         `json:"lsngrpcurudpsessions,omitempty"`
	Lsngrpcurudpsessionsrate  int         `json:"lsngrpcurudpsessionsrate,omitempty"`
	Lsngrpicmpdrppktsrate     int         `json:"lsngrpicmpdrppktsrate,omitempty"`
	Lsngrpicmptranslbytesrate int         `json:"lsngrpicmptranslbytesrate,omitempty"`
	Lsngrpicmptranslpktsrate  int         `json:"lsngrpicmptranslpktsrate,omitempty"`
	Lsngrptcpdrppktsrate      int         `json:"lsngrptcpdrppktsrate,omitempty"`
	Lsngrptcptranslbytesrate  int         `json:"lsngrptcptranslbytesrate,omitempty"`
	Lsngrptcptranslpktsrate   int         `json:"lsngrptcptranslpktsrate,omitempty"`
	Lsngrptoticmpdrppkts      int         `json:"lsngrptoticmpdrppkts,omitempty"`
	Lsngrptoticmptranslbytes  int         `json:"lsngrptoticmptranslbytes,omitempty"`
	Lsngrptoticmptranslpkts   int         `json:"lsngrptoticmptranslpkts,omitempty"`
	Lsngrptottcpdrppkts       int         `json:"lsngrptottcpdrppkts,omitempty"`
	Lsngrptottcptranslbytes   int         `json:"lsngrptottcptranslbytes,omitempty"`
	Lsngrptottcptranslpkts    int         `json:"lsngrptottcptranslpkts,omitempty"`
	Lsngrptottranslbytes      int         `json:"lsngrptottranslbytes,omitempty"`
	Lsngrptottranslpkts       int         `json:"lsngrptottranslpkts,omitempty"`
	Lsngrptotudpdrppkts       int         `json:"lsngrptotudpdrppkts,omitempty"`
	Lsngrptotudptranslbytes   int         `json:"lsngrptotudptranslbytes,omitempty"`
	Lsngrptotudptranslpkts    int         `json:"lsngrptotudptranslpkts,omitempty"`
	Lsngrptranslbytesrate     int         `json:"lsngrptranslbytesrate,omitempty"`
	Lsngrptranslpktsrate      int         `json:"lsngrptranslpktsrate,omitempty"`
	Lsngrpudpdrppktsrate      int         `json:"lsngrpudpdrppktsrate,omitempty"`
	Lsngrpudptranslbytesrate  int         `json:"lsngrpudptranslbytesrate,omitempty"`
	Lsngrpudptranslpktsrate   int         `json:"lsngrpudptranslpktsrate,omitempty"`
	Lsnpool                   interface{} `json:"lsnpool,omitempty"`
	Pcpserver                 interface{} `json:"pcpserver,omitempty"`
}
