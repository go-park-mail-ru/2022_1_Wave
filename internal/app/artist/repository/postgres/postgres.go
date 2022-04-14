package ArtistPostgres

import (
	"errors"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/jmoiron/sqlx"
	"os"
)

type ArtistRepo struct {
	Sqlx *sqlx.DB
}

func NewArtistPostgresRepo(db *sqlx.DB) domain.ArtistRepo {
	return &ArtistRepo{
		Sqlx: db,
	}
}

// ----------------------------------------------------------------------
func (table ArtistRepo) Insert(dom domain.Artist) error {
	query := `
		INSERT INTO artist (name, count_followers, count_listening)
		VALUES ($1, $2, $3)
		RETURNING id`

	// do query
	_, err := table.Sqlx.Exec(query, dom.Name, dom.CountFollowers, dom.CountListening)

	return err
}

func (table ArtistRepo) Update(dom domain.Artist) error {
	query := `
		UPDATE artist SET name = $1, count_followers = $2, count_listening = $3
		WHERE id = $4`

	res, err := table.Sqlx.Exec(query, dom.Name, dom.CountFollowers, dom.CountListening, dom.Id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New(constants.ErrorNothingToUpdate)
	}

	return nil
}

func (table ArtistRepo) Delete(id int) error {
	query := `DELETE FROM artist WHERE id = $1`

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

func (table ArtistRepo) SelectByID(id int) (*domain.Artist, error) {
	query := `SELECT * FROM artist WHERE id = $1;`
	holder := domain.Artist{}
	if err := table.Sqlx.Get(&holder, query, id); err != nil {
		return nil, err
	}
	return &holder, nil
}

func (table ArtistRepo) GetAll() ([]domain.Artist, error) {
	query := `SELECT * FROM artist ORDER BY id;`

	var artists []domain.Artist
	err := table.Sqlx.Select(&artists, query)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func (table ArtistRepo) GetPopular() ([]domain.Artist, error) {
	query := `
			SELECT *
			FROM artist
			ORDER BY count_listening DESC
			LIMIT $1;`

	var artists []domain.Artist
	err := table.Sqlx.Select(&artists, query, constants.Top)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func (table ArtistRepo) GetLastId() (int, error) {
	query := `SELECT max(id) from artist;`

	lastId := 0
	err := table.Sqlx.Get(&lastId, query)

	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (table ArtistRepo) GetSize() (int, error) {
	query := `SELECT count(*) From artist;`
	size := 0
	if err := table.Sqlx.Get(&size, query); err != nil {
		os.Exit(1)
	}
	return size, nil
}
