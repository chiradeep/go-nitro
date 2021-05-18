package system

type Systemcpustats struct {
	Clearstats string `json:"clearstats,omitempty"`
	Id         int    `json:"id,omitempty"`
	Percpuuse  int    `json:"percpuuse,omitempty"`
}
