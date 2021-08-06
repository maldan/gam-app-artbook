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
	Time    int       `json:"time"`
	Created time.Time `json:"created"`
	Files   [][]byte  `json:"files"`
}

type ArgsTraining struct {
	Id      string    `json:"id"`
	Title   string    `json:"title"`
	Tags    string    `json:"tags"`
	Time    int       `json:"time"`
	Created time.Time `json:"created"`
	Files   [][]byte  `json:"files"`
}

type ArgsReference struct {
	Id       string   `json:"id"`
	Category string   `json:"category"`
	Tags     string   `json:"tags"`
	Files    [][]byte `json:"files"`
}

type ArgsDate struct {
	Date time.Time
}
