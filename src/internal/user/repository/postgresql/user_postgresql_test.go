package postgresql

import (
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestInsertSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	user := &domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "img.ru/aboba/avatar",
		Password:       "aboba_pass",
		CountFollowing: 0,
	}

	query := `INSERT INTO users \(username, email, avatar, password_hash\) VALUES \(\:username, \:email, \:avatar, \:password_hash\) RETURNING id`

	mock.ExpectExec(query).WithArgs(user.Username, user.Email, user.Avatar, user.Password).WillReturnResult(sqlmock.NewResult(1, 1))

	a := NewUserPostgresRepo(sqlxDb)

	err = a.Insert(user)
	assert.NoError(t, err)
}

func TestUpdateSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	userToUpdate1 := &domain.User{
		ID:       3,
		Username: "abobus",
		Email:    "",
		Password: "",
		Avatar:   "new avatar url",
	}
	userToUpdate2 := &domain.User{
		ID:       4,
		Username: "",
		Email:    "aboba@aboba.ru",
		Password: "aboba",
		Avatar:   "",
	}

	query1 := `UPDATE users SET username \= \$1, avatar \= \$2 WHERE id \= \$3`
	mock.ExpectExec(query1).WithArgs(userToUpdate1.Username, userToUpdate1.Avatar, userToUpdate1.ID).WillReturnResult(sqlmock.NewResult(int64(userToUpdate1.ID), 1))
	query2 := `UPDATE users SET email \= \$1, password_hash \= \$2 WHERE id \= \$3`
	mock.ExpectExec(query2).WithArgs(userToUpdate2.Email, userToUpdate2.Password, userToUpdate2.ID).WillReturnResult(sqlmock.NewResult(int64(userToUpdate2.ID), 1))

	a := NewUserPostgresRepo(sqlxDb)

	err = a.Update(userToUpdate1.ID, userToUpdate1)
	assert.NoError(t, err)
	err = a.Update(userToUpdate2.ID, userToUpdate2)
	assert.NoError(t, err)
}

func TestUpdateError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	userToUpdate := &domain.User{
		ID:       4,
		Username: "",
		Email:    "aboba@aboba.ru",
		Password: "aboba",
		Avatar:   "",
	}
	query := `UPDATE users SET email \= \$1, password_hash \= \$2 WHERE id \= \$3`
	mock.ExpectExec(query).WithArgs(userToUpdate.Email, userToUpdate.Password, 5).WillReturnError(errors.New("user with such id does not exist"))

	a := NewUserPostgresRepo(sqlxDb)
	err = a.Update(5, userToUpdate)
	assert.ErrorIs(t, err, ErrorUpdateUser)
}
