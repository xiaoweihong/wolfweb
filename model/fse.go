package model

type Repo struct {
	Id       string `json:"id"`
	Type     string `json:"type"`
	Level    string `json:"level"`
	Capacity int64  `json:"capacity"`
	Size     int64  `json:"size"`
}