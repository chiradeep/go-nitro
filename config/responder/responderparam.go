package responder

type Responderparam struct {
	Timeout     int    `json:"timeout,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
}
