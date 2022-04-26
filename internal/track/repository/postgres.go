package TrackPostgres

import (
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	"github.com/jmoiron/sqlx"
	"os"
)

type TrackRepo struct {
	Sqlx *sqlx.DB
}

func NewTrackPostgresRepo(db *sqlx.DB) domain.TrackRepo {
	return &TrackRepo{
		Sqlx: db,
	}
}

func (table TrackRepo) Create(dom *trackProto.Track) error {
	query := `
		INSERT INTO track (album_id, artist_id, title, duration, count_likes, count_listening)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	// do query
	_, err := table.Sqlx.Exec(query, dom.AlbumId, dom.ArtistId, dom.Title, dom.Duration, dom.CountLikes, dom.CountListenings)

	return err
}

func (table TrackRepo) Update(dom *trackProto.Track) error {
	query := `
		UPDATE track
		SET album_id = $1, artist_id = $2, title = $3, duration = $4, count_likes = $5, count_listening = $6
		WHERE id = $7`

	_, err := table.Sqlx.Exec(query, dom.AlbumId, dom.ArtistId, dom.Title, dom.Duration, dom.CountLikes, dom.CountListenings, dom.Id)
	return err
}

func (table TrackRepo) Delete(id int64) error {
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

func (table TrackRepo) SelectByID(id int64) (*trackProto.Track, error) {
	query := `SELECT * FROM track WHERE id = $1;`
	holder := trackProto.Track{}
	if err := table.Sqlx.Get(&holder, query, id); err != nil {
		return nil, err
	}
	return &holder, nil
}

func (table TrackRepo) GetAll() ([]*trackProto.Track, error) {
	query := `SELECT * FROM track ORDER BY id;`

	var tracks []*trackProto.Track
	err := table.Sqlx.Select(&tracks, query)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

func (table TrackRepo) GetPopular() ([]*trackProto.Track, error) {
	query := `
			SELECT *
			FROM track
			ORDER BY count_listening DESC
			LIMIT $1;`

	var tracks []*trackProto.Track
	err := table.Sqlx.Select(&tracks, query, constants.Top)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

func (table TrackRepo) GetLastId() (int64, error) {
	query := `SELECT max(id) from track;`

	lastId := int64(0)
	err := table.Sqlx.Get(&lastId, query)

	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (table TrackRepo) GetSize() (int64, error) {
	query := `SELECT count(*) From track;`
	size := int64(0)
	if err := table.Sqlx.Get(&size, query); err != nil {
		os.Exit(1)
	}
	return size, nil
}

func (table TrackRepo) GetTracksFromAlbum(albumId int64) ([]*trackProto.Track, error) {
	var tracks []*trackProto.Track
	if err := table.Sqlx.Select(&tracks, `SELECT * FROM track WHERE album_id = $1`, albumId); err != nil {
		return nil, err
	}
	return tracks, nil
}

func (table TrackRepo) GetPopularTracksFromArtist(artistId int64) ([]*trackProto.Track, error) {
	var tracks []*trackProto.Track
	if err := table.Sqlx.Select(&tracks, `
			SELECT * FROM track
			WHERE artist_id = $1
			ORDER BY count_listening DESC
			LIMIT $2;`, artistId, constants.Top); err != nil {
		return nil, err
	}

	return tracks, nil
}

func (table TrackRepo) Like(trackId int64) error {
	track, err := table.SelectByID(trackId)
	if err != nil {
		return err
	}
	track.CountLikes = track.CountLikes + 1
	if err := table.Update(track); err != nil {
		return err
	}
	return nil
}

func (table TrackRepo) Listen(trackId int64) error {
	track, err := table.SelectByID(trackId)
	if err != nil {
		return err
	}
	track.CountListenings = track.CountListenings + 1
	if err := table.Update(track); err != nil {
		return err
	}
	return nil
}

func (table TrackRepo) SearchByTitle(title string) ([]*trackProto.Track, error) {
	query := `
			SELECT *
			FROM track
			WHERE to_tsvector("title") @@ plainto_tsquery($1)
			ORDER BY ts_rank(to_tsvector("title"), plainto_tsquery($1)) DESC;
			`

	var tracks []*trackProto.Track
	err := table.Sqlx.Select(&tracks, query, title)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}
