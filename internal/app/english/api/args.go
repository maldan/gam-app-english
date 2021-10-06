package api

import "time"

type ArgsEmpty struct {
}

type ArgsCategory struct {
	Category string
}
type ArgsDate struct {
	Date time.Time
}
