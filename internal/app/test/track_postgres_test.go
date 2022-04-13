package test

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	TrackPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/app/track/repository"
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

	track := domain.Track{
		AlbumId:        500,
		ArtistId:       5000,
		Title:          "some track",
		Duration:       300,
		CountLikes:     1000,
		CountListening: 0,
	}

	query := `INSERT INTO track \(album_id, artist_id, title, duration, count_likes, count_listening\) VALUES \(\$1, \$2, \$3, \$4, \$5\, \$6\) RETURNING id`

	mock.ExpectExec(query).WithArgs(track.AlbumId, track.ArtistId, track.Title, track.Duration, track.CountLikes, track.CountListening).WillReturnResult(sqlmock.NewResult(1, 1))
	a := TrackPostgres.NewTrackPostgresRepo(sqlxDb)
	err = a.Insert(track)

	assert.NoError(t, err)
}

func TestUpdateTrackSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	tr := domain.Track{
		AlbumId:        500,
		ArtistId:       5000,
		Title:          "some track",
		Duration:       300,
		CountLikes:     1000,
		CountListening: 0,
	}

	query1 := `UPDATE track SET album_id \= \$1, artist_id \= \$2, title \= \$3, duration \= \$4, count_likes \= \$5, count_listening \= \$6 WHERE id \= \$7`
	mock.ExpectExec(query1).WithArgs(tr.AlbumId, tr.ArtistId, tr.Title, tr.Duration, tr.CountLikes, tr.CountListening, tr.Id).WillReturnResult(sqlmock.NewResult(int64(tr.Id), 1))

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
