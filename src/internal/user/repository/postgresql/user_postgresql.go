package postgresql

import (
	"fmt"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

// TODO: вынести эту логику в usecase
func getPasswordHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func checkPassword(password string, passwordHash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)) == nil
}

type UserPostrgesRepo struct {
	DB *sqlx.DB
}

func NewUserPostgresRepo(db *sqlx.DB) domain.UserRepo {
	return &UserPostrgesRepo{
		DB: db,
	}
}

func (a *UserPostrgesRepo) Insert(user *domain.User) error {
	passwordHash, err := getPasswordHash(string(user.Password))
	if err != nil {
		return ErrorGetPasswordHash
	}

	user.Password = string(passwordHash)

	_, err = a.DB.NamedQuery(insertUserCommand, user)
	if err != nil {
		return ErrorInsertUser
	}

	return nil
}

func (a *UserPostrgesRepo) Update(id uint, user *domain.User) error {
	updateQuery := `UPDATE users SET `
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
		passwordHash, err := getPasswordHash(string(user.Password))
		if err != nil {
			return ErrorGetPasswordHash
		}
		user.Password = string(passwordHash)

		updateQuery += fmt.Sprintf(`password_hash = $%d, `, i)
		updateParams = append(updateParams, user.Password)
		i++
	}
	if user.Avatar != "" {
		updateQuery += fmt.Sprintf(`username = $%d`, i)
		updateParams = append(updateParams, user.Avatar)
		i++
	}

	updateQuery = updateQuery[:len(updateQuery)-2]

	updateQuery += fmt.Sprintf(` WHERE id = $%d`, i)

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

func (a *UserPostrgesRepo) SelectByID(id uint) (*domain.User, error) {
	var user domain.User

	err := a.DB.Get(&user, selectUserByIdCommand, id)
	if err != nil {
		return nil, ErrorSelectUser
	}

	return &user, nil
}

func (a *UserPostrgesRepo) SelectByUsername(username string) (*domain.User, error) {
	var user domain.User

	err := a.DB.Get(&user, selectUserByUsernameCommand, username)
	if err != nil {
		return nil, ErrorSelectUser
	}

	return &user, nil
}

func (a *UserPostrgesRepo) SelectByEmail(email string) (*domain.User, error) {
	var user domain.User

	err := a.DB.Get(&user, selectUserByEmailCommand, email)
	if err != nil {
		return nil, ErrorSelectUser
	}

	return &user, nil
}

func (a *UserPostrgesRepo) CheckUsernameAndPassword(username string, password string) bool {
	user, err := a.SelectByUsername(username)
	if err != nil {
		return false
	}

	return checkPassword(password, user.Password)
}

func (a *UserPostrgesRepo) CheckEmailAndPassword(email string, password string) bool {
	user, err := a.SelectByEmail(email)
	if err != nil {
		return false
	}

	return checkPassword(password, user.Password)
}
