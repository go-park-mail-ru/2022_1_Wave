package user_microservice_domain

type User struct {
	ID             uint   `json:"id,omitempty" db:"id"`
	Username       string `json:"username,omitempty" db:"username"`
	Email          string `json:"email,omitempty" db:"email"`
	Avatar         string `json:"avatar,omitempty" db:"avatar"`
	Password       string `json:"password,omitempty" db:"password_hash"`
	CountFollowing int    `json:"count_following,omitempty" db:"count_following"`
}

type UserRepo interface {
	Insert(user *User) error
	Update(id uint, user *User) error
	Delete(id uint) error
	SelectByID(id uint) (*User, error)
	SelectByUsername(username string) (*User, error)
	SelectByEmail(email string) (*User, error)
	GetSize() (int, error)
}
