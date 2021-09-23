package core

type Translate struct {
	Noun      []string `json:"noun"`
	Verb      []string `json:"verb"`
	Adjective []string `json:"adjective"`
}

type Word struct {
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Translate Translate `json:"translate"`
}

var DataDir = ""
