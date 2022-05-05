package ArtistPostgres

import (
	"errors"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/artistProto"
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
func (table ArtistRepo) Create(dom *artistProto.Artist) error {
	query := `
		INSERT INTO artist (name, count_followers, count_listening, count_likes)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	// do query
	_, err := table.Sqlx.Exec(query, dom.Name, dom.CountFollowers, dom.CountListenings, dom.CountLikes)

	return err
}

func (table ArtistRepo) Update(dom *artistProto.Artist) error {
	query := `
		UPDATE artist SET name = $1, count_followers = $2, count_listening = $3, count_likes = $4
		WHERE id = $5`

	res, err := table.Sqlx.Exec(query, dom.Name, dom.CountFollowers, dom.CountListenings, dom.CountLikes, dom.Id)
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

func (table ArtistRepo) Delete(id int64) error {
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

func (table ArtistRepo) SelectByID(id int64) (*artistProto.Artist, error) {
	query := `SELECT * FROM artist WHERE id = $1 ORDER BY id;`
	holder := artistProto.Artist{}
	if err := table.Sqlx.Get(&holder, query, id); err != nil {
		return nil, err
	}
	return &holder, nil
}

func (table ArtistRepo) GetAll() ([]*artistProto.Artist, error) {
	query := `SELECT * FROM artist ORDER BY id;`

	var artists []*artistProto.Artist
	err := table.Sqlx.Select(&artists, query)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func (table ArtistRepo) GetPopular() ([]*artistProto.Artist, error) {
	query := `
			SELECT *
			FROM artist
			ORDER BY count_listening DESC
			LIMIT $1;`

	var artists []*artistProto.Artist
	err := table.Sqlx.Select(&artists, query, constants.Top)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func (table ArtistRepo) GetLastId() (int64, error) {
	query := `SELECT max(id) from artist;`

	lastId := int64(0)
	err := table.Sqlx.Get(&lastId, query)

	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (table ArtistRepo) GetSize() (int64, error) {
	query := `SELECT count(*) From artist;`
	size := int64(0)
	if err := table.Sqlx.Get(&size, query); err != nil {
		os.Exit(1)
	}
	return size, nil
}

func (table ArtistRepo) Like(artistId int64, userId int64) error {
	artist, err := table.SelectByID(artistId)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO userArtistsLike (user_id, artist_id)
		VALUES ($1, $2)
		RETURNING artist_id`

	if _, err := table.Sqlx.Exec(query, userId, artist.Id); err != nil {
		return err
	}

	artist.CountLikes = artist.CountLikes + 1
	return table.Update(artist)
}

func (table ArtistRepo) Listen(trackId int64) error {
	artist, err := table.SelectByID(trackId)
	if err != nil {
		return err
	}
	artist.CountListenings = artist.CountListenings + 1
	if err := table.Update(artist); err != nil {
		return err
	}
	return nil
}

func (table ArtistRepo) SearchByName(title string) ([]*artistProto.Artist, error) {
	query := `
			SELECT *
			FROM artist
			WHERE to_tsvector("name") @@ plainto_tsquery($1)
			ORDER BY ts_rank(to_tsvector("name"), plainto_tsquery($1)) DESC
			LIMIT $2;
			`

	var artists []*artistProto.Artist
	err := table.Sqlx.Select(&artists, query, title, constants.SearchTop)
	if err != nil {
		return nil, err
	}

	if len(artists) == 0 {
		arg := title + "%"
		query := `
			SELECT *
			FROM artist
			WHERE lower(name) LIKE lower($1)
			LIMIT $2;
			`
		err := table.Sqlx.Select(&artists, query, arg, constants.SearchTop)
		if err != nil {
			return nil, err
		}
	}

	return artists, nil
}

func (table ArtistRepo) AddToFavorites(artistId int64, userId int64) error {
	artist, err := table.SelectByID(artistId)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO userFavoriteArtists (user_id, artist_id)
		VALUES ($1, $2)
		RETURNING artist_id`

	// do query
	_, err = table.Sqlx.Exec(query, userId, artist.Id)

	return err
}

func (table ArtistRepo) GetFavorites(userId int64) ([]*artistProto.Artist, error) {
	query := `SELECT id, name, count_likes, count_listening, count_followers FROM Artist
			  JOIN userFavoriteArtists favorite ON favorite.artist_id = artist.id
    	      WHERE favorite.user_id = $1 ORDER BY artist_id;`

	// do query
	var artists []*artistProto.Artist
	err := table.Sqlx.Select(&artists, query, userId)

	return artists, err
}

func (table ArtistRepo) RemoveFromFavorites(trackId int64, userId int64) error {
	query := `DELETE FROM userFavoriteArtists WHERE user_id = $1 and artist_id = $2`

	_, err := table.Sqlx.Exec(query, userId, trackId)
	return err
}
