package AlbumCoverPostgres

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestInsertAlbumCoverSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	cover := &albumProto.AlbumCover{
		Quote:  "qweqwe",
		IsDark: true,
	}

	query := `INSERT INTO albumcover \(quote, is_dark\) VALUES \(\$1, \$2\) RETURNING id`

	mock.ExpectExec(query).WithArgs(cover.Quote, cover.IsDark).WillReturnResult(sqlmock.NewResult(1, 1))
	a := NewAlbumCoverPostgresRepo(sqlxDb)
	err = a.Create(cover)

	assert.NoError(t, err)
}

func TestUpdateAlbumCoverSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	cvr := &albumProto.AlbumCover{
		Id:     4,
		Quote:  "qweqwe",
		IsDark: true,
	}

	query1 := `UPDATE albumcover SET quote \= \$1, is_dark \= \$2 WHERE id \= \$3`
	mock.ExpectExec(query1).WithArgs(cvr.Quote, cvr.IsDark, cvr.Id).WillReturnResult(sqlmock.NewResult(cvr.Id, 1))

	a := NewAlbumCoverPostgresRepo(sqlxDb)

	err = a.Update(cvr)
	assert.NoError(t, err)
}

func TestDeleteAlbumCoverSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	query := `DELETE FROM albumcover WHERE id \= \$1`
	mock.ExpectExec(query).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	a := NewAlbumCoverPostgresRepo(sqlxDb)
	err = a.Delete(1)
	assert.NoError(t, err)
}

func TestSelectAlbumCoverByIdSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")
	rows := sqlmock.NewRows([]string{"id", "quote", "is_dark"}).
		AddRow(1, "amogus", true)

	query := `SELECT \* FROM albumcover WHERE id \= \$1`
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	a := NewAlbumCoverPostgresRepo(sqlxDb)
	user, err := a.SelectByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestSelectAllAlbumCoversSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")
	rows := sqlmock.NewRows([]string{"id", "quote", "is_dark"}).
		AddRow(1, "amogus1", true).
		AddRow(10, "amogus2", false).
		AddRow(100, "amogus3", true)

	query := `SELECT \* FROM albumcover ORDER BY id`
	mock.ExpectQuery(query).WillReturnRows(rows)

	a := NewAlbumCoverPostgresRepo(sqlxDb)
	user, err := a.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestGetLastIdAlbumCoverSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	rows := sqlmock.NewRows([]string{"id"}).AddRow(100)
	query := `SELECT max\(id\) from albumcover`
	mock.ExpectQuery(query).WillReturnRows(rows)

	a := NewAlbumCoverPostgresRepo(sqlxDb)
	id, err := a.GetLastId()
	assert.NoError(t, err)
	assert.Equal(t, int64(100), id)
}

func TestGetSizeAlbumCoverSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDb := sqlx.NewDb(db, "sqlmock")

	rows := sqlmock.NewRows([]string{"count"}).AddRow(100)

	query := `SELECT count\(\*\) From albumcover`
	mock.ExpectQuery(query).WillReturnRows(rows)

	a := NewAlbumCoverPostgresRepo(sqlxDb)
	size, err := a.GetSize()
	assert.NoError(t, err)
	assert.Equal(t, int64(100), size)
}
