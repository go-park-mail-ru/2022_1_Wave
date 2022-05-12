package ArtistPostgres

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/artistProto"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestInsertArtistSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	artist := &artistProto.Artist{
		Id:              5,
		Name:            "imagine",
		CountLikes:      100,
		CountFollowers:  11110,
		CountListenings: 1000,
	}

	query := `INSERT INTO artist \(name, count_followers, count_listening, count_likes\) VALUES \(\$1, \$2, \$3, \$4\) RETURNING id`

	mock.ExpectExec(query).WithArgs(artist.Name, artist.CountFollowers, artist.CountListenings, artist.CountLikes).WillReturnResult(sqlmock.NewResult(1, 1))
	a := NewArtistPostgresRepo(sqlxDb)
	err = a.Create(artist)

	assert.NoError(t, err)
}

func TestUpdateArtistSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	al1 := &artistProto.Artist{
		Id:              0,
		Name:            "qweqwe",
		CountLikes:      110,
		CountFollowers:  1110,
		CountListenings: 1111110,
	}

	query1 := `UPDATE artist SET name \= \$1, count_followers \= \$2, count_listening \= \$3, count_likes \= \$4 WHERE id \= \$5`
	mock.ExpectExec(query1).WithArgs(al1.Name, al1.CountFollowers, al1.CountListenings, al1.CountLikes, al1.Id).WillReturnResult(sqlmock.NewResult(int64(al1.Id), 1))

	a := NewArtistPostgresRepo(sqlxDb)

	err = a.Update(al1)
	assert.NoError(t, err)
}

func TestDeleteArtistSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	query := `DELETE FROM artist WHERE id \= \$1`
	mock.ExpectExec(query).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	a := NewArtistPostgresRepo(sqlxDb)
	err = a.Delete(1)
	assert.NoError(t, err)
}

func TestSelectArtistByIdSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")
	rows := sqlmock.NewRows([]string{"name", "count_followers", "count_listening"}).
		AddRow("aboba", 500, 1230)

	query := `SELECT \* FROM artist WHERE id \= \$1`
	mock.ExpectQuery(query).WithArgs(10).WillReturnRows(rows)

	a := NewArtistPostgresRepo(sqlxDb)
	user, err := a.SelectByID(10)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestSelectAllArtistsSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")
	rows := sqlmock.NewRows([]string{"id", "name", "count_followers", "count_listening"}).
		AddRow(1, "aboba1", 12500, 1230).
		AddRow(2, "aboba2", 321500, 1230).
		AddRow(3, "aboba3", 512300, 1230)

	query := `SELECT \* FROM artist ORDER BY id`
	mock.ExpectQuery(query).WillReturnRows(rows)

	a := NewArtistPostgresRepo(sqlxDb)
	user, err := a.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestSelectPopularArtistsSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")
	rows := sqlmock.NewRows([]string{"id", "name", "count_followers", "count_listening"}).
		AddRow(1, "aboba1", 12500, 1230).
		AddRow(2, "aboba2", 321500, 1230).
		AddRow(3, "aboba3", 512300, 1230)

	query := `SELECT \* FROM artist ORDER BY count_listening DESC LIMIT \$1`
	mock.ExpectQuery(query).WillReturnRows(rows)

	a := NewArtistPostgresRepo(sqlxDb)
	user, err := a.GetPopular()

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestGetLastIdArtistsSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	rows := sqlmock.NewRows([]string{"id"}).AddRow(100)
	query := `SELECT max\(id\) from artist`
	mock.ExpectQuery(query).WillReturnRows(rows)

	a := NewArtistPostgresRepo(sqlxDb)
	id, err := a.GetLastId()
	assert.NoError(t, err)
	assert.Equal(t, int64(100), id)
}

func TestGetSizeArtistsSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	rows := sqlmock.NewRows([]string{"count"}).AddRow(100)

	query := `SELECT count\(\*\) From artist`
	mock.ExpectQuery(query).WillReturnRows(rows)

	a := NewArtistPostgresRepo(sqlxDb)
	size, err := a.GetSize()
	assert.NoError(t, err)
	assert.Equal(t, int64(100), size)
}
