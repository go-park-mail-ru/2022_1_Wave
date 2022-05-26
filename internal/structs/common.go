package structs

type TrackIdWrapper struct {
	TrackId int `json:"trackId" example:"4" validate:"min=1,nonnil"`
}
