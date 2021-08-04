package core

import "time"

var DataDir = ""

type Config struct {
	SPACES_KEY      string `json:"SPACES_KEY"`
	SPACES_SECRET   string `json:"SPACES_SECRET"`
	SPACES_ENDPOINT string `json:"SPACES_ENDPOINT"`
	SPACES_BUCKET   string `json:"SPACES_BUCKET"`
}

type Image struct {
	Id      string    `json:"id"`
	WorkId  string    `json:"workId"`
	Path    string    `json:"path"`
	Time    int       `json:"time"`
	Created time.Time `json:"created"`
}

type Work struct {
	Id        string   `json:"id"`
	Title     string   `json:"title"`
	Tags      []string `json:"tags"`
	ImageList []Image  `json:"imageList"`
}

type Training struct {
	Id        string   `json:"id"`
	Title     string   `json:"title"`
	Tags      []string `json:"tags"`
	ImageList []Image  `json:"imageList"`
}
