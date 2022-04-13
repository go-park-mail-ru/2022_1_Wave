package TrackPostgres

import (
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	utilsInterfaces "github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	"github.com/jmoiron/sqlx"
	"os"
)

type TrackRepo struct {
	Sqlx *sqlx.DB
}

func NewTrackPostgresRepo(db *sqlx.DB) utilsInterfaces.TrackRepoInterface {
	return &TrackRepo{
		Sqlx: db,
	}
}

func (table TrackRepo) Insert(dom domain.Track) error {
	query := `
		INSERT INTO track (album_id, artist_id, title, duration, count_likes, count_listening)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	// do query
	_, err := table.Sqlx.Exec(query, dom.AlbumId, dom.ArtistId, dom.Title, dom.Duration, dom.CountLikes, dom.CountListening)

	return err
}

func (table TrackRepo) Update(dom domain.Track) error {
	query := `
		UPDATE track
		SET album_id = $1, artist_id = $2, title = $3, duration = $4, count_likes = $5, count_listening = $6
		WHERE id = $7`

	_, err := table.Sqlx.Exec(query, dom.AlbumId, dom.ArtistId, dom.Title, dom.Duration, dom.CountLikes, dom.CountListening, dom.Id)
	return err
}

func (table TrackRepo) Delete(id int) error {
	query := `DELETE FROM track WHERE id = $1`

	_, err := table.Sqlx.Exec(query, id)
	return err
	//
	//if err != nil {
	//	return err
	//}
	//
	//deleted, err := res.RowsAffected()
	//
	//if err != nil {
	//	return err
	//}
	//
	//if deleted == 0 {
	//	return errors.New(constants.IndexOutOfRange)
	//}
	//return nil
}

func (table TrackRepo) SelectByID(id int) (*domain.Track, error) {
	query := `SELECT * FROM track WHERE id = $1;`
	holder := domain.Track{}
	if err := table.Sqlx.Get(&holder, query, id); err != nil {
		return nil, err
	}
	return &holder, nil
}

func (table TrackRepo) GetAll() ([]domain.Track, error) {
	query := `SELECT * FROM track ORDER BY id;`

	var tracks []domain.Track
	err := table.Sqlx.Select(&tracks, query)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

func (table TrackRepo) GetPopular() ([]domain.Track, error) {
	query := `
			SELECT *
			FROM track
			ORDER BY count_listening DESC
			LIMIT $1;`

	var tracks []domain.Track
	err := table.Sqlx.Select(&tracks, query, constants.Top)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

func (table TrackRepo) GetLastId() (int, error) {
	query := `SELECT max(id) from track;`

	lastId := 0
	err := table.Sqlx.Get(&lastId, query)

	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (table TrackRepo) GetSize() (int, error) {
	query := `SELECT count(*) From track;`
	size := 0
	if err := table.Sqlx.Get(&size, query); err != nil {
		os.Exit(1)
	}
	return size, nil
}

func (table TrackRepo) GetTracksFromAlbum(albumId int) ([]domain.Track, error) {
	var tracks []domain.Track
	if err := table.Sqlx.Select(&tracks, `SELECT * FROM track WHERE album_id = $1`, albumId); err != nil {
		return nil, err
	}
	return tracks, nil
}

func (table TrackRepo) GetPopularTracksFromArtist(artistId int) ([]domain.Track, error) {
	var tracks []domain.Track
	if err := table.Sqlx.Select(&tracks, `
			SELECT * FROM track
			WHERE artist_id = $1
			ORDER BY count_listening DESC
			LIMIT $2;`, artistId, constants.Top); err != nil {
		return nil, err
	}

	return tracks, nil
}
