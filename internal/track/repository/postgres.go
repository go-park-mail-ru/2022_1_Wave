package TrackPostgres

import (
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	"github.com/jmoiron/sqlx"
	"os"
	"time"
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
			ORDER BY count_likes DESC
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
	if err := table.Sqlx.Select(&tracks, `SELECT * FROM track WHERE album_id = $1 ORDER BY count_likes DESC`, albumId); err != nil {
		return nil, err
	}
	return tracks, nil
}

func (table TrackRepo) GetPopularTracksFromArtist(artistId int64) ([]*trackProto.Track, error) {
	var tracks []*trackProto.Track
	if err := table.Sqlx.Select(&tracks, `
			SELECT * FROM track
			WHERE artist_id = $1
			ORDER BY count_likes DESC
			LIMIT $2;`, artistId, constants.Top); err != nil {
		return nil, err
	}

	return tracks, nil
}

func (table TrackRepo) Like(trackId int64, userId int64) error {
	track, err := table.SelectByID(trackId)

	if err != nil {
		return err
	}

	query := `
		INSERT INTO userTracksLike (user_id, track_id)
		VALUES ($1, $2)
		RETURNING track_id`
	if _, err := table.Sqlx.Exec(query, userId, track.Id); err != nil {
		return err
	}

	track.CountLikes = track.CountLikes + 1
	return table.Update(track)
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
			ORDER BY ts_rank(to_tsvector("title"), plainto_tsquery($1)) DESC
			LIMIT $2;
			`

	var tracks []*trackProto.Track
	err := table.Sqlx.Select(&tracks, query, title, constants.SearchTop)
	if err != nil {
		return nil, err
	}

	if len(tracks) == 0 {
		arg := title + "%"
		query := `
			SELECT *
			FROM track
			WHERE lower(title) LIKE lower($1)
			LIMIT $2;
			`
		err := table.Sqlx.Select(&tracks, query, arg, constants.SearchTop)
		if err != nil {
			return nil, err
		}
	}

	return tracks, nil
}

func (table TrackRepo) AddToFavorites(trackId int64, userId int64) error {
	if err := table.Like(trackId, userId); err != nil {
		return err
	}

	query := `
		INSERT INTO userfavoritetracks (user_id, track_id)
		VALUES ($1, $2)
		RETURNING track_id`

	_, err := table.Sqlx.Exec(query, userId, trackId)

	return err
}

func (table TrackRepo) GetFavorites(userId int64) ([]*trackProto.Track, error) {
	query := `SELECT id, album_id, artist_id, title, duration, count_likes, count_listening FROM track
			  JOIN userFavoriteTracks favorite ON favorite.track_id = track.id
    	      WHERE user_id = $1 ORDER BY track_id;`
	// do query
	var tracks []*trackProto.Track
	err := table.Sqlx.Select(&tracks, query, userId)

	return tracks, err
}

func (table TrackRepo) RemoveFromFavorites(trackId int64, userId int64) error {
	query := `DELETE FROM userTracksLike WHERE user_id = $1 AND track_id = $2`
	_, err := table.Sqlx.Exec(query, userId, trackId)
	if err != nil {
		return err
	}
	query = `DELETE FROM userFavoriteTracks WHERE user_id = $1 and track_id = $2`
	_, err = table.Sqlx.Exec(query, userId, trackId)
	return err
}

func (table TrackRepo) GetTracksFromPlaylist(playlistId int64) ([]*trackProto.Track, error) {
	query := `SELECT id, album_id, artist_id, title, duration, count_likes, count_listening FROM track
			  JOIN playlisttrack ON playlisttrack.track_id = track.id and playlisttrack.playlist_id = $1
    	      ORDER BY track.id;`
	// do query
	var tracks []*trackProto.Track
	err := table.Sqlx.Select(&tracks, query, playlistId)
	return tracks, err
}

func (table TrackRepo) LikeCheckByUser(trackId int64, userId int64) (bool, error) {
	track, err := table.SelectByID(trackId)

	if err != nil {
		return false, err
	}

	query := `
		SELECT track_id FROM userTracksLike
		WHERE track_id = $1 and user_id = $2`

	likedTrackId := -1
	err = table.Sqlx.Get(&likedTrackId, query, track.Id, userId)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (table TrackRepo) CountPopularTrackOfWeek() (bool, error) {
	tracks, err := table.GetAll()
	if err != nil {
		return false, err
	}
	for _, track := range tracks {
		query := `SELECT track_id, last_week_likes, current_week_likes, date
			  	  FROM popularTracksByWeek
			  	  WHERE track_id = $1`
		holder := trackProto.PopularTrackOfWeek{}
		err := table.Sqlx.Get(&holder, query, track.Id)
		if err != nil {
			query = `
					INSERT INTO popularTracksByWeek (track_id, last_week_likes, current_week_likes, date)
					VALUES ($1, $2, $3, $4)
					RETURNING track_id`
			if _, err := table.Sqlx.Exec(query, track.Id, 0, 0, time.Now().Unix()); err != nil {
				return false, err
			}
		} else {
			lastLikes := holder.CurrentWeekLikes
			currentLikes := track.CountLikes
			date := time.Now().Unix()
			query = `
					 UPDATE popularTracksByWeek SET last_week_likes = $1, current_week_likes = $2, date = $3
					WHERE track_id = $4`

			if _, err := table.Sqlx.Exec(query, lastLikes, currentLikes, date, track.Id); err != nil {
				return false, err
			}
		}

	}
	return true, nil
}

func (table TrackRepo) GetPopularTrackOfWeekTop20() ([]*trackProto.Track, error) {
	query := `
		SELECT id, title, artist_id, track.count_likes, track.count_listening, track.duration, track.album_id FROM track
		JOIN popularTracksByWeek p ON p.track_id = track.id 
		ORDER BY (p.current_week_likes - p.last_week_likes) DESC, count_likes DESC
		LIMIT $1;`

	var tracks []*trackProto.Track
	if err := table.Sqlx.Select(&tracks, query, constants.Top); err != nil {
		return nil, err
	}

	return tracks, nil
}
