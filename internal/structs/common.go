package structs

type TrackIdWrapper struct {
	TrackId int `json:"trackId" example:"4"`
}

type ArtistIdWrapper struct {
	ArtistId int64 `json:"artistId" example:"4"`
}

type AlbumIdWrapper struct {
	AlbumId int64 `json:"albumId" example:"4"`
}
