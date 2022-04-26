package test

import (
	AlbumPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/album/repository/postgres"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestInsertAlbumSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	album := &albumProto.Album{
		Id:              5,
		Title:           "some album",
		ArtistId:        10,
		CountLikes:      100,
		CountListenings: 1000,
		Date:            0,
	}

	query := `INSERT INTO album \(title, artist_id, count_likes, count_listening, date\) VALUES \(\$1, \$2, \$3, \$4, \$5\) RETURNING id`

	mock.ExpectExec(query).WithArgs(album.Title, album.ArtistId, album.CountLikes, album.CountListenings, album.Date).WillReturnResult(sqlmock.NewResult(1, 1))
	a := AlbumPostgres.NewAlbumPostgresRepo(sqlxDb)
	err = a.Create(album)

	assert.NoError(t, err)
}

func TestUpdateAlbumSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	al1 := &albumProto.Album{
		Id:              0,
		Title:           "",
		ArtistId:        0,
		CountLikes:      0,
		CountListenings: 0,
		Date:            0,
	}

	query1 := `UPDATE album SET title \= \$1, artist_id \= \$2, count_likes \= \$3, count_listening \= \$4, date \= \$5 WHERE id \= \$6`
	mock.ExpectExec(query1).WithArgs(al1.Title, al1.ArtistId, al1.CountLikes, al1.CountListenings, al1.Date, al1.Id).WillReturnResult(sqlmock.NewResult(int64(al1.Id), 1))

	a := AlbumPostgres.NewAlbumPostgresRepo(sqlxDb)

	err = a.Update(al1)
	assert.NoError(t, err)
}

func TestDeleteAlbumSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	query := `DELETE FROM album WHERE id \= \$1`
	mock.ExpectExec(query).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	a := AlbumPostgres.NewAlbumPostgresRepo(sqlxDb)
	err = a.Delete(1)
	assert.NoError(t, err)
}

func TestSelectAlbumByIdSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")
	rows := sqlmock.NewRows([]string{"id", "title", "artist_id", "count_likes", "count_listening", "date"}).
		AddRow(10, "aboba", 5, 100, 1000, 0)

	query := `SELECT \* FROM album WHERE id \= \$1`
	mock.ExpectQuery(query).WithArgs(10).WillReturnRows(rows)

	a := AlbumPostgres.NewAlbumPostgresRepo(sqlxDb)
	user, err := a.SelectByID(10)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestSelectAllAlbumsSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")
	rows := sqlmock.NewRows([]string{"id", "title", "artist_id", "count_likes", "count_listening", "date"}).
		AddRow(10, "aboba", 5, 100, 1000, 0).
		AddRow(11, "kek", 500, 1000, 2000, 0).
		AddRow(12, "kekus", 5000, 123, 321, 0)

	query := `SELECT \* FROM album ORDER BY id`
	mock.ExpectQuery(query).WillReturnRows(rows)

	a := AlbumPostgres.NewAlbumPostgresRepo(sqlxDb)
	user, err := a.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestSelectPopularAlbumsSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")
	rows := sqlmock.NewRows([]string{"id", "title", "artist_id", "count_likes", "count_listening", "date"}).
		AddRow(16, "aboba", 5, 100, 10000, 0).
		AddRow(13, "kek", 500, 1000, 2000, 0).
		AddRow(567, "kekus", 5000, 123, 321, 0)

	query := `SELECT \* FROM album ORDER BY count_listening DESC LIMIT \$1`
	mock.ExpectQuery(query).WillReturnRows(rows)

	a := AlbumPostgres.NewAlbumPostgresRepo(sqlxDb)
	user, err := a.GetPopular()

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestGetLastIdAlbumsSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	rows := sqlmock.NewRows([]string{"id"}).AddRow(100)
	query := `SELECT max\(id\) from album`
	mock.ExpectQuery(query).WillReturnRows(rows)

	a := AlbumPostgres.NewAlbumPostgresRepo(sqlxDb)
	id, err := a.GetLastId()
	assert.NoError(t, err)
	assert.Equal(t, int64(100), id)
}

func TestGetSizeAlbumsSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	rows := sqlmock.NewRows([]string{"count"}).AddRow(100)

	query := `SELECT count\(\*\) From album`
	mock.ExpectQuery(query).WillReturnRows(rows)

	a := AlbumPostgres.NewAlbumPostgresRepo(sqlxDb)
	size, err := a.GetSize()
	assert.NoError(t, err)
	assert.Equal(t, int64(100), size)
}

func TestGetAlbumsFromArtistSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	expected := []*albumProto.Album{
		{
			Id:              16,
			Title:           "aboba",
			ArtistId:        5,
			CountLikes:      100,
			CountListenings: 10000,
			Date:            0,
		},
		{
			Id:              13,
			Title:           "kek",
			ArtistId:        5,
			CountLikes:      1000,
			CountListenings: 2000,
			Date:            0,
		},
		{
			Id:              567,
			Title:           "kekus",
			ArtistId:        5,
			CountLikes:      123,
			CountListenings: 321,
			Date:            0,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "artist_id", "count_likes", "count_listening", "date"}).
		AddRow(16, "aboba", 5, 100, 10000, 0).
		AddRow(13, "kek", 5, 1000, 2000, 0).
		AddRow(567, "kekus", 5, 123, 321, 0)

	query := `SELECT \* FROM album WHERE artist_id \= \$1`
	mock.ExpectQuery(query).WithArgs(5).WillReturnRows(rows)

	a := AlbumPostgres.NewAlbumPostgresRepo(sqlxDb)
	albums, err := a.GetAlbumsFromArtist(5)
	assert.NoError(t, err)
	assert.Equal(t, expected, albums)
}
