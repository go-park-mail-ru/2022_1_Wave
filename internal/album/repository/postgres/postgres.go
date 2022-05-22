package AlbumPostgres

import (
	"errors"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	"github.com/jmoiron/sqlx"
	"time"
)

type AlbumRepo struct {
	Sqlx *sqlx.DB
}

func NewAlbumPostgresRepo(db *sqlx.DB) domain.AlbumRepo {
	return &AlbumRepo{
		Sqlx: db,
	}
}

// ----------------------------------------------------------------------
func (table AlbumRepo) Create(dom *albumProto.Album) error {
	query := `
		INSERT INTO album (title, artist_id, count_likes, count_listening, date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	// do query
	_, err := table.Sqlx.Exec(query, dom.Title, dom.ArtistId, dom.CountLikes, dom.CountListenings, dom.Date)

	return err
}

func (table AlbumRepo) Update(dom *albumProto.Album) error {
	query := `
		UPDATE album SET title = $1, artist_id = $2, count_likes = $3, count_listening = $4, date = $5
		WHERE id = $6`

	_, err := table.Sqlx.Exec(query, dom.Title, dom.ArtistId, dom.CountLikes, dom.CountListenings, dom.Date, dom.Id)
	return err
}

func (table AlbumRepo) Delete(id int64) error {
	query := `DELETE FROM album WHERE id = $1`

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

func (table AlbumRepo) SelectByID(id int64) (*albumProto.Album, error) {
	query := `SELECT * FROM album WHERE id = $1 ORDER BY id`
	holder := albumProto.Album{}
	if err := table.Sqlx.Get(&holder, query, id); err != nil {
		return nil, err
	}
	return &holder, nil

}

func (table AlbumRepo) GetAll() ([]*albumProto.Album, error) {
	query := `SELECT * FROM album ORDER BY id;`

	var albums []*albumProto.Album

	err := table.Sqlx.Select(&albums, query)

	if err != nil {
		return nil, err
	}

	return albums, nil
}

func (table AlbumRepo) GetPopular() ([]*albumProto.Album, error) {
	query := `
			SELECT *
			FROM album
			ORDER BY count_listening DESC
			LIMIT $1;`

	var albums []*albumProto.Album
	err := table.Sqlx.Select(&albums, query, constants.Top)
	if err != nil {
		return nil, err
	}

	return albums, nil
}

func (table AlbumRepo) GetLastId() (int64, error) {
	query := `SELECT max(id) from album;`

	lastId := int64(0)
	err := table.Sqlx.Get(&lastId, query)

	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (table AlbumRepo) GetSize() (int64, error) {
	query := `SELECT count(*) From album;`
	size := int64(0)
	if err := table.Sqlx.Get(&size, query); err != nil {
		return -1, err
	}
	return size, nil
}

func (table AlbumRepo) GetAlbumsFromArtist(artistId int64) ([]*albumProto.Album, error) {
	var albums []*albumProto.Album
	if err := table.Sqlx.Select(&albums, `SELECT * FROM album WHERE artist_id = $1 ORDER BY id`, artistId); err != nil {
		return nil, err
	}

	return albums, nil

}

func (table AlbumRepo) Like(albumId int64, userId int64) error {
	album, err := table.SelectByID(albumId)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO userAlbumsLike (user_id, album_id)
		VALUES ($1, $2)
		RETURNING album_id`

	if _, err := table.Sqlx.Exec(query, userId, album.Id); err != nil {
		return err
	}

	album.CountLikes = album.CountLikes + 1
	return table.Update(album)
}

func (table AlbumRepo) Listen(trackId int64) error {
	album, err := table.SelectByID(trackId)
	if err != nil {
		return err
	}
	album.CountListenings = album.CountListenings + 1
	if err := table.Update(album); err != nil {
		return err
	}
	return nil
}

func (table AlbumRepo) SearchByTitle(title string) ([]*albumProto.Album, error) {
	query := `
			SELECT *
			FROM album
			WHERE to_tsvector("title") @@ plainto_tsquery($1)
			ORDER BY ts_rank(to_tsvector("title"), plainto_tsquery($1)) DESC
			LIMIT $2;
			`

	var albums []*albumProto.Album
	err := table.Sqlx.Select(&albums, query, title, constants.SearchTop)
	if err != nil {
		return nil, err
	}

	if len(albums) == 0 {
		arg := title + "%"
		query := `
			SELECT *
			FROM album
			WHERE lower(title) LIKE lower($1)
			LIMIT $2
			`
		err := table.Sqlx.Select(&albums, query, arg, constants.SearchTop)
		if err != nil {
			return nil, err
		}
	}

	return albums, nil
}

func (table AlbumRepo) AddToFavorites(albumId int64, userId int64) error {
	album, err := table.SelectByID(albumId)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO userFavoriteAlbums (user_id, album_id)
		VALUES ($1, $2)
		RETURNING album_id`

	// do query
	_, err = table.Sqlx.Exec(query, userId, album.Id)

	return err
}

func (table AlbumRepo) GetFavorites(userId int64) ([]*albumProto.Album, error) {
	query := `SELECT id, title, artist_id, count_likes, count_listening, date FROM album
			  JOIN userFavoriteAlbums favorite ON favorite.album_id = album.id
    	      WHERE user_id = $1 ORDER BY album_id;`

	// do query
	var albums []*albumProto.Album
	err := table.Sqlx.Select(&albums, query, userId)

	return albums, err
}

func (table AlbumRepo) RemoveFromFavorites(albumId int64, userId int64) error {
	query := `DELETE FROM userFavoriteAlbums WHERE user_id = $1 and album_id = $2`

	_, err := table.Sqlx.Exec(query, userId, albumId)
	return err
}

func (table AlbumRepo) LikeCheckByUser(albumId int64, userId int64) (bool, error) {
	album, err := table.SelectByID(albumId)

	if err != nil {
		return false, err
	}

	query := `
		SELECT album_id FROM userAlbumsLike
		WHERE album_id = $1 and user_id = $2`

	likedAlbumId := -1
	err = table.Sqlx.Get(&likedAlbumId, query, album.Id, userId)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (table AlbumRepo) CountPopularAlbumOfWeek() (bool, error) {
	albums, err := table.GetAll()
	if err != nil {
		return false, err
	}
	for _, album := range albums {
		query := `SELECT album_id, last_week_likes, current_week_likes, date
			  	  FROM popularAlbumsByWeek
			  	  WHERE album_id = $1`
		holder := albumProto.PopularAlbumOfWeek{}
		err := table.Sqlx.Get(&holder, query, album.Id)
		if err != nil {
			query = `
					INSERT INTO popularAlbumsByWeek (album_id, last_week_likes, current_week_likes, date)
					VALUES ($1, $2, $3, $4)
					RETURNING album_id`
			if _, err := table.Sqlx.Exec(query, album.Id, 0, 0, time.Now().Unix()); err != nil {
				return false, err
			}
		} else {
			lastLikes := holder.CurrentWeekLikes
			currentLikes := album.CountLikes
			date := time.Now().Unix()
			query = `
					 UPDATE popularAlbumsByWeek SET last_week_likes = $1, current_week_likes = $2, date = $3
					WHERE album_id = $4`

			if _, err := table.Sqlx.Exec(query, lastLikes, currentLikes, date, album.Id); err != nil {
				return false, err
			}
		}

	}
	return true, nil
}

func (table AlbumRepo) GetPopularAlbumOfWeekTop20() ([]*albumProto.Album, error) {
	query := `
		SELECT id, title, artist_id, album.count_likes, album.count_listening, album.date FROM album
		JOIN popularAlbumsByWeek p ON p.album_id = album.id 
		ORDER BY (p.current_week_likes - p.last_week_likes) DESC
		LIMIT $1;`

	var albums []*albumProto.Album
	if err := table.Sqlx.Select(&albums, query, constants.Top); err != nil {
		return nil, err
	}

	return albums, nil
}

// ----------------------------------------------------------------------
