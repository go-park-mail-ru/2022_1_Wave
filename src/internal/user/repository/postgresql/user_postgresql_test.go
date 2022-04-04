package postgresql

import (
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

	//query := `INSERT INTO users \(username, email, avatar, password_hash\) VALUES \(\$1, \$2, \$3, \$4\)`
	query := `INSERT INTO users \(username, email, avatar, password_hash\) VALUES \(\:username, \:email, \:avatar, \:password_hash\) RETURNING id`

	//mock.ExpectExec("INSERT INTO users").WillReturnResult(sqlmock.NewResult(1, 1))
	//prep := mock.ExpectPrepare(query)
	//prep.ExpectExec().WithArgs(user.Username, user.Email, user.Avatar, user.Password).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(query).WithArgs(user.Username, user.Email, user.Avatar, user.Password).WillReturnResult(sqlmock.NewResult(1, 1))

	a := NewUserPostgresRepo(sqlxDb)

	err = a.Insert(user)
	assert.NoError(t, err)
}
