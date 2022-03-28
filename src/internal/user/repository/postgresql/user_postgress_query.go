package postgresql

import "errors"

const (
	insertUserCommand           = `INSERT INTO users (email, username, password_hash) VALUES (:email, :username, :password_hash) RETURNING id`
	deleteUserCommand           = `DELETE FROM users WHERE id = $1`
	selectUserByIdCommand       = `SELECT id, email, username, count_following FROM users WHERE id = $1`
	selectUserByUsernameCommand = `SELECT id, email, username, count_following FROM users WHERE username = $1`
	selectUserByEmailCommand    = `SELECT id, email, username, count_following FROM users WHERE email = $1`
)

var (
	ErrorGetPasswordHash = errors.New("error getting the password hash")
	ErrorInsertUser      = errors.New("error inserting user")
	ErrorUpdateUser      = errors.New("error updating user")
	ErrorDeleteUser      = errors.New("error deleting user")
	ErrorSelectUser      = errors.New("error selecting user")
)
