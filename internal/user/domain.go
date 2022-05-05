package user_domain

import (
	"errors"
	user_microservice_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user"
)

type UserAgent interface {
	Create(user *user_microservice_domain.User) error
	Update(id uint, user *user_microservice_domain.User) error
	Delete(id uint) error
	GetById(id uint) (*user_microservice_domain.User, error)
	GetByUsername(username string) (*user_microservice_domain.User, error)
	GetByEmail(email string) (*user_microservice_domain.User, error)
}

type UserUseCase interface {
	GetById(userId uint) (*user_microservice_domain.User, error)
	GetByUsername(username string) (*user_microservice_domain.User, error)
	GetByEmail(email string) (*user_microservice_domain.User, error)
	GetBySessionId(sessionId string) (*user_microservice_domain.User, error)
	Update(id uint, user *user_microservice_domain.User) error
	DeleteById(userId uint) error
	DeleteByUsername(username string) error
	DeleteByEmail(email string) error
	DeleteBySessionId(sessionId string) error
	CheckUsernameAndPassword(username string, password string) bool
	CheckEmailAndPassword(email string, password string) bool
}

var (
	ErrUserAlreadyExist    = errors.New("user already exist")
	ErrUserDoesNotExist    = errors.New("user does not exist")
	ErrSessionDoesNotExist = errors.New("session does not exist")
)
