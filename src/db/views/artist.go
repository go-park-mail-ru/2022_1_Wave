package views

type Artist struct {
	Name  string `json:"name" example:"Mercury"`
	Cover string `json:"cover" example:"assets/artist_1.png"`
}

func SetArtistsViewsFromInterfaces(data []interface{}) []Artist {
	artists := make([]Artist, len(data))
	for idx, it := range data {
		temp := it.(map[string]interface{})
		artists[idx] = Artist{
			Name:  temp["name"].(string),
			Cover: temp["cover"].(string),
		}
	}
	return artists
}
