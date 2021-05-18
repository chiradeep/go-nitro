package network

type Inatstats struct {
	Clearstats          string `json:"clearstats,omitempty"`
	Inatnat46drop46     int    `json:"inatnat46drop46,omitempty"`
	Inatnat46drop46rate int    `json:"inatnat46drop46rate,omitempty"`
	Inatnat46drop64     int    `json:"inatnat46drop64,omitempty"`
	Inatnat46drop64rate int    `json:"inatnat46drop64rate,omitempty"`
	Inatnat46icmp46     int    `json:"inatnat46icmp46,omitempty"`
	Inatnat46icmp46rate int    `json:"inatnat46icmp46rate,omitempty"`
	Inatnat46icmp64     int    `json:"inatnat46icmp64,omitempty"`
	Inatnat46icmp64rate int    `json:"inatnat46icmp64rate,omitempty"`
	Inatnat46tcp46      int    `json:"inatnat46tcp46,omitempty"`
	Inatnat46tcp46rate  int    `json:"inatnat46tcp46rate,omitempty"`
	Inatnat46tcp64      int    `json:"inatnat46tcp64,omitempty"`
	Inatnat46tcp64rate  int    `json:"inatnat46tcp64rate,omitempty"`
	Inatnat46udp46      int    `json:"inatnat46udp46,omitempty"`
	Inatnat46udp46rate  int    `json:"inatnat46udp46rate,omitempty"`
	Inatnat46udp64      int    `json:"inatnat46udp64,omitempty"`
	Inatnat46udp64rate  int    `json:"inatnat46udp64rate,omitempty"`
	Name                string `json:"name,omitempty"`
	Nat46drop46rate     int    `json:"nat46drop46rate,omitempty"`
	Nat46drop64rate     int    `json:"nat46drop64rate,omitempty"`
	Nat46icmp46rate     int    `json:"nat46icmp46rate,omitempty"`
	Nat46icmp64rate     int    `json:"nat46icmp64rate,omitempty"`
	Nat46tcp46rate      int    `json:"nat46tcp46rate,omitempty"`
	Nat46tcp64rate      int    `json:"nat46tcp64rate,omitempty"`
	Nat46totdrop46      int    `json:"nat46totdrop46,omitempty"`
	Nat46totdrop64      int    `json:"nat46totdrop64,omitempty"`
	Nat46toticmp46      int    `json:"nat46toticmp46,omitempty"`
	Nat46toticmp64      int    `json:"nat46toticmp64,omitempty"`
	Nat46tottcp46       int    `json:"nat46tottcp46,omitempty"`
	Nat46tottcp64       int    `json:"nat46tottcp64,omitempty"`
	Nat46totudp46       int    `json:"nat46totudp46,omitempty"`
	Nat46totudp64       int    `json:"nat46totudp64,omitempty"`
	Nat46udp46rate      int    `json:"nat46udp46rate,omitempty"`
	Nat46udp64rate      int    `json:"nat46udp64rate,omitempty"`
}
