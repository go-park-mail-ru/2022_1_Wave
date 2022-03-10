package views

type Track struct {
	Title  string `json:"title" example:"Mercury"`
	Artist string `json:"artist" example:"Hexed"`
	Cover  string `json:"cover" example:"assets/album_1.png"`
}

func FromInterfaceToTrackView(data interface{}) Track {
	temp := data.(map[string]interface{})
	track := Track{
		Title:  temp["title"].(string),
		Artist: temp["artist"].(string),
		Cover:  temp["cover"].(string),
	}
	return track
}

func GetTracksViewsFromInterfaces(data []interface{}) []Track {
	tracks := make([]Track, len(data))
	for idx, it := range data {
		tracks[idx] = FromInterfaceToTrackView(it)
	}
	return tracks
}
