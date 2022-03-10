package models

type Artist struct {
	Id             uint64 `json:"id" example:"43"`
	Name           string `json:"name" example:"Imagine Dragons"`
	Photo          string `json:"photo" example:"assets/artist_1.png"`
	CountFollowers uint64 `json:"countFollowers" example:"1001"`
	CountListening uint64 `json:"countListening" example:"7654"`
}
