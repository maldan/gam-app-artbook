package core

import "time"

type Image struct {
	Id        string    `json:"id"`
	WorkId    string    `json:"workId"`
	Url       string    `json:"url"`
	Thumbnail string    `json:"thumbnail"`
	Time      int       `json:"time"`
	Created   time.Time `json:"created"`
}

type Work struct {
	Id        string   `json:"id"`
	Title     string   `json:"title"`
	Tags      []string `json:"tags"`
	ImageList []Image  `json:"imageList"`
}

type Training struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Tags      []string  `json:"tags"`
	Time      int       `json:"time"`
	Url       string    `json:"url"`
	Thumbnail string    `json:"thumbnail"`
	Created   time.Time `json:"created"`
}

type Reference struct {
	Id        string   `json:"id"`
	Category  string   `json:"category"`
	Tags      []string `json:"tags"`
	Url       string   `json:"url"`
	Thumbnail string   `json:"thumbnail"`
}

type Config struct {
	SPACES_KEY      string `json:"SPACES_KEY"`
	SPACES_SECRET   string `json:"SPACES_SECRET"`
	SPACES_ENDPOINT string `json:"SPACES_ENDPOINT"`
	SPACES_BUCKET   string `json:"SPACES_BUCKET"`
}

var DataDir = ""
var AppConfig = Config{}
