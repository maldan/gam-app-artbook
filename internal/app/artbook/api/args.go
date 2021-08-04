package api

import "time"

type ArgsEmpty struct {
}

type ArgsId struct {
	Id string
}

type ArgsImage struct {
	Id      string    `json:"id"`
	WorkId  string    `json:"workId"`
	Files   [][]byte  `json:"files"`
	Time    int       `json:"time"`
	Created time.Time `json:"created"`
}
