package test

import (
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/user/repository/postgresql"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
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

	query := `INSERT INTO Users \(username, email, avatar, password_hash\) VALUES \(\$1, \$2, \$3, \$4\) RETURNING id`

	mock.ExpectExec(query).WithArgs(user.Username, user.Email, user.Avatar, user.Password).WillReturnResult(sqlmock.NewResult(1, 1))

	a := postgresql.NewUserPostgresRepo(sqlxDb)

	err = a.Insert(user)
	assert.NoError(t, err)
}

func TestInsertError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDb := sqlx.NewDb(db, "sqlmock")

	user := &domain.User{
		ID:       4,
		Username: "",
		Email:    "aboba@aboba.ru",
		Password: "aboba",
		Avatar:   "",
	}
	query := `INSERT INTO Users \(username, email, avatar, password_hash\) VALUES \(\$1, \$2, \$3, \$4\) RETURNING id`
	mock.ExpectExec(query).WithArgs(user.Username, user.Email, user.Avatar, user.Password).WillReturnError(errors.New("insert error"))

	a := postgresql.NewUserPostgresRepo(sqlxDb)
	err = a.Insert(user)
	assert.ErrorIs(t, err, postgresql.ErrorInsertUser)
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

	query1 := `UPDATE Users SET username \= \$1, avatar \= \$2 WHERE id \= \$3`
	mock.ExpectExec(query1).WithArgs(userToUpdate1.Username, userToUpdate1.Avatar, userToUpdate1.ID).WillReturnResult(sqlmock.NewResult(int64(userToUpdate1.ID), 1))
	query2 := `UPDATE Users SET email \= \$1, password_hash \= \$2 WHERE id \= \$3`
	mock.ExpectExec(query2).WithArgs(userToUpdate2.Email, userToUpdate2.Password, userToUpdate2.ID).WillReturnResult(sqlmock.NewResult(int64(userToUpdate2.ID), 1))

	a := postgresql.NewUserPostgresRepo(sqlxDb)

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
	query := `UPDATE Users SET email \= \$1, password_hash \= \$2 WHERE id \= \$3`
	mock.ExpectExec(query).WithArgs(userToUpdate.Email, userToUpdate.Password, userToUpdate.ID).WillReturnError(errors.New("user with such id does not exist"))

	a := postgresql.NewUserPostgresRepo(sqlxDb)
	err = a.Update(5, userToUpdate)
	assert.ErrorIs(t, err, postgresql.ErrorUpdateUser)
}

func TestDeleteSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	query := `DELETE FROM Users WHERE id \= \$1`
	mock.ExpectExec(query).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	a := postgresql.NewUserPostgresRepo(sqlxDb)
	err = a.Delete(1)
	assert.NoError(t, err)
}

func TestDeleteError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	query := `DELETE FROM Users WHERE id \= \$1`
	mock.ExpectExec(query).WithArgs(1).WillReturnError(errors.New("user with such id does not exist"))

	a := postgresql.NewUserPostgresRepo(sqlxDb)
	err = a.Delete(1)
	assert.ErrorIs(t, err, postgresql.ErrorDeleteUser)
}

func TestSelectByIdSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")
	rows := sqlmock.NewRows([]string{"id", "username", "email", "avatar", "count_following"}).
		AddRow(1, "aboba", "aboba@aboba.ru", "url_to_avatar", 0)

	query := `SELECT id, username, email, avatar, password_hash, count_following FROM Users WHERE id \= \$1`
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	a := postgresql.NewUserPostgresRepo(sqlxDb)
	user, err := a.SelectByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestSelectByIdError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	query := `SELECT id, username, email, avatar, password_hash, count_following FROM Users WHERE id \= \$1`
	mock.ExpectExec(query).WithArgs(1).WillReturnError(errors.New("user with such id does not exist"))

	a := postgresql.NewUserPostgresRepo(sqlxDb)
	user, err := a.SelectByID(1)

	assert.ErrorIs(t, err, postgresql.ErrorSelectUser)
	assert.Nil(t, user)
}

func TestSelectByUsernameSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")
	rows := sqlmock.NewRows([]string{"id", "username", "email", "avatar", "count_following"}).
		AddRow(1, "aboba", "aboba@aboba.ru", "url_to_avatar", 0)

	query := `SELECT id, username, email, avatar, password_hash, count_following FROM Users WHERE username \= \$1`
	mock.ExpectQuery(query).WithArgs("aboba").WillReturnRows(rows)

	a := postgresql.NewUserPostgresRepo(sqlxDb)
	user, err := a.SelectByUsername("aboba")

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestSelectByUsernameError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	query := `SELECT id, username, email, avatar, password_hash, count_following FROM Users WHERE username \= \$1`
	mock.ExpectExec(query).WithArgs("not_aboba").WillReturnError(errors.New("user with such username does not exist"))

	a := postgresql.NewUserPostgresRepo(sqlxDb)
	user, err := a.SelectByUsername("not_aboba")

	assert.ErrorIs(t, err, postgresql.ErrorSelectUser)
	assert.Nil(t, user)
}

func TestSelectByEmailSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")
	rows := sqlmock.NewRows([]string{"id", "username", "email", "avatar", "count_following"}).
		AddRow(1, "aboba", "aboba@aboba.ru", "url_to_avatar", 0)

	query := `SELECT id, username, email, avatar, password_hash, count_following FROM Users WHERE email \= \$1`
	mock.ExpectQuery(query).WithArgs("aboba@aboba.ru").WillReturnRows(rows)

	a := postgresql.NewUserPostgresRepo(sqlxDb)
	user, err := a.SelectByEmail("aboba@aboba.ru")

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestSelectByEmailError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	query := `SELECT id, username, email, avatar, password_hash, count_following FROM Users WHERE email \= \$1`
	mock.ExpectExec(query).WithArgs("not_aboba@aboba.ru").WillReturnError(errors.New("user with such username does not exist"))

	a := postgresql.NewUserPostgresRepo(sqlxDb)
	user, err := a.SelectByEmail("not_aboba@aboba.ru")

	assert.ErrorIs(t, err, postgresql.ErrorSelectUser)
	assert.Nil(t, user)
}
