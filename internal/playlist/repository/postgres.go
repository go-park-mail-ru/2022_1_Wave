package PlaylistPostgres

import (
	"fmt"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/playlist/playlistProto"
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

func (table PlaylistRepo) SelectByIDOfCurrentUser(userId int64, playlistId int64) (*playlistProto.Playlist, error) {
	query := `SELECT id, title FROM playlist
			JOIN userPlaylist ON userPlaylist.user_id = $1 and userPlaylist.playlist_id = $2 
			WHERE userplaylist.playlist_id = playlist.id;`

	holder := playlistProto.Playlist{}
	if err := table.Sqlx.Get(&holder, query, userId, playlistId); err != nil {
		return nil, err
	}

	return &holder, nil
}

func (table PlaylistRepo) SelectById(playlistId int64) (*playlistProto.Playlist, error) {
	query := `SELECT * FROM playlist WHERE id = $1 ORDER BY id;`

	holder := playlistProto.Playlist{}
	if err := table.Sqlx.Get(&holder, query, playlistId); err != nil {
		return nil, err
	}

	return &holder, nil
}

func (table PlaylistRepo) Create(userId int64, playlist *playlistProto.Playlist) error {
	query := `
		INSERT INTO Playlist (title)
		VALUES ($1)
		RETURNING id`

	// do query

	if _, err := table.Sqlx.Exec(query, playlist.Title); err != nil {
		if err != nil {
			return err
		}
	}

	id, err := table.GetLastId()
	if err != nil {
		return err
	}

	playlist.Id = id

	query = `
			INSERT INTO userPlaylist(user_id, playlist_id) 
			VALUES ($1, $2)
			RETURNING playlist_id`

	_, err = table.Sqlx.Exec(query, userId, playlist.Id)
	return err
}

func (table PlaylistRepo) Update(userId int64, playlist *playlistProto.Playlist) error {
	selected, err := table.SelectByIDOfCurrentUser(userId, playlist.Id)
	if err != nil {
		return err
	}

	fmt.Println(selected.Title, selected.Id)

	query := `
		UPDATE playlist
		SET title = $1
		WHERE id = $2`

	_, err = table.Sqlx.Exec(query, playlist.Title, selected.Id)
	return err
}

func (table PlaylistRepo) Delete(userId int64, playlistId int64) error {
	selected, err := table.SelectByIDOfCurrentUser(userId, playlistId)
	if err != nil {
		return err
	}

	query := `
		DELETE FROM playlist
		WHERE id = $1
		`

	_, err = table.Sqlx.Exec(query, selected.Id)
	return err
}

func (table PlaylistRepo) GetAll() ([]*playlistProto.Playlist, error) {
	query := `SELECT * FROM playlist ORDER BY id;`

	var playlists []*playlistProto.Playlist
	err := table.Sqlx.Select(&playlists, query)
	if err != nil {
		return nil, err
	}

	return playlists, nil
}

func (table PlaylistRepo) GetAllOfCurrentUser(userId int64) ([]*playlistProto.Playlist, error) {
	query := `SELECT id, title
			  FROM playlist 
			  JOIN userPlaylist ON userPlaylist.user_id = $1 and playlist.id = userPlaylist.playlist_id
			  ORDER BY id;`

	var playlists []*playlistProto.Playlist
	err := table.Sqlx.Select(&playlists, query, userId)
	if err != nil {
		return nil, err
	}

	return playlists, nil
}

func (table PlaylistRepo) GetLastId() (int64, error) {
	query := `SELECT max(id) from playlist`

	lastId := int64(0)
	err := table.Sqlx.Get(&lastId, query)

	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (table PlaylistRepo) GetLastIdOfCurrentUser(userId int64) (int64, error) {
	query := `SELECT max(playlist_id) from userPlaylist
			  where userPlaylist.user_id = $1;`

	lastId := int64(0)
	err := table.Sqlx.Get(&lastId, query, userId)

	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (table PlaylistRepo) GetSize() (int64, error) {
	query := `SELECT count(*) From playlist;`
	size := int64(0)
	if err := table.Sqlx.Get(&size, query); err != nil {
		os.Exit(1)
	}
	return size, nil
}

func (table PlaylistRepo) GetSizeOfCurrentUser(userId int64) (int64, error) {
	query := `SELECT count(*) From userPlaylist WHERE user_id = user_id;`
	size := int64(0)
	if err := table.Sqlx.Get(&size, query, userId); err != nil {
		os.Exit(1)
	}
	return size, nil
}

func (table PlaylistRepo) AddToPlaylist(userId int64, playlistId int64, trackId int64) error {
	selected, err := table.SelectByIDOfCurrentUser(userId, playlistId)
	if err != nil {
		return err
	}
	query := `INSERT INTO playlistTrack(playlist_id, track_id) 
			  VALUES ($1, $2)
			  `

	if _, err := table.Sqlx.Exec(query, selected.Id, trackId); err != nil {
		return err
	}

	return nil
}

func (table PlaylistRepo) RemoveFromPlaylist(userId int64, playlistId int64, trackId int64) error {
	selected, err := table.SelectByIDOfCurrentUser(userId, playlistId)
	if err != nil {
		return err
	}

	query := `DELETE FROM playlistTrack
			  WHERE playlistTrack.playlist_id = $1 and playlistTrack.track_id = $2 
			  `

	if _, err := table.Sqlx.Exec(query, selected.Id, trackId); err != nil {
		return err
	}

	return nil
}
