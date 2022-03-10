package views

type Track struct {
	Title  string `json:"title" example:"Mercury"`
	Artist string `json:"artist" example:"Hexed"`
	Cover  string `json:"cover" example:"assets/album_1.png"`
}

func SetTracksViewsFromInterfaces(data []interface{}) []Track {
	tracks := make([]Track, len(data))
	for idx, it := range data {
		temp := it.(map[string]interface{})
		tracks[idx] = Track{
			Title:  temp["title"].(string),
			Artist: temp["artist"].(string),
			Cover:  temp["cover"].(string),
		}
	}
	return tracks
}
