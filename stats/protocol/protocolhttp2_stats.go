package protocol

type Protocolhttp2stats struct {
  Clearstats string `json:"clearstats,omitempty"`
  Http2continuationframesrcvd int `json:"http2continuationframesrcvd,omitempty"`
  Http2continuationframesrcvdrate int `json:"http2continuationframesrcvdrate,omitempty"`
  Http2continuationframessent int `json:"http2continuationframessent,omitempty"`
  Http2continuationframessentrate int `json:"http2continuationframessentrate,omitempty"`
  Http2dataframesrcvd int `json:"http2dataframesrcvd,omitempty"`
  Http2dataframesrcvdrate int `json:"http2dataframesrcvdrate,omitempty"`
  Http2dataframessent int `json:"http2dataframessent,omitempty"`
  Http2dataframessentrate int `json:"http2dataframessentrate,omitempty"`
  Http2direct int `json:"http2direct,omitempty"`
  Http2directrate int `json:"http2directrate,omitempty"`
  Http2errempfraflood int `json:"http2errempfraflood,omitempty"`
  Http2errempfrafloodrate int `json:"http2errempfrafloodrate,omitempty"`
  Http2errresfraflood int `json:"http2errresfraflood,omitempty"`
  Http2errresfrafloodrate int `json:"http2errresfrafloodrate,omitempty"`
  Http2errsetflood int `json:"http2errsetflood,omitempty"`
  Http2errsetfloodrate int `json:"http2errsetfloodrate,omitempty"`
  Http2frametoobig int `json:"http2frametoobig,omitempty"`
  Http2frametoobigrate int `json:"http2frametoobigrate,omitempty"`
  Http2goawayframesrcvd int `json:"http2goawayframesrcvd,omitempty"`
  Http2goawayframesrcvdrate int `json:"http2goawayframesrcvdrate,omitempty"`
  Http2goawayframessent int `json:"http2goawayframessent,omitempty"`
  Http2goawayframessentrate int `json:"http2goawayframessentrate,omitempty"`
  Http2grpcfailurerate int `json:"http2grpcfailurerate,omitempty"`
  Http2grpcrequestrate int `json:"http2grpcrequestrate,omitempty"`
  Http2grpcresponserate int `json:"http2grpcresponserate,omitempty"`
  Http2grpcsuccessrate int `json:"http2grpcsuccessrate,omitempty"`
  Http2headerframesrcvd int `json:"http2headerframesrcvd,omitempty"`
  Http2headerframesrcvdrate int `json:"http2headerframesrcvdrate,omitempty"`
  Http2headerframessent int `json:"http2headerframessent,omitempty"`
  Http2headerframessentrate int `json:"http2headerframessentrate,omitempty"`
  Http2incontinuationframes int `json:"http2incontinuationframes,omitempty"`
  Http2incontinuationframesrate int `json:"http2incontinuationframesrate,omitempty"`
  Http2indataframes int `json:"http2indataframes,omitempty"`
  Http2indataframesrate int `json:"http2indataframesrate,omitempty"`
  Http2ingoawayframes int `json:"http2ingoawayframes,omitempty"`
  Http2ingoawayframesrate int `json:"http2ingoawayframesrate,omitempty"`
  Http2inheaderframes int `json:"http2inheaderframes,omitempty"`
  Http2inheaderframesrate int `json:"http2inheaderframesrate,omitempty"`
  Http2inpingframes int `json:"http2inpingframes,omitempty"`
  Http2inpingframesrate int `json:"http2inpingframesrate,omitempty"`
  Http2inpriorityframes int `json:"http2inpriorityframes,omitempty"`
  Http2inpriorityframesrate int `json:"http2inpriorityframesrate,omitempty"`
  Http2inpushpromiseframes int `json:"http2inpushpromiseframes,omitempty"`
  Http2inpushpromiseframesrate int `json:"http2inpushpromiseframesrate,omitempty"`
  Http2inrststreamframes int `json:"http2inrststreamframes,omitempty"`
  Http2inrststreamframesrate int `json:"http2inrststreamframesrate,omitempty"`
  Http2insettingframes int `json:"http2insettingframes,omitempty"`
  Http2insettingframesrate int `json:"http2insettingframesrate,omitempty"`
  Http2inwindowupdateframes int `json:"http2inwindowupdateframes,omitempty"`
  Http2inwindowupdateframesrate int `json:"http2inwindowupdateframesrate,omitempty"`
  Http2nomatcipher int `json:"http2nomatcipher,omitempty"`
  Http2nomatcipherrate int `json:"http2nomatcipherrate,omitempty"`
  Http2pingflood int `json:"http2pingflood,omitempty"`
  Http2pingfloodrate int `json:"http2pingfloodrate,omitempty"`
  Http2pingframesrcvd int `json:"http2pingframesrcvd,omitempty"`
  Http2pingframesrcvdrate int `json:"http2pingframesrcvdrate,omitempty"`
  Http2pingframessent int `json:"http2pingframessent,omitempty"`
  Http2pingframessentrate int `json:"http2pingframessentrate,omitempty"`
  Http2priorityframesrcvd int `json:"http2priorityframesrcvd,omitempty"`
  Http2priorityframesrcvdrate int `json:"http2priorityframesrcvdrate,omitempty"`
  Http2priorityframessent int `json:"http2priorityframessent,omitempty"`
  Http2priorityframessentrate int `json:"http2priorityframessentrate,omitempty"`
  Http2pushpromframesrcvd int `json:"http2pushpromframesrcvd,omitempty"`
  Http2pushpromframesrcvdrate int `json:"http2pushpromframesrcvdrate,omitempty"`
  Http2pushpromiseframessent int `json:"http2pushpromiseframessent,omitempty"`
  Http2pushpromiseframessentrate int `json:"http2pushpromiseframessentrate,omitempty"`
  Http2requests int `json:"http2requests,omitempty"`
  Http2requestsrate int `json:"http2requestsrate,omitempty"`
  Http2requestupgradefailed int `json:"http2requestupgradefailed,omitempty"`
  Http2requestupgradefailedrate int `json:"http2requestupgradefailedrate,omitempty"`
  Http2requpg int `json:"http2requpg,omitempty"`
  Http2requpgrate int `json:"http2requpgrate,omitempty"`
  Http2responses int `json:"http2responses,omitempty"`
  Http2responsesrate int `json:"http2responsesrate,omitempty"`
  Http2rststreamframesrcvd int `json:"http2rststreamframesrcvd,omitempty"`
  Http2rststreamframesrcvdrate int `json:"http2rststreamframesrcvdrate,omitempty"`
  Http2rststreamframessent int `json:"http2rststreamframessent,omitempty"`
  Http2rststreamframessentrate int `json:"http2rststreamframessentrate,omitempty"`
  Http2serverdirect int `json:"http2serverdirect,omitempty"`
  Http2serverdirectfailed int `json:"http2serverdirectfailed,omitempty"`
  Http2serverdirectfailedrate int `json:"http2serverdirectfailedrate,omitempty"`
  Http2serverdirectrate int `json:"http2serverdirectrate,omitempty"`
  Http2serverupgradefailed int `json:"http2serverupgradefailed,omitempty"`
  Http2serverupgradefailedrate int `json:"http2serverupgradefailedrate,omitempty"`
  Http2settingframesrcvd int `json:"http2settingframesrcvd,omitempty"`
  Http2settingframesrcvdrate int `json:"http2settingframesrcvdrate,omitempty"`
  Http2settingframessent int `json:"http2settingframessent,omitempty"`
  Http2settingframessentrate int `json:"http2settingframessentrate,omitempty"`
  Http2totgrpcfailure int `json:"http2totgrpcfailure,omitempty"`
  Http2totgrpcrequest int `json:"http2totgrpcrequest,omitempty"`
  Http2totgrpcresponse int `json:"http2totgrpcresponse,omitempty"`
  Http2totgrpcsuccess int `json:"http2totgrpcsuccess,omitempty"`
  Http2windowupdateframessent int `json:"http2windowupdateframessent,omitempty"`
  Http2windowupdateframessentrate int `json:"http2windowupdateframessentrate,omitempty"`
  Http2winupdateframesrcvd int `json:"http2winupdateframesrcvd,omitempty"`
  Http2winupdateframesrcvdrate int `json:"http2winupdateframesrcvdrate,omitempty"`
}
