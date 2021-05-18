package appflow

type Appflowstats struct {
	Appflowflowsrate        int    `json:"appflowflowsrate,omitempty"`
	Appflowignoctetsrate    int    `json:"appflowignoctetsrate,omitempty"`
	Appflowignpacketssrate  int    `json:"appflowignpacketssrate,omitempty"`
	Appflownotxflowsrate    int    `json:"appflownotxflowsrate,omitempty"`
	Appflownotxoctetsrate   int    `json:"appflownotxoctetsrate,omitempty"`
	Appflownotxpacketsrate  int    `json:"appflownotxpacketsrate,omitempty"`
	Appflowtotalflows       int    `json:"appflowtotalflows,omitempty"`
	Appflowtotalignoctets   int    `json:"appflowtotalignoctets,omitempty"`
	Appflowtotalignpacketss int    `json:"appflowtotalignpacketss,omitempty"`
	Appflowtotalnotxflows   int    `json:"appflowtotalnotxflows,omitempty"`
	Appflowtotalnotxoctets  int    `json:"appflowtotalnotxoctets,omitempty"`
	Appflowtotalnotxpackets int    `json:"appflowtotalnotxpackets,omitempty"`
	Appflowtotaltxmessagess int    `json:"appflowtotaltxmessagess,omitempty"`
	Appflowtotaltxoctets    int    `json:"appflowtotaltxoctets,omitempty"`
	Appflowtxmessagessrate  int    `json:"appflowtxmessagessrate,omitempty"`
	Appflowtxoctetsrate     int    `json:"appflowtxoctetsrate,omitempty"`
	Clearstats              string `json:"clearstats,omitempty"`
}
