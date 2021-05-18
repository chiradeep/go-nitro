package contentinspection

type Contentinspectionstats struct {
	Clearstats                 string `json:"clearstats,omitempty"`
	Icap100contrecv            int    `json:"icap100contrecv,omitempty"`
	Icap204enabledrequests     int    `json:"icap204enabledrequests,omitempty"`
	Icap204nocontentrecv       int    `json:"icap204nocontentrecv,omitempty"`
	Icapadaptiverequests       int    `json:"icapadaptiverequests,omitempty"`
	Icapadaptiveresponses      int    `json:"icapadaptiveresponses,omitempty"`
	Icapcalloutcompleted       int    `json:"icapcalloutcompleted,omitempty"`
	Icapcalloutinitiated       int    `json:"icapcalloutinitiated,omitempty"`
	Icaperrorshandled          int    `json:"icaperrorshandled,omitempty"`
	Icappreviewenabledrequests int    `json:"icappreviewenabledrequests,omitempty"`
	Icapreqmodrequests         int    `json:"icapreqmodrequests,omitempty"`
	Icaprespmodrequests        int    `json:"icaprespmodrequests,omitempty"`
	Icapserverdownbypass       int    `json:"icapserverdownbypass,omitempty"`
	Icapserverdowndrop         int    `json:"icapserverdowndrop,omitempty"`
	Icapserverdownreset        int    `json:"icapserverdownreset,omitempty"`
	Inlinegeneratedresponses   int    `json:"inlinegeneratedresponses,omitempty"`
	Inlinereqbytesrecv         int    `json:"inlinereqbytesrecv,omitempty"`
	Inlinereqbytessent         int    `json:"inlinereqbytessent,omitempty"`
	Inlinerequestssent         int    `json:"inlinerequestssent,omitempty"`
	Inlinerespbytesrecv        int    `json:"inlinerespbytesrecv,omitempty"`
	Inlinerespbytessent        int    `json:"inlinerespbytessent,omitempty"`
	Inlineresponsessent        int    `json:"inlineresponsessent,omitempty"`
	Inlineserverdownbypass     int    `json:"inlineserverdownbypass,omitempty"`
	Inlineserverdowndrop       int    `json:"inlineserverdowndrop,omitempty"`
	Inlineserverdownreset      int    `json:"inlineserverdownreset,omitempty"`
	Mirrorreqbytessent         int    `json:"mirrorreqbytessent,omitempty"`
	Mirrorrequestssent         int    `json:"mirrorrequestssent,omitempty"`
	Mirrorrespbytessent        int    `json:"mirrorrespbytessent,omitempty"`
	Mirrorresponsessent        int    `json:"mirrorresponsessent,omitempty"`
	Mirrorserverdownbypass     int    `json:"mirrorserverdownbypass,omitempty"`
	Mirrorserverdowndrop       int    `json:"mirrorserverdowndrop,omitempty"`
	Mirrorserverdownreset      int    `json:"mirrorserverdownreset,omitempty"`
}
