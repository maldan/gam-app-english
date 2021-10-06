package core

type Translate struct {
	Noun      []string `json:"noun"`
	Verb      []string `json:"verb"`
	Adjective []string `json:"adjective"`
}

type Category struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type Word struct {
	Name      string    `json:"name"`
	Category  []string  `json:"category"`
	Translate Translate `json:"translate"`
	Power     int       `json:"power"`
}

var DataDir = ""

func UniqueStringSlice(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if entry == "" {
			continue
		}
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
