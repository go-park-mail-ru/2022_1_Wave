package domain

import "errors"

var (
	ErrSetUserPlayerState = errors.New("error while set user player state")
	ErrGetUserPlayerState = errors.New("error while get user player state")
	ErrUnmarshal          = errors.New("error while unmarshal user player state")
	ErrDeletePlayerState  = errors.New("error while delete user player state")
)
