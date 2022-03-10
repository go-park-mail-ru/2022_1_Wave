package forms

type Result struct {
	Status string `json:"status" example:"OK"`
	Result string `json:"result,omitempty" example:"some success note"`
	Error  string `json:"error,omitempty" example:"some error note"`
}
