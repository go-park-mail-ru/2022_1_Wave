package views

const (
	AlbumTitleField  = "title"
	AlbumArtistField = "artist"
	AlbumCoverField  = "cover"
)

type Album struct {
	Title  string `json:"title" example:"Mercury"`
	Artist string `json:"artist" example:"Hexed"`
	Cover  string `json:"cover" example:"assets/album_1.png"`
}
