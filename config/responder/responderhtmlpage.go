package responder

type Responderhtmlpage struct {
	Cacertfile string `json:"cacertfile,omitempty"`
	Comment    string `json:"comment,omitempty"`
	Name       string `json:"name,omitempty"`
	Overwrite  bool   `json:"overwrite,omitempty"`
	Response   string `json:"response,omitempty"`
	Src        string `json:"src,omitempty"`
}
