package postgresql

import "errors"

const (
	insertUserCommand           = `INSERT INTO Users (username, email, avatar, password_hash) VALUES ($1, $2, $3, $4) RETURNING id`
	deleteUserCommand           = `DELETE FROM Users WHERE id = $1`
	selectUserByIdCommand       = `SELECT id, username, email, avatar, count_following FROM Users WHERE id = $1`
	selectUserByUsernameCommand = `SELECT id, username, email, avatar, count_following FROM Users WHERE username = $1`
	selectUserByEmailCommand    = `SELECT id, username, email, avatar, count_following FROM Users WHERE email = $1`
)

var (
	ErrorGetPasswordHash = errors.New("error getting the password hash")
	ErrorInsertUser      = errors.New("error inserting user")
	ErrorUpdateUser      = errors.New("error updating user")
	ErrorDeleteUser      = errors.New("error deleting user")
	ErrorSelectUser      = errors.New("error selecting user")
)
