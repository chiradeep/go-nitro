package pq

type Pqpolicystats struct {
	Clearstats                       string `json:"clearstats,omitempty"`
	Clienttransactionsrate           int    `json:"clienttransactionsrate,omitempty"`
	Policyname                       string `json:"policyname,omitempty"`
	Pqavgclienttransactiontimems     int    `json:"pqavgclienttransactiontimems,omitempty"`
	Pqavgclienttransactiontimemsrate int    `json:"pqavgclienttransactiontimemsrate,omitempty"`
	Pqavgqueuewaittimerate           int    `json:"pqavgqueuewaittimerate,omitempty"`
	Pqclientconnectionsrate          int    `json:"pqclientconnectionsrate,omitempty"`
	Pqcurrentclientconnections       int    `json:"pqcurrentclientconnections,omitempty"`
	Pqcurrentclientconnectionsrate   int    `json:"pqcurrentclientconnectionsrate,omitempty"`
	Pqdropped                        int    `json:"pqdropped,omitempty"`
	Pqdroppedrate                    int    `json:"pqdroppedrate,omitempty"`
	Pqqdepth                         int    `json:"pqqdepth,omitempty"`
	Pqqdepthrate                     int    `json:"pqqdepthrate,omitempty"`
	Pqqueuedepthrate                 int    `json:"pqqueuedepthrate,omitempty"`
	Pqtotavgqueuewaittime            int    `json:"pqtotavgqueuewaittime,omitempty"`
	Pqtotclientconnections           int    `json:"pqtotclientconnections,omitempty"`
	Pqtotqueuedepth                  int    `json:"pqtotqueuedepth,omitempty"`
	Pqvserverip                      string `json:"pqvserverip,omitempty"`
	Pqvserverport                    int    `json:"pqvserverport,omitempty"`
	Totclienttransactions            int    `json:"totclienttransactions,omitempty"`
}
