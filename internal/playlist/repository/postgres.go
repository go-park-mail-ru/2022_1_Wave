package TrackPostgres

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/playlist/playlistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	"github.com/jmoiron/sqlx"
	"os"
)

type PlaylistRepo struct {
	Sqlx *sqlx.DB
}

func NewPlaylistPostgresRepo(db *sqlx.DB) domain.PlaylistRepo {
	return &PlaylistRepo{
		Sqlx: db,
	}
}

func (table PlaylistRepo) Create(userId int64, playlist *playlistProto.Playlist) error {
	query := `
		INSERT INTO Playlist (title, tracks_id)
		VALUES ($1, $2)
		RETURNING id`

	// do query

	if _, err := table.Sqlx.Exec(query, playlist.Title, playlist.TracksId); err != nil {
		return err
	}

	query = `
			INSERT INTO userPlaylist(user_id, playlist_id) 
			VALUES ($1, $2)
			RETURNING playlist_id`

	_, err := table.Sqlx.Exec(query, userId, playlist.Id)
	return err
}

func (table PlaylistRepo) Update(userId int64, playlist *playlistProto.Playlist) error {
	query := `SELECT * FROM userPlaylist WHERE user_id = $1`

	query = `
		UPDATE playlist
		SET title = $1, tracks_id = $2
		WHERE id = $3`

	_, err := table.Sqlx.Exec(query, playlist.Title, playlist.TracksId, playlist.Id)
	return err
}

func (table PlaylistRepo) Delete(userId int64, playlistId int64) error {
	query := `DELETE FROM track WHERE id = $1`

	_, err := table.Sqlx.Exec(query, id)
	return err
}

func (table PlaylistRepo) SelectByID(userId int64, playlistId int64) (*playlistProto.Playlist, error) {
	query := `SELECT * FROM track WHERE id = $1 ORDER BY id;`
	holder := trackProto.Track{}
	if err := table.Sqlx.Get(&holder, query, id); err != nil {
		return nil, err
	}
	return &holder, nil
}

func (table PlaylistRepo) GetAll(userId int64) ([]*playlistProto.Playlist, error) {
	query := `SELECT * FROM track ORDER BY id;`

	var tracks []*trackProto.Track
	err := table.Sqlx.Select(&tracks, query)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

func (table PlaylistRepo) GetLastId(userId int64) (int64, error) {
	query := `SELECT max(id) from track;`

	lastId := int64(0)
	err := table.Sqlx.Get(&lastId, query)

	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (table PlaylistRepo) GetSize(userId int64) (int64, error) {
	query := `SELECT count(*) From track;`
	size := int64(0)
	if err := table.Sqlx.Get(&size, query); err != nil {
		os.Exit(1)
	}
	return size, nil
}
