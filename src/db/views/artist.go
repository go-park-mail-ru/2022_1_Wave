package views

type Artist struct {
	Name  string `json:"name" example:"Mercury"`
	Cover string `json:"cover" example:"assets/artist_1.png"`
}

func FromInterfaceToArtistView(data interface{}) Artist {
	temp := data.(map[string]interface{})
	artist := Artist{
		Name:  temp["name"].(string),
		Cover: temp["cover"].(string),
	}
	return artist
}

func GetArtistsViewsFromInterfaces(data []interface{}) []Artist {
	artists := make([]Artist, len(data))
	for idx, it := range data {
		artists[idx] = FromInterfaceToArtistView(it)
	}
	return artists
}
