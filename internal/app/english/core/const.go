package core

type Translate struct {
	Noun      []string `json:"noun"`
	Verb      []string `json:"verb"`
	Adjective []string `json:"adjective"`
}

type Word struct {
	Name      string    `json:"name"`
	Translate Translate `json:"translate"`
}

var DataDir = ""
