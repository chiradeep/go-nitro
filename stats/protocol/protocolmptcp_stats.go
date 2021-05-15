package protocol

type Protocolmptcpstats struct {
  Clearstats string `json:"clearstats,omitempty"`
  Mptcpaddrremovedrate int `json:"mptcpaddrremovedrate,omitempty"`
  Mptcpconnestrate int `json:"mptcpconnestrate,omitempty"`
  Mptcpcurmpcapablesessions int `json:"mptcpcurmpcapablesessions,omitempty"`
  Mptcpcurpendingjoin int `json:"mptcpcurpendingjoin,omitempty"`
  Mptcpcursesswithoutsfs int `json:"mptcpcursesswithoutsfs,omitempty"`
  Mptcpcursfconnections int `json:"mptcpcursfconnections,omitempty"`
  Mptcpdatalessthandatalenrate int `json:"mptcpdatalessthandatalenrate,omitempty"`
  Mptcpdssarate int `json:"mptcpdssarate,omitempty"`
  Mptcpdssfreshackrate int `json:"mptcpdssfreshackrate,omitempty"`
  Mptcpdssmrate int `json:"mptcpdssmrate,omitempty"`
  Mptcpdssrate int `json:"mptcpdssrate,omitempty"`
  Mptcperraddrid0 int `json:"mptcperraddrid0,omitempty"`
  Mptcperraddrid0rate int `json:"mptcperraddrid0rate,omitempty"`
  Mptcperraddridexist int `json:"mptcperraddridexist,omitempty"`
  Mptcperraddridexistrate int `json:"mptcperraddridexistrate,omitempty"`
  Mptcperrbadcksum int `json:"mptcperrbadcksum,omitempty"`
  Mptcperrbadcksumrate int `json:"mptcperrbadcksumrate,omitempty"`
  Mptcperrbadmapconndrop int `json:"mptcperrbadmapconndrop,omitempty"`
  Mptcperrbadmapconndroprate int `json:"mptcperrbadmapconndroprate,omitempty"`
  Mptcperrcryptonotsupported int `json:"mptcperrcryptonotsupported,omitempty"`
  Mptcperrcryptonotsupportedrate int `json:"mptcperrcryptonotsupportedrate,omitempty"`
  Mptcperrdatafinpassive int `json:"mptcperrdatafinpassive,omitempty"`
  Mptcperrdatafinpassiverate int `json:"mptcperrdatafinpassiverate,omitempty"`
  Mptcperrdupmaprecvd int `json:"mptcperrdupmaprecvd,omitempty"`
  Mptcperrdupmaprecvdrate int `json:"mptcperrdupmaprecvdrate,omitempty"`
  Mptcperrextnflagset int `json:"mptcperrextnflagset,omitempty"`
  Mptcperrextnflagsetrate int `json:"mptcperrextnflagsetrate,omitempty"`
  Mptcperrfastclose int `json:"mptcperrfastclose,omitempty"`
  Mptcperrfastclosekey int `json:"mptcperrfastclosekey,omitempty"`
  Mptcperrfastclosekeyrate int `json:"mptcperrfastclosekeyrate,omitempty"`
  Mptcperrfastclosepassive int `json:"mptcperrfastclosepassive,omitempty"`
  Mptcperrfastclosepassiverate int `json:"mptcperrfastclosepassiverate,omitempty"`
  Mptcperrfastcloserate int `json:"mptcperrfastcloserate,omitempty"`
  Mptcperrinvalcookie int `json:"mptcperrinvalcookie,omitempty"`
  Mptcperrinvalcookierate int `json:"mptcperrinvalcookierate,omitempty"`
  Mptcperrinvalidsfn int `json:"mptcperrinvalidsfn,omitempty"`
  Mptcperrinvalidsfnrate int `json:"mptcperrinvalidsfnrate,omitempty"`
  Mptcperrinvalmac int `json:"mptcperrinvalmac,omitempty"`
  Mptcperrinvalmacrate int `json:"mptcperrinvalmacrate,omitempty"`
  Mptcperrinvalopts int `json:"mptcperrinvalopts,omitempty"`
  Mptcperrinvaloptsrate int `json:"mptcperrinvaloptsrate,omitempty"`
  Mptcperrinvalremaddr int `json:"mptcperrinvalremaddr,omitempty"`
  Mptcperrinvalremaddrrate int `json:"mptcperrinvalremaddrrate,omitempty"`
  Mptcperrjoinafterfallback int `json:"mptcperrjoinafterfallback,omitempty"`
  Mptcperrjoinafterfallbackrate int `json:"mptcperrjoinafterfallbackrate,omitempty"`
  Mptcperrjointhreshold int `json:"mptcperrjointhreshold,omitempty"`
  Mptcperrjointhresholdrate int `json:"mptcperrjointhresholdrate,omitempty"`
  Mptcperrmapexists int `json:"mptcperrmapexists,omitempty"`
  Mptcperrmapexistsrate int `json:"mptcperrmapexistsrate,omitempty"`
  Mptcperrmaxsf int `json:"mptcperrmaxsf,omitempty"`
  Mptcperrmaxsfrate int `json:"mptcperrmaxsfrate,omitempty"`
  Mptcperrmpcapsessionallocfail int `json:"mptcperrmpcapsessionallocfail,omitempty"`
  Mptcperrmpcapsessionallocfailrate int `json:"mptcperrmpcapsessionallocfailrate,omitempty"`
  Mptcperrnomappktrcvd int `json:"mptcperrnomappktrcvd,omitempty"`
  Mptcperrnomappktrcvdrate int `json:"mptcperrnomappktrcvdrate,omitempty"`
  Mptcperrnopayloadlenpkt int `json:"mptcperrnopayloadlenpkt,omitempty"`
  Mptcperrnopayloadlenpktrate int `json:"mptcperrnopayloadlenpktrate,omitempty"`
  Mptcperrnosffreensb int `json:"mptcperrnosffreensb,omitempty"`
  Mptcperrnosffreensbrate int `json:"mptcperrnosffreensbrate,omitempty"`
  Mptcperrnsballocfailed int `json:"mptcperrnsballocfailed,omitempty"`
  Mptcperrnsballocfailedrate int `json:"mptcperrnsballocfailedrate,omitempty"`
  Mptcperroptiondiscarded int `json:"mptcperroptiondiscarded,omitempty"`
  Mptcperroptiondiscardedrate int `json:"mptcperroptiondiscardedrate,omitempty"`
  Mptcperroptsnosession int `json:"mptcperroptsnosession,omitempty"`
  Mptcperroptsnosessionrate int `json:"mptcperroptsnosessionrate,omitempty"`
  Mptcperroptssendrst int `json:"mptcperroptssendrst,omitempty"`
  Mptcperroptssendrstrate int `json:"mptcperroptssendrstrate,omitempty"`
  Mptcperrremaddrself int `json:"mptcperrremaddrself,omitempty"`
  Mptcperrremaddrselfrate int `json:"mptcperrremaddrselfrate,omitempty"`
  Mptcperrresflagset int `json:"mptcperrresflagset,omitempty"`
  Mptcperrresflagsetrate int `json:"mptcperrresflagsetrate,omitempty"`
  Mptcperrretxpktrcvd int `json:"mptcperrretxpktrcvd,omitempty"`
  Mptcperrretxpktrcvdrate int `json:"mptcperrretxpktrcvdrate,omitempty"`
  Mptcperrrssffail int `json:"mptcperrrssffail,omitempty"`
  Mptcperrrssffailrate int `json:"mptcperrrssffailrate,omitempty"`
  Mptcperrsfsessionallocfail int `json:"mptcperrsfsessionallocfail,omitempty"`
  Mptcperrsfsessionallocfailrate int `json:"mptcperrsfsessionallocfailrate,omitempty"`
  Mptcperrunknowntoken int `json:"mptcperrunknowntoken,omitempty"`
  Mptcperrunknowntokenrate int `json:"mptcperrunknowntokenrate,omitempty"`
  Mptcperrunsupportedmssnegotiated int `json:"mptcperrunsupportedmssnegotiated,omitempty"`
  Mptcperrunsupportedmssnegotiatedrate int `json:"mptcperrunsupportedmssnegotiatedrate,omitempty"`
  Mptcperrversionnotsupported int `json:"mptcperrversionnotsupported,omitempty"`
  Mptcperrversionnotsupportedrate int `json:"mptcperrversionnotsupportedrate,omitempty"`
  Mptcpestsfreplacedrate int `json:"mptcpestsfreplacedrate,omitempty"`
  Mptcpfreshackfrwdrate int `json:"mptcpfreshackfrwdrate,omitempty"`
  Mptcpinfinitemapfrwdrate int `json:"mptcpinfinitemapfrwdrate,omitempty"`
  Mptcpinfinitemaprecvd int `json:"mptcpinfinitemaprecvd,omitempty"`
  Mptcpinfinitemaprecvdrate int `json:"mptcpinfinitemaprecvdrate,omitempty"`
  Mptcpmoredatarcvdrate int `json:"mptcpmoredatarcvdrate,omitempty"`
  Mptcpmpcapfackrecvdrate int `json:"mptcpmpcapfackrecvdrate,omitempty"`
  Mptcpmpcapsessionrate int `json:"mptcpmpcapsessionrate,omitempty"`
  Mptcpmpcapsfpcballocrate int `json:"mptcpmpcapsfpcballocrate,omitempty"`
  Mptcpmpcapsteeredrate int `json:"mptcpmpcapsteeredrate,omitempty"`
  Mptcpmpcapsynacksentrate int `json:"mptcpmpcapsynacksentrate,omitempty"`
  Mptcpmpcapsynrate int `json:"mptcpmpcapsynrate,omitempty"`
  Mptcpmpcballocfailedrate int `json:"mptcpmpcballocfailedrate,omitempty"`
  Mptcpmpfailrecvd int `json:"mptcpmpfailrecvd,omitempty"`
  Mptcpmpfailrecvdrate int `json:"mptcpmpfailrecvdrate,omitempty"`
  Mptcpmpfailsent int `json:"mptcpmpfailsent,omitempty"`
  Mptcpmpfailsentrate int `json:"mptcpmpfailsentrate,omitempty"`
  Mptcpmpjoin4thacksentrate int `json:"mptcpmpjoin4thacksentrate,omitempty"`
  Mptcpmpjoinfackrecvdrate int `json:"mptcpmpjoinfackrecvdrate,omitempty"`
  Mptcpmpjoinsteeredrate int `json:"mptcpmpjoinsteeredrate,omitempty"`
  Mptcpmpjoinsynacksentrate int `json:"mptcpmpjoinsynacksentrate,omitempty"`
  Mptcpmpjoinsynrate int `json:"mptcpmpjoinsynrate,omitempty"`
  Mptcppendsfreplacedrate int `json:"mptcppendsfreplacedrate,omitempty"`
  Mptcpplainackfallback int `json:"mptcpplainackfallback,omitempty"`
  Mptcpplainackfallbackrate int `json:"mptcpplainackfallbackrate,omitempty"`
  Mptcpplainackrst int `json:"mptcpplainackrst,omitempty"`
  Mptcpplainackrstrate int `json:"mptcpplainackrstrate,omitempty"`
  Mptcppriobackuprx int `json:"mptcppriobackuprx,omitempty"`
  Mptcppriobackuprxrate int `json:"mptcppriobackuprxrate,omitempty"`
  Mptcpprioclearbackuprx int `json:"mptcpprioclearbackuprx,omitempty"`
  Mptcpprioclearbackuprxrate int `json:"mptcpprioclearbackuprxrate,omitempty"`
  Mptcprxdatafinrate int `json:"mptcprxdatafinrate,omitempty"`
  Mptcprxdssrate int `json:"mptcprxdssrate,omitempty"`
  Mptcpsfconnrate int `json:"mptcpsfconnrate,omitempty"`
  Mptcptotaddrremoved int `json:"mptcptotaddrremoved,omitempty"`
  Mptcptotconnest int `json:"mptcptotconnest,omitempty"`
  Mptcptotdatalessthandatalen int `json:"mptcptotdatalessthandatalen,omitempty"`
  Mptcptotdss int `json:"mptcptotdss,omitempty"`
  Mptcptotdssa int `json:"mptcptotdssa,omitempty"`
  Mptcptotdssfreshack int `json:"mptcptotdssfreshack,omitempty"`
  Mptcptotdssm int `json:"mptcptotdssm,omitempty"`
  Mptcptotestsfreplaced int `json:"mptcptotestsfreplaced,omitempty"`
  Mptcptotfreshackfrwd int `json:"mptcptotfreshackfrwd,omitempty"`
  Mptcptotinfinitemapfrwd int `json:"mptcptotinfinitemapfrwd,omitempty"`
  Mptcptotmoredatarcvd int `json:"mptcptotmoredatarcvd,omitempty"`
  Mptcptotmpcapfackrecvd int `json:"mptcptotmpcapfackrecvd,omitempty"`
  Mptcptotmpcapsession int `json:"mptcptotmpcapsession,omitempty"`
  Mptcptotmpcapsfpcballoc int `json:"mptcptotmpcapsfpcballoc,omitempty"`
  Mptcptotmpcapsteered int `json:"mptcptotmpcapsteered,omitempty"`
  Mptcptotmpcapsyn int `json:"mptcptotmpcapsyn,omitempty"`
  Mptcptotmpcapsynacksent int `json:"mptcptotmpcapsynacksent,omitempty"`
  Mptcptotmpcballocfailed int `json:"mptcptotmpcballocfailed,omitempty"`
  Mptcptotmpjoin4thacksent int `json:"mptcptotmpjoin4thacksent,omitempty"`
  Mptcptotmpjoinfackrecvd int `json:"mptcptotmpjoinfackrecvd,omitempty"`
  Mptcptotmpjoinsteered int `json:"mptcptotmpjoinsteered,omitempty"`
  Mptcptotmpjoinsyn int `json:"mptcptotmpjoinsyn,omitempty"`
  Mptcptotmpjoinsynacksent int `json:"mptcptotmpjoinsynacksent,omitempty"`
  Mptcptotpendsfreplaced int `json:"mptcptotpendsfreplaced,omitempty"`
  Mptcptotrxdatafin int `json:"mptcptotrxdatafin,omitempty"`
  Mptcptotrxdss int `json:"mptcptotrxdss,omitempty"`
  Mptcptotsfconn int `json:"mptcptotsfconn,omitempty"`
  Mptcptottxdatafin int `json:"mptcptottxdatafin,omitempty"`
  Mptcptottxdss int `json:"mptcptottxdss,omitempty"`
  Mptcptottxsffin int `json:"mptcptottxsffin,omitempty"`
  Mptcptxdatafinrate int `json:"mptcptxdatafinrate,omitempty"`
  Mptcptxdssrate int `json:"mptcptxdssrate,omitempty"`
  Mptcptxsffinrate int `json:"mptcptxsffinrate,omitempty"`
}
