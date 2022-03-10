package models

type Track struct {
	Id             uint64 `json:"id" example:"777"`
	AlbumId        uint64 `json:"albumId" example:"143"`
	AuthorId       uint64 `json:"authorId" example:"121"`
	Title          string `json:"title" example:"Rain"`
	Duration       uint64 `json:"duration" example:"180"`
	Mp4            string `json:"mp4" example:"assets/track_1.mp4"`
	CoverId        uint64 `json:"coverId" example:"254"`
	CountLikes     uint64 `json:"countLikes" example:"54"`
	CountListening uint64 `json:"countListening" example:"15632"`
}
