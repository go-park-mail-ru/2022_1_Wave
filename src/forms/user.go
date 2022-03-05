package forms

type User struct {
	ID       string `json:"id" example:"5125112"`
	Name     string `json:"name" example:"Martin"`
	Email    string `json:"email" example:"hello@example.com"`
	Password string `json:"password" example:"1fsgh2rfafas"`
}
