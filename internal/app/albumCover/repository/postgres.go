package AlbumCoverPostgres

import (
	"errors"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/album/albumProto"
	"github.com/jmoiron/sqlx"
	"os"
)

type AlbumCoverRepo struct {
	Sqlx *sqlx.DB
}

func NewAlbumCoverPostgresRepo(db *sqlx.DB) domain.AlbumCoverRepo {
	return &AlbumCoverRepo{
		Sqlx: db,
	}
}

// ----------------------------------------------------------------------
func (table AlbumCoverRepo) Create(cover *albumProto.AlbumCover) error {
	query := `
		INSERT INTO albumcover (quote, is_dark)
		VALUES ($1, $2)
		RETURNING id`

	// do query
	_, err := table.Sqlx.Exec(query, cover.Quote, cover.IsDark)

	return err
}

func (table AlbumCoverRepo) Update(cover *albumProto.AlbumCover) error {
	query := `
		UPDATE albumcover
		SET quote = $1, is_dark = $2
		WHERE id = $3`

	_, err := table.Sqlx.Exec(query, cover.Quote, cover.IsDark, cover.Id)
	return err
	//if err != nil {
	//	return err
	//}

	//rowsAffected, err := res.RowsAffected()
	//if err != nil {
	//	return err
	//}
	//
	//if rowsAffected == 0 {
	//	return errors.New(constants.ErrorNothingToUpdate)
	//}

	//return nil
}

func (table AlbumCoverRepo) Delete(id int64) error {
	query := `DELETE FROM albumcover WHERE id = $1`

	res, err := table.Sqlx.Exec(query, id)

	if err != nil {
		return err
	}

	deleted, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if deleted == 0 {
		return errors.New(constants.IndexOutOfRange)
	}
	return nil
}

func (table AlbumCoverRepo) SelectByID(id int64) (*albumProto.AlbumCover, error) {
	query := `SELECT * FROM albumcover WHERE id = $1;`
	holder := albumProto.AlbumCover{}
	if err := table.Sqlx.Get(&holder, query, id); err != nil {
		return nil, err
	}
	return &holder, nil
}

func (table AlbumCoverRepo) GetAll() ([]*albumProto.AlbumCover, error) {
	query := `SELECT * FROM albumcover ORDER BY id;`

	var covers []*albumProto.AlbumCover
	err := table.Sqlx.Select(&covers, query)
	if err != nil {
		return nil, err
	}

	return covers, nil
}

func (table AlbumCoverRepo) GetLastId() (int64, error) {
	query := `SELECT max(id) from albumcover;`

	lastId := int64(0)
	err := table.Sqlx.Get(&lastId, query)

	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (table AlbumCoverRepo) GetSize() (int64, error) {
	query := `SELECT count(*) From albumcover;`
	size := int64(0)
	if err := table.Sqlx.Get(&size, query); err != nil {
		os.Exit(1)
	}
	return size, nil
}
