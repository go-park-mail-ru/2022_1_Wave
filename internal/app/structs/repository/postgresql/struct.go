package structRepoPostgres

import (
	"database/sql"
	"errors"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	utilsInterfaces "github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	domainCreator "github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/domain"
	"github.com/jmoiron/sqlx"
	"os"
	"reflect"
	"sync"
)

type Table struct {
	Sqlx *sqlx.DB
	name string
}

// ----------------------------------------------------------------------
func (table Table) Insert(dom utilsInterfaces.Domain, mutex *sync.RWMutex) (utilsInterfaces.RepoInterface, error) {
	mutex.Lock()

	defer mutex.Unlock()

	id, err := table.GetSize(mutex)
	if err != nil {
		return nil, err
	}
	dom, err = dom.SetId(id + 1)
	if err != nil {
		return nil, err
	}

	query := ""

	var holder interface{}
	switch table.GetTableName() {
	case constants.Album:
		holder = dom.(domain.Album)
		query = `
		INSERT INTO album (title, artist_id, count_likes, count_listening, date)
		VALUES (:title, :artist_id, :count_likes, :count_listening, :date)
		RETURNING id`
	case constants.Artist:
		holder = dom.(domain.Artist)
		query = `
		INSERT INTO artist (name, count_followers, count_listening)
		VALUES (:name, :count_followers, :count_listening)
		RETURNING id`
	case constants.Track:
		holder = dom.(domain.Track)
		query = `
		INSERT INTO track (album_id, artist_id, title, duration, count_likes, count_listening)
		VALUES (:album_id, :artist_id, :title, :duration, :count_likes, :count_listening)
		RETURNING id`
	case constants.AlbumCover:
		holder = dom.(domain.AlbumCover)
		query = `
		INSERT INTO albumCover (quote, is_dark)
		VALUES (:quote, :is_dark)
		RETURNING id`
	default:
		return table, errors.New(constants.BadType)
	}

	// do query
	res, err := table.Sqlx.PrepareNamed(query)
	if err != nil {
		return table, err
	}
	var lastId uint64
	err = res.Get(&lastId, holder)

	if err := res.Close(); err != nil {
		return table, err
	}

	if err != nil {
		return table, err
	}
	return table, nil
}

func (table Table) Update(dom utilsInterfaces.Domain, mutex *sync.RWMutex) (utilsInterfaces.RepoInterface, error) {
	mutex.Lock()
	defer mutex.Unlock()

	lastId, err := table.GetLastId(mutex)
	if err != nil {
		return table, err
	}

	id := dom.GetId()

	if id > lastId {
		return table, errors.New(constants.IndexOutOfRange)
	}

	query := ""

	var holder interface{}
	switch table.GetTableName() {
	case constants.Album:
		holder = dom.(domain.Album)
		query = `
		UPDATE album SET title=:title, artist_id=:artist_id, count_likes=:count_likes, count_listening=:count_listening, date=:date
		WHERE id = :id`
	case constants.Artist:
		holder = dom.(domain.Artist)
		query = `
		UPDATE artist SET name=:name, count_followers=:count_followers, count_listening=:count_listening
		WHERE id = :id`
	case constants.Track:
		holder = dom.(domain.Track)
		query = `
		UPDATE track
		SET album_id=:album_id, artist_id=:artist_id, title=:title, duration=:duration, count_likes=:count_likes, count_listening=:count_listening
		WHERE id = :id`
	case constants.AlbumCover:
		holder = dom.(domain.AlbumCover)
		query = `
		UPDATE albumcover
		SET quote=:quote, is_dark=:is_dark
		WHERE id = :id`
	default:
		return table, errors.New(constants.BadType)
	}

	res, err := table.Sqlx.NamedExec(query, holder)
	if err != nil {
		return table, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return table, err
	}

	if rowsAffected == 0 {
		return table, errors.New(constants.ErrorNothingToUpdate)
	}

	return table, nil
}

func (table Table) Delete(id uint64, mutex *sync.RWMutex) (utilsInterfaces.RepoInterface, error) {
	mutex.Lock()
	defer mutex.Unlock()

	query := ""
	switch table.GetTableName() {
	case constants.Album:
		query = `DELETE FROM album WHERE id = $1`
	case constants.Artist:
		query = `DELETE FROM artist WHERE id = $1`
	case constants.Track:
		query = `DELETE FROM track WHERE id = $1`
	case constants.AlbumCover:
		query = `DELETE FROM albumcover WHERE id = $1`
	default:
		return table, errors.New(constants.BadType)
	}

	res, err := table.Sqlx.Exec(query, id)

	if err != nil {
		return table, err
	}

	deleted, err := res.RowsAffected()

	if err != nil {
		return table, err
	}

	if deleted == 0 {
		return table, errors.New(constants.IndexOutOfRange)
	}
	return table, nil
}

func (table Table) SelectByID(id uint64, mutex *sync.RWMutex) (utilsInterfaces.Domain, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	lastId, err := table.GetLastId(mutex)
	if err != nil {
		return nil, err
	}
	if id > lastId {
		return nil, errors.New(constants.IndexOutOfRange)
	}

	query := ""
	//var holder interface{}
	//holder, err := domainCreator.ToDomainPtrByTableName(table.name)

	switch table.GetTableName() {
	case constants.Album:
		query = `SELECT * FROM album WHERE id = $1;`
		holder := domain.Album{}
		err = table.Sqlx.Get(&holder, query, id)
		if err != nil {
			return nil, err
		}
		return holder, nil
	case constants.Artist:
		query = `SELECT * FROM artist WHERE id = $1;`
		holder := domain.Artist{}
		err = table.Sqlx.Get(&holder, query, id)
		if err != nil {
			return nil, err
		}
		return holder, nil
	case constants.Track:
		query = `SELECT * FROM track WHERE id = $1;`
		holder := domain.Track{}
		err = table.Sqlx.Get(&holder, query, id)
		if err != nil {
			return nil, err
		}
		temp, ok := holder.AlbumId.(int64)
		if !ok {
			return nil, err
		}
		holder.AlbumId = uint64(temp)
		return holder, nil
	case constants.AlbumCover:
		query = `SELECT * FROM albumcover WHERE id = $1;`
		holder := domain.AlbumCover{}
		err = table.Sqlx.Get(&holder, query, id)
		if err != nil {
			return nil, err
		}
		return holder, nil
	default:
		return nil, errors.New(constants.BadType)
	}

	//if err := table.Sqlx.Get(holder, query, id); err != nil {
	//	return nil, err
	//}

	//return holder.(utilsInterfaces.Domain), nil

	//if err != nil {
	//	return nil, err
	//}

	//return *holder.(utilsInterfaces.Domain), nil
}

func (table Table) getManyObjects(query string, args ...interface{}) ([]utilsInterfaces.Domain, error) {
	var holder interface{}
	err := domainCreator.ToDomainsArrayPtr(&holder, table.GetTableName())
	if err != nil {
		return nil, err
	}

	err = table.Sqlx.Select(holder, query, args...)
	if err != nil {
		return nil, err
	}

	values, err := domainCreator.GetValues(holder, table.GetTableName())

	return values, nil

	// it's doesn't work :( what a pity
	//var holder interface{}
	//if err := tools.Converter.ToDomainsArrayPtr(&holder, table.Name); err != nil {
	//	return nil, err
	//}
	//var result []utilsInterfaces.Domain
	//if err := table.Sqlx.Select(holder, query); err != nil {
	//	return nil, err
	//}
	//return holder.(*[]utilsInterfaces.Domain), nil

}

func (table Table) GetAll(mutex *sync.RWMutex) ([]utilsInterfaces.Domain, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	query := ""
	switch table.GetTableName() {
	case constants.Album:
		query = `SELECT * FROM album ORDER BY id;`
	case constants.Artist:
		query = `SELECT * FROM artist ORDER BY id;`
	case constants.Track:
		query = `SELECT * FROM track ORDER BY id;`
	case constants.AlbumCover:
		query = `SELECT * FROM albumcover ORDER BY id;`
	default:
		return nil, errors.New(constants.BadType)
	}

	manyObj, err := table.getManyObjects(query)

	if err != nil {
		return nil, err
	}

	return manyObj, nil
}

func (table Table) GetPopular(mutex *sync.RWMutex) ([]utilsInterfaces.Domain, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	query := ""
	switch table.GetTableName() {
	case constants.Album:
		query = `
			SELECT *
			FROM album
			ORDER BY count_listening DESC
			LIMIT $1;`
	case constants.Artist:
		query = `
			SELECT *
			FROM artist
			ORDER BY count_listening DESC
			LIMIT $1;`
	case constants.Track:
		query = `
			SELECT *
			FROM track
			ORDER BY count_listening DESC
			LIMIT $1;`
	case constants.AlbumCover:
		query = `
			SELECT *
			FROM albumcover
			ORDER BY count_listening DESC
			LIMIT $1;`
	default:
		return nil, errors.New(constants.BadType)
	}

	manyObj, err := table.getManyObjects(query, constants.Top)

	if err != nil {
		return nil, err
	}

	return manyObj, nil
}

func (table Table) GetLastId(mutex *sync.RWMutex) (uint64, error) {
	//mutex.RLock()
	//defer mutex.RUnlock()

	query := ""
	switch table.GetTableName() {
	case constants.Album:
		query = `SELECT max(id) from album;`
	case constants.Artist:
		query = `SELECT max(id) from artist;`
	case constants.Track:
		query = `SELECT max(id) from track;`
	case constants.AlbumCover:
		query = `SELECT max(id) from albumcover;`
	default:
		return 0, errors.New(constants.BadType)
	}

	lastId := sql.NullInt64{}
	err := table.Sqlx.Get(&lastId, query)

	if err != nil {
		return 0, err
	}

	if !lastId.Valid {
		return constants.NullId, errors.New(constants.ErrorDbIsEmpty)
	}

	return uint64(lastId.Int64), nil
}

func (table Table) GetTableName() string {
	return table.name
}
func (table Table) SetTableName(name string) (Table, error) {
	table.name = name
	return table, nil
}

func (table Table) GetType(mutex *sync.RWMutex) reflect.Type {
	mutex.RLock()
	defer mutex.RUnlock()
	return reflect.TypeOf(table)
}

func (table Table) GetSize(mutex *sync.RWMutex) (uint64, error) {
	//mutex.RLock()
	//defer mutex.RUnlock()

	query := ""
	switch table.GetTableName() {
	case constants.Album:
		query = `SELECT count(*) From album;`
	case constants.Artist:
		query = `SELECT count(*) From artist;`
	case constants.Track:
		query = `SELECT count(*) From track;`
	case constants.AlbumCover:
		query = `SELECT count(*) From albumcover;`
	default:
		return 0, errors.New(constants.BadType)
	}

	size := uint64(0)
	if err := table.Sqlx.Get(&size, query); err != nil {
		os.Exit(1)
	}

	return size, nil
}

// todo пока кастыль, так как не успеваем
func (table Table) GetTracksFromAlbum(albumId uint64, mutex *sync.RWMutex) (interface{}, error) {
	if table.GetTableName() != constants.Album {
		return nil, errors.New(constants.BadType)
	}

	mutex.RLock()
	defer mutex.RUnlock()

	var tracks []domain.Track
	if err := table.Sqlx.Select(&tracks, `SELECT * FROM track WHERE album_id = $1`, albumId); err != nil {
		return nil, err
	}

	return tracks, nil

}

// todo пока кастыль, так как не успеваем
func (table Table) GetAlbumsFromArtist(artistId uint64, mutex *sync.RWMutex) (interface{}, error) {
	if table.GetTableName() != constants.Artist {
		return nil, errors.New(constants.BadType)
	}

	mutex.RLock()
	defer mutex.RUnlock()

	var albums []domain.Album
	if err := table.Sqlx.Select(&albums, `SELECT * FROM album WHERE artist_id = $1`, artistId); err != nil {
		return nil, err
	}

	return albums, nil

}

// todo пока кастыль, так как не успеваем
func (table Table) GetPopularTracksFromArtist(artistId uint64, mutex *sync.RWMutex) (interface{}, error) {
	if table.GetTableName() != constants.Artist {
		return nil, errors.New(constants.BadType)
	}

	mutex.RLock()
	defer mutex.RUnlock()

	var tracks []domain.Track
	if err := table.Sqlx.Select(&tracks, `
			SELECT * FROM track
			WHERE artist_id = $1
			ORDER BY count_listening DESC
			LIMIT $2;`, int(artistId), constants.Top); err != nil {
		return nil, err
	}

	return tracks, nil

}

// ----------------------------------------------------------------------
