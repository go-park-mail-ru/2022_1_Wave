package test

import (
	ArtistPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/app/artist/repository/postgres"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
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

	artist := domain.Artist{
		Id:             5,
		Name:           "imagine",
		CountLikes:     100,
		CountFollowers: 11110,
		CountListening: 1000,
	}

	query := `INSERT INTO artist \(name, count_followers, count_listening\) VALUES \(\$1, \$2, \$3\) RETURNING id`

	mock.ExpectExec(query).WithArgs(artist.Name, artist.CountFollowers, artist.CountListening).WillReturnResult(sqlmock.NewResult(1, 1))
	a := ArtistPostgres.NewArtistPostgresRepo(sqlxDb)
	err = a.Insert(artist)

	assert.NoError(t, err)
}

func TestUpdateArtistSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	al1 := domain.Artist{
		Id:             0,
		Name:           "qweqwe",
		CountLikes:     110,
		CountFollowers: 1110,
		CountListening: 1111110,
	}

	query1 := `UPDATE artist SET name \= \$1, count_followers \= \$2, count_listening \= \$3 WHERE id \= \$4`
	mock.ExpectExec(query1).WithArgs(al1.Name, al1.CountFollowers, al1.CountListening, al1.Id).WillReturnResult(sqlmock.NewResult(int64(al1.Id), 1))

	a := ArtistPostgres.NewArtistPostgresRepo(sqlxDb)

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

	a := ArtistPostgres.NewArtistPostgresRepo(sqlxDb)
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

	a := ArtistPostgres.NewArtistPostgresRepo(sqlxDb)
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

	a := ArtistPostgres.NewArtistPostgresRepo(sqlxDb)
	user, err := a.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, user)
}
