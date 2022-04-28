package postgresql

import (
	"fmt"
	user_microservice_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
)

type UserPostrgesRepo struct {
	DB *sqlx.DB
}

func NewUserPostgresRepo(db *sqlx.DB) user_microservice_domain.UserRepo {
	return &UserPostrgesRepo{
		DB: db,
	}
}

func (a *UserPostrgesRepo) Insert(user *user_microservice_domain.User) error {
	_, err := a.DB.Exec(insertUserCommand, user.Username, user.Email, user.Avatar, user.Password)
	if err != nil {
		fmt.Println(err)
		return ErrorInsertUser
	}

	return nil
}

func (a *UserPostrgesRepo) Update(id uint, user *user_microservice_domain.User) error {
	updateQuery := `UPDATE Users SET `
	var updateParams []interface{}

	i := 1
	if user.Username != "" {
		updateQuery += fmt.Sprintf(`username = $%d, `, i)
		updateParams = append(updateParams, user.Username)
		i++
	}
	if user.Email != "" {
		updateQuery += fmt.Sprintf(`email = $%d, `, i)
		updateParams = append(updateParams, user.Email)
		i++
	}
	if user.Password != "" {
		updateQuery += fmt.Sprintf(`password_hash = $%d, `, i)
		updateParams = append(updateParams, user.Password)
		i++
	}
	if user.Avatar != "" {
		updateQuery += fmt.Sprintf(`avatar = $%d, `, i)
		updateParams = append(updateParams, user.Avatar)
		i++
	}

	updateQuery = updateQuery[:len(updateQuery)-2]

	updateQuery += fmt.Sprintf(` WHERE id = $%d`, i)
	updateParams = append(updateParams, id)

	_, err := a.DB.Exec(updateQuery, updateParams...)
	if err != nil {
		return ErrorUpdateUser
	}

	return nil
}

func (a *UserPostrgesRepo) Delete(id uint) error {
	_, err := a.DB.Exec(deleteUserCommand, id)
	if err != nil {
		return ErrorDeleteUser
	}

	return nil
}

func (a *UserPostrgesRepo) SelectByID(id uint) (*user_microservice_domain.User, error) {
	var user user_microservice_domain.User

	err := a.DB.Get(&user, selectUserByIdCommand, id)
	if err != nil {
		return nil, ErrorSelectUser
	}

	return &user, nil
}

func (a *UserPostrgesRepo) SelectByUsername(username string) (*user_microservice_domain.User, error) {
	var user user_microservice_domain.User

	err := a.DB.Get(&user, selectUserByUsernameCommand, username)
	if err != nil {
		return nil, ErrorSelectUser
	}

	return &user, nil
}

func (a *UserPostrgesRepo) SelectByEmail(email string) (*user_microservice_domain.User, error) {
	var user user_microservice_domain.User

	err := a.DB.Get(&user, selectUserByEmailCommand, email)
	if err != nil {
		return nil, ErrorSelectUser
	}

	return &user, nil
}

func (a *UserPostrgesRepo) GetSize() (int, error) {
	query := `SELECT count(*) From users;`
	size := 0
	if err := a.DB.Get(&size, query); err != nil {
		return -1, err
	}
	return size, nil
}
