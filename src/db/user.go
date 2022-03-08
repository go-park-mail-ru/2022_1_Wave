package db

import (
	"errors"
	"sync"
)

const (
	NoUserWithUsername           = "user with such username does not exist"
	NoUserWithId                 = "user with such id does not exist"
	NoUserWithEmail              = "user with such email does not exist"
	UserWithUsernameAlreadyExist = "user with such username has already exist"
	UserWithEmailAlreadyExist    = "user with such email has already exist"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type UserStorage struct {
	users []User
	m     sync.RWMutex
}

type UserRep interface {
	Insert(user *User) error
	Update(user *User) error
	Delete(id uint) error
	SelectByID(id uint) (*User, error)
	SelectByUsername(username string) (*User, error)
	SelectByEmail(email string) (*User, error)
	CheckUsernameAndPassword(username string, password string) bool
	CheckEmailAndPassword(email string, password string) bool
}

var MyUserStorage UserRep = &UserStorage{
	users: make([]User, 0),
}

func (storage *UserStorage) SelectByEmail(email string) (*User, error) {
	storage.m.RLock()
	defer storage.m.RUnlock()
	for _, user := range storage.users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, errors.New(NoUserWithEmail)
}

func (storage *UserStorage) CheckEmailAndPassword(email string, password string) bool {
	user, err := storage.SelectByEmail(email)
	if err != nil {
		return false
	}
	return user.Password == password
}

func (storage *UserStorage) CheckUsernameAndPassword(username string, password string) bool {
	user, err := storage.SelectByUsername(username)
	if err != nil {
		return false
	}
	return user.Password == password
}

func (storage *UserStorage) Delete(id uint) error {
	storage.m.Lock()
	defer storage.m.Unlock()
	for i := 0; i < len(storage.users); i++ {
		if storage.users[i].ID == id {
			storage.users = append(storage.users[:i], storage.users[i+1:]...)
			return nil
		}
	}

	return errors.New(NoUserWithId)
}

func (storage *UserStorage) SelectByUsername(username string) (*User, error) {
	storage.m.RLock()
	defer storage.m.RUnlock()
	for _, user := range storage.users {
		if user.Username == username {
			return &user, nil
		}
	}

	return nil, errors.New(NoUserWithUsername)
}

func (storage *UserStorage) Insert(user *User) error {
	storage.m.RLock()
	if _, err := storage.SelectByUsername(user.Username); err == nil {
		return errors.New(UserWithUsernameAlreadyExist)
	}
	if _, err := storage.SelectByEmail(user.Email); err == nil {
		return errors.New(UserWithEmailAlreadyExist)
	}
	storage.m.RUnlock()

	user.ID = uint(len(storage.users)) + 1

	storage.m.Lock()
	storage.users = append(storage.users, *user)
	storage.m.Unlock()
	return nil
}

func (storage *UserStorage) Update(user *User) error {
	userInDb, err := storage.SelectByID(user.ID)
	if err != nil {
		return err
	}

	storage.m.Lock()
	userInDb.Username = user.Username
	userInDb.Password = user.Password
	userInDb.Email = user.Email
	storage.m.Unlock()

	return nil
}

func (storage *UserStorage) SelectByID(id uint) (*User, error) {
	storage.m.RLock()
	defer storage.m.RUnlock()
	for i := 0; i < len(storage.users); i++ {
		if storage.users[i].ID == id {
			return &storage.users[i], nil
		}
	}

	return nil, errors.New(NoUserWithId)
}
