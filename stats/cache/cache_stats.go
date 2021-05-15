package cache

type Cachestats struct {
  Cache304hitsrate int `json:"cache304hitsrate,omitempty"`
  Cache64maxmemorykb int `json:"cache64maxmemorykb,omitempty"`
  Cachebytesserved int `json:"cachebytesserved,omitempty"`
  Cachebytesservedrate int `json:"cachebytesservedrate,omitempty"`
  Cachecompressedbytesserved int `json:"cachecompressedbytesserved,omitempty"`
  Cachecompressedbytesservedrate int `json:"cachecompressedbytesservedrate,omitempty"`
  Cachecurhits int `json:"cachecurhits,omitempty"`
  Cachecurmisses int `json:"cachecurmisses,omitempty"`
  Cacheerrmemalloc int `json:"cacheerrmemalloc,omitempty"`
  Cacheexpireatlastbyterate int `json:"cacheexpireatlastbyterate,omitempty"`
  Cacheflashcachehitsrate int `json:"cacheflashcachehitsrate,omitempty"`
  Cacheflashcachemissesrate int `json:"cacheflashcachemissesrate,omitempty"`
  Cachefulltoconditionalrequestrate int `json:"cachefulltoconditionalrequestrate,omitempty"`
  Cachehitsrate int `json:"cachehitsrate,omitempty"`
  Cacheinvalidationrequestsrate int `json:"cacheinvalidationrequestsrate,omitempty"`
  Cachelargestresponsereceived int `json:"cachelargestresponsereceived,omitempty"`
  Cachemaxmemoryactivekb int `json:"cachemaxmemoryactivekb,omitempty"`
  Cachemaxmemorykb int `json:"cachemaxmemorykb,omitempty"`
  Cachemissesrate int `json:"cachemissesrate,omitempty"`
  Cachenon304hitsrate int `json:"cachenon304hitsrate,omitempty"`
  Cachenonparameterizedinvalidationrequestsrate int `json:"cachenonparameterizedinvalidationrequestsrate,omitempty"`
  Cachenonstoreablemissesrate int `json:"cachenonstoreablemissesrate,omitempty"`
  Cachenumcached int `json:"cachenumcached,omitempty"`
  Cachenummarker int `json:"cachenummarker,omitempty"`
  Cacheparameterized304hitsrate int `json:"cacheparameterized304hitsrate,omitempty"`
  Cacheparameterizedhitsrate int `json:"cacheparameterizedhitsrate,omitempty"`
  Cacheparameterizedinvalidationrequestsrate int `json:"cacheparameterizedinvalidationrequestsrate,omitempty"`
  Cacheparameterizednon304hitsrate int `json:"cacheparameterizednon304hitsrate,omitempty"`
  Cacheparameterizedrequestsrate int `json:"cacheparameterizedrequestsrate,omitempty"`
  Cachepercent304hits int `json:"cachepercent304hits,omitempty"`
  Cachepercentbytehit int `json:"cachepercentbytehit,omitempty"`
  Cachepercenthit int `json:"cachepercenthit,omitempty"`
  Cachepercentoriginbandwidthsaved int `json:"cachepercentoriginbandwidthsaved,omitempty"`
  Cachepercentparameterized304hits int `json:"cachepercentparameterized304hits,omitempty"`
  Cachepercentpethits int `json:"cachepercentpethits,omitempty"`
  Cachepercentstoreablemiss int `json:"cachepercentstoreablemiss,omitempty"`
  Cachepercentsuccessfulrevalidation int `json:"cachepercentsuccessfulrevalidation,omitempty"`
  Cachepethitsrate int `json:"cachepethitsrate,omitempty"`
  Cachepetrequestsrate int `json:"cachepetrequestsrate,omitempty"`
  Cacherecentpercent304hits int `json:"cacherecentpercent304hits,omitempty"`
  Cacherecentpercentbytehit int `json:"cacherecentpercentbytehit,omitempty"`
  Cacherecentpercenthit int `json:"cacherecentpercenthit,omitempty"`
  Cacherecentpercentoriginbandwidthsaved int `json:"cacherecentpercentoriginbandwidthsaved,omitempty"`
  Cacherecentpercentparameterizedhits int `json:"cacherecentpercentparameterizedhits,omitempty"`
  Cacherecentpercentstoreablemiss int `json:"cacherecentpercentstoreablemiss,omitempty"`
  Cacherecentpercentsuccessfulrevalidation int `json:"cacherecentpercentsuccessfulrevalidation,omitempty"`
  Cacherequestsrate int `json:"cacherequestsrate,omitempty"`
  Cacheresponsebytesrate int `json:"cacheresponsebytesrate,omitempty"`
  Cacherevalidationmissrate int `json:"cacherevalidationmissrate,omitempty"`
  Cachesqlhitsrate int `json:"cachesqlhitsrate,omitempty"`
  Cachestoreablemissesrate int `json:"cachestoreablemissesrate,omitempty"`
  Cachesuccessfulrevalidationrate int `json:"cachesuccessfulrevalidationrate,omitempty"`
  Cachetot304hits int `json:"cachetot304hits,omitempty"`
  Cachetotexpireatlastbyte int `json:"cachetotexpireatlastbyte,omitempty"`
  Cachetotflashcachehits int `json:"cachetotflashcachehits,omitempty"`
  Cachetotflashcachemisses int `json:"cachetotflashcachemisses,omitempty"`
  Cachetotfulltoconditionalrequest int `json:"cachetotfulltoconditionalrequest,omitempty"`
  Cachetothits int `json:"cachetothits,omitempty"`
  Cachetotinvalidationrequests int `json:"cachetotinvalidationrequests,omitempty"`
  Cachetotmisses int `json:"cachetotmisses,omitempty"`
  Cachetotnon304hits int `json:"cachetotnon304hits,omitempty"`
  Cachetotnonparameterizedinvalidationrequests int `json:"cachetotnonparameterizedinvalidationrequests,omitempty"`
  Cachetotnonstoreablemisses int `json:"cachetotnonstoreablemisses,omitempty"`
  Cachetotparameterized304hits int `json:"cachetotparameterized304hits,omitempty"`
  Cachetotparameterizedhits int `json:"cachetotparameterizedhits,omitempty"`
  Cachetotparameterizedinvalidationrequests int `json:"cachetotparameterizedinvalidationrequests,omitempty"`
  Cachetotparameterizednon304hits int `json:"cachetotparameterizednon304hits,omitempty"`
  Cachetotparameterizedrequests int `json:"cachetotparameterizedrequests,omitempty"`
  Cachetotpethits int `json:"cachetotpethits,omitempty"`
  Cachetotpetrequests int `json:"cachetotpetrequests,omitempty"`
  Cachetotrequests int `json:"cachetotrequests,omitempty"`
  Cachetotresponsebytes int `json:"cachetotresponsebytes,omitempty"`
  Cachetotrevalidationmiss int `json:"cachetotrevalidationmiss,omitempty"`
  Cachetotsqlhits int `json:"cachetotsqlhits,omitempty"`
  Cachetotstoreablemisses int `json:"cachetotstoreablemisses,omitempty"`
  Cachetotsuccessfulrevalidation int `json:"cachetotsuccessfulrevalidation,omitempty"`
  Cacheutilizedmemorykb int `json:"cacheutilizedmemorykb,omitempty"`
  Clearstats string `json:"clearstats,omitempty"`
}
