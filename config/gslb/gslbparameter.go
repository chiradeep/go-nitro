package gslb

type Gslbparameter struct {
	Automaticconfigsync   string      `json:"automaticconfigsync,omitempty"`
	Builtin               interface{} `json:"builtin,omitempty"`
	Dropldnsreq           string      `json:"dropldnsreq,omitempty"`
	Feature               string      `json:"feature,omitempty"`
	Flags                 int         `json:"flags,omitempty"`
	Gslbsvcstatedelaytime int         `json:"gslbsvcstatedelaytime,omitempty"`
	Ldnsentrytimeout      int         `json:"ldnsentrytimeout,omitempty"`
	Ldnsmask              string      `json:"ldnsmask,omitempty"`
	Ldnsprobeorder        interface{} `json:"ldnsprobeorder,omitempty"`
	Mepkeepalivetimeout   int         `json:"mepkeepalivetimeout,omitempty"`
	Rtttolerance          int         `json:"rtttolerance,omitempty"`
	Svcstatelearningtime  int         `json:"svcstatelearningtime,omitempty"`
	V6ldnsmasklen         int         `json:"v6ldnsmasklen,omitempty"`
}
