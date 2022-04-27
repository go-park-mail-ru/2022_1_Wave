package domain

import "errors"

var (
	ErrInvalidLoginOrPassword = errors.New("invalid login or password")
	ErrUserAlreadyExist       = errors.New("user already exist")
	ErrUserDoesNotExist       = errors.New("user does not exist")

	ErrInsert             = errors.New("insertion error")
	ErrDatabaseUnexpected = errors.New("database unexpected error")

	ErrWhileSetNewSession       = errors.New("error while set new session")
	ErrWhileChangeSession       = errors.New("error while change session")
	ErrSessionStorageUnexpected = errors.New("session storage unexpected error")
	ErrSessionDoesNotExist      = errors.New("session does not exist")

	ErrGetSession    = errors.New("error while get session")
	ErrSetSession    = errors.New("error while set session")
	ErrDeleteSession = errors.New("error while delete session")
)
