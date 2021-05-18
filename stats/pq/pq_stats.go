package pq

type Pqstats struct {
	Clearstats              string `json:"clearstats,omitempty"`
	Pqpolicymatchesrate     int    `json:"pqpolicymatchesrate,omitempty"`
	Pqpriority1requests     int    `json:"pqpriority1requests,omitempty"`
	Pqpriority1requestsrate int    `json:"pqpriority1requestsrate,omitempty"`
	Pqpriority2requests     int    `json:"pqpriority2requests,omitempty"`
	Pqpriority2requestsrate int    `json:"pqpriority2requestsrate,omitempty"`
	Pqpriority3requests     int    `json:"pqpriority3requests,omitempty"`
	Pqpriority3requestsrate int    `json:"pqpriority3requestsrate,omitempty"`
	Pqthresholdfailedrate   int    `json:"pqthresholdfailedrate,omitempty"`
	Pqtotalpolicymatches    int    `json:"pqtotalpolicymatches,omitempty"`
	Pqtotalthresholdfailed  int    `json:"pqtotalthresholdfailed,omitempty"`
}
