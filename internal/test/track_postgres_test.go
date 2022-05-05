package test

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	TrackPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/track/repository"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestInsertTrackSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	track := &trackProto.Track{
		AlbumId:         500,
		ArtistId:        5000,
		Title:           "some track",
		Duration:        300,
		CountLikes:      1000,
		CountListenings: 0,
	}

	query := `INSERT INTO track \(album_id, artist_id, title, duration, count_likes, count_listening\) VALUES \(\$1, \$2, \$3, \$4, \$5\, \$6\) RETURNING id`

	mock.ExpectExec(query).WithArgs(track.AlbumId, track.ArtistId, track.Title, track.Duration, track.CountLikes, track.CountListenings).WillReturnResult(sqlmock.NewResult(1, 1))
	a := TrackPostgres.NewTrackPostgresRepo(sqlxDb)
	err = a.Create(track)

	assert.NoError(t, err)
}

func TestUpdateTrackSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	tr := &trackProto.Track{
		AlbumId:         500,
		ArtistId:        5000,
		Title:           "some track",
		Duration:        300,
		CountLikes:      1000,
		CountListenings: 0,
	}

	query1 := `UPDATE track SET album_id \= \$1, artist_id \= \$2, title \= \$3, duration \= \$4, count_likes \= \$5, count_listening \= \$6 WHERE id \= \$7`
	mock.ExpectExec(query1).WithArgs(tr.AlbumId, tr.ArtistId, tr.Title, tr.Duration, tr.CountLikes, tr.CountListenings, tr.Id).WillReturnResult(sqlmock.NewResult(int64(tr.Id), 1))

	a := TrackPostgres.NewTrackPostgresRepo(sqlxDb)

	err = a.Update(tr)
	assert.NoError(t, err)
}

func TestDeleteTrackSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	query := `DELETE FROM track WHERE id \= \$1`
	mock.ExpectExec(query).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	a := TrackPostgres.NewTrackPostgresRepo(sqlxDb)
	err = a.Delete(1)
	assert.NoError(t, err)
}

func TestSelectTrackByIdSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")
	rows := sqlmock.NewRows([]string{"id", "album_id", "artist_id", "title", "duration", "count_likes", "count_listening"}).
		AddRow(10, 4, 5, "amogus", 1000, 0, 321)

	query := `SELECT \* FROM track WHERE id \= \$1`
	mock.ExpectQuery(query).WithArgs(10).WillReturnRows(rows)

	a := TrackPostgres.NewTrackPostgresRepo(sqlxDb)
	user, err := a.SelectByID(10)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestSelectAllTracksSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")
	rows := sqlmock.NewRows([]string{"id", "album_id", "artist_id", "title", "duration", "count_likes", "count_listening"}).
		AddRow(105, 41, 45, "amogus1", 100, 0, 4321).
		AddRow(106, 43, 55, "amogus2", 10, 0, 1321).
		AddRow(107, 42, 85, "amogus3", 41, 0, 2321)

	query := `SELECT \* FROM track ORDER BY id`
	mock.ExpectQuery(query).WillReturnRows(rows)

	a := TrackPostgres.NewTrackPostgresRepo(sqlxDb)
	user, err := a.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestSelectPopularTracksSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")
	rows := sqlmock.NewRows([]string{"id", "album_id", "artist_id", "title", "duration", "count_likes", "count_listening"}).
		AddRow(105, 41, 45, "amogus1", 100, 0, 4321).
		AddRow(106, 43, 55, "amogus2", 10, 0, 1321).
		AddRow(107, 42, 85, "amogus3", 41, 0, 2321)

	query := `SELECT \* FROM track ORDER BY count_listening DESC LIMIT \$1`
	mock.ExpectQuery(query).WillReturnRows(rows)

	a := TrackPostgres.NewTrackPostgresRepo(sqlxDb)
	user, err := a.GetPopular()

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestGetLastIdTracksSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	rows := sqlmock.NewRows([]string{"id"}).AddRow(100)
	query := `SELECT max\(id\) from track`
	mock.ExpectQuery(query).WillReturnRows(rows)

	a := TrackPostgres.NewTrackPostgresRepo(sqlxDb)
	id, err := a.GetLastId()
	assert.NoError(t, err)
	assert.Equal(t, int64(100), id)
}

func TestGetSizeTracksSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	rows := sqlmock.NewRows([]string{"count"}).AddRow(100)

	query := `SELECT count\(\*\) From track`
	mock.ExpectQuery(query).WillReturnRows(rows)

	a := TrackPostgres.NewTrackPostgresRepo(sqlxDb)
	size, err := a.GetSize()
	assert.NoError(t, err)
	assert.Equal(t, int64(100), size)
}

func TestSelectPopularTracksFromArtistSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	expected := []*trackProto.Track{
		{
			Id:              1,
			AlbumId:         12,
			ArtistId:        5,
			Title:           "some song",
			Duration:        310,
			CountLikes:      10,
			CountListenings: 11110,
		},
		{
			Id:              500,
			AlbumId:         98,
			ArtistId:        5,
			Title:           "and more",
			Duration:        10,
			CountLikes:      40,
			CountListenings: 760,
		},
		{
			Id:              56,
			AlbumId:         111,
			ArtistId:        5,
			Title:           "and finish",
			Duration:        110,
			CountLikes:      76,
			CountListenings: 10,
		},
	}

	sqlxDb := sqlx.NewDb(db, "sqlmock")
	rows := sqlmock.NewRows([]string{"id", "album_id", "artist_id", "title", "duration", "count_likes", "count_listening"}).
		AddRow(1, 12, 5, "some song", 310, 10, 11110).
		AddRow(500, 98, 5, "and more", 10, 40, 760).
		AddRow(56, 111, 5, "and finish", 110, 76, 10)

	query := `SELECT \* FROM track WHERE artist_id \= \$1 ORDER BY count_listening DESC LIMIT \$2`
	mock.ExpectQuery(query).WillReturnRows(rows)

	a := TrackPostgres.NewTrackPostgresRepo(sqlxDb)
	tracks, err := a.GetPopularTracksFromArtist(5)

	assert.NoError(t, err)
	assert.NotNil(t, tracks)
	assert.Equal(t, expected, tracks)
}
