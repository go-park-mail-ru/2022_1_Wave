package domain

type User struct {
	ID       uint   `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserRepo interface {
	Insert(user *User) error
	Update(id uint, user *User) error
	Delete(id uint) error
	SelectByID(id uint) (*User, error)
	SelectByUsername(username string) (*User, error)
	SelectByEmail(email string) (*User, error)
	CheckUsernameAndPassword(username string, password string) bool
	CheckEmailAndPassword(email string, password string) bool
}

type UserUseCase interface {
	GetById(userId uint) (*User, error)
	GetByUsername(username string) (*User, error)
	GetByEmail(email string) (*User, error)
	GetBySessionId(sessionId string) (*User, error)
	Update(id uint, user *User) error
	DeleteById(userId uint) error
	DeleteByUsername(username string) error
	DeleteByEmail(email string) error
	DeleteBySessionId(sessionId string) error
	CheckUsernameAndPassword(username string, password string) bool
	CheckEmailAndPassword(email string, password string) bool
}
