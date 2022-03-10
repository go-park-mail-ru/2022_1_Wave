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

func GetAlbumsViewsFromInterfaces(data []interface{}) []Album {
	albums := make([]Album, len(data))
	for idx, it := range data {
		temp := it.(map[string]interface{})
		albums[idx] = Album{
			Title:  temp["title"].(string),
			Artist: temp["artist"].(string),
			Cover:  temp["cover"].(string),
		}
	}
	return albums
}
