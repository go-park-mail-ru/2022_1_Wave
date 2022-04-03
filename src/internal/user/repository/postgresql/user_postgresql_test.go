package postgresql

import (
	"fmt"
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

	user := &domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "img.ru/aboba/avatar",
		Password:       "aboba_pass",
		CountFollowing: 0,
	}

	query := `INSERT INTO users (username, email, avatar, password_hash) VALUES ($1, $2, $3, $4) RETURNING id`
	//query := `INSERT INTO users \(username, email, avatar, password_hash\) VALUES \(\:username, \:email, \:avatar, \:password_hash\) RETURNING id`

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(user.Username, user.Email, user.Avatar, user.Password).WillReturnResult(sqlmock.NewResult(1, 1))

	db_sqlx := sqlx.NewDb(db, "sql")
	fmt.Println(db_sqlx)
	a := NewUserPostgresRepo(db_sqlx)

	err = a.Insert(user)
	assert.NoError(t, err)
}
