package structRepoPostgres

import (
	"database/sql"
	"errors"
	"fmt"
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
	Name string
}

// ----------------------------------------------------------------------
func (table Table) Insert(dom *utilsInterfaces.Domain, mutex *sync.RWMutex) (utilsInterfaces.RepoInterface, error) {
	mutex.Lock()
	defer mutex.Unlock()

	id, err := table.GetSize(mutex)
	if err != nil {
		return nil, err
	}
	*dom, err = (*dom).SetId(id + 1)
	if err != nil {
		return nil, err
	}

	query := ""

	var holder interface{}
	switch table.Name {
	case constants.Album:
		holder = (*dom).(domain.Album)
		query = fmt.Sprintf(`
		INSERT INTO %s (title, artist_id, count_likes, count_listening, date)
		VALUES (:title, :artist_id, :count_likes, :count_listening, :date)
		RETURNING id`, table.Name)
	case constants.Artist:
		holder = (*dom).(domain.Artist)
		query = fmt.Sprintf(`
		INSERT INTO %s (name, count_followers, count_listening)
		VALUES (:name, :count_followers, :count_listening)
		RETURNING id`, table.Name)
	case constants.Track:
		holder = (*dom).(domain.Track)
		query = fmt.Sprintf(`
		INSERT INTO %s (album_id, artist_id, title, duration, count_likes, count_listening)
		VALUES (:album_id, :artist_id, :title, :duration, :count_likes, :count_listening)
		RETURNING id`, table.Name)
	case constants.AlbumCover:
		holder = (*dom).(domain.AlbumCover)
		query = fmt.Sprintf(`
		INSERT INTO %s (title, quote, is_dark)
		VALUES (:title, :quote, :is_dark)
		RETURNING id`, table.Name)

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
	if err != nil {
		return table, err
	}

	return table, nil
}

func (table Table) Update(dom *utilsInterfaces.Domain, mutex *sync.RWMutex) (utilsInterfaces.RepoInterface, error) {
	mutex.Lock()
	defer mutex.Unlock()

	lastId, err := table.GetLastId(mutex)
	if err != nil {
		return table, err
	}

	id := (*dom).GetId()

	if id > lastId {
		return table, errors.New(constants.IndexOutOfRange)
	}

	query := ""

	var holder interface{}
	switch table.Name {
	case constants.Album:
		holder = (*dom).(domain.Album)
		query = fmt.Sprintf(`
		UPDATE %s SET title=:title, artist_id=:artist_id, count_likes=:count_likes, count_listening=:count_listening, date=:date
		WHERE id = %d`, table.Name, id)
	case constants.Artist:
		holder = (*dom).(domain.Artist)
		query = fmt.Sprintf(`
		UPDATE %s SET name=:name, count_followers=:count_followers, count_listening=:count_listening
		WHERE id = %d`, table.Name, id)
	case constants.Track:
		holder = (*dom).(domain.Track)
		query = fmt.Sprintf(`
		UPDATE %s
		SET album_id=:album_id, artist_id=:artist_id, title=:title, duration=:duration, count_likes=:count_likes, count_listening=:count_listening
		WHERE id = %d`, table.Name, id)
	case constants.AlbumCover:
		holder = (*dom).(domain.AlbumCover)
		query = fmt.Sprintf(`
		UPDATE %s
		SET title=:title, quote=:quote, is_dark=:is_dark
		WHERE id = %d`, table.Name, id)
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

	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, table.Name)

	res, err := table.Sqlx.Exec(query, id)

	deleted, err := res.RowsAffected()

	if err != nil {
		return table, err
	}

	if deleted == 0 {
		return table, errors.New(constants.IndexOutOfRange)
	}
	return table, nil
}

// TODO решить проблему с двойной ссылочностью
func (table Table) SelectByID(id uint64, mutex *sync.RWMutex) (*utilsInterfaces.Domain, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	lastId, err := table.GetLastId(mutex)
	if err != nil {
		return nil, err
	}
	if id > lastId {
		return nil, errors.New(constants.IndexOutOfRange)
	}

	//var holder interface{}
	//if err := domainCreator.ToDomainPtr(holder, table.Name); err != nil {
	//	return nil, err
	//}

	holder, err := domainCreator.ToDomainPtr(table.Name)

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = %d;`, table.Name, id)

	if err := table.Sqlx.Get(holder, query); err != nil {
		return nil, err
	}

	return &holder, nil
}

func (table Table) getManyObjects(query string) ([]utilsInterfaces.Domain, error) {
	var holder interface{}
	err := domainCreator.ToDomainsArrayPtr(&holder, table.Name)
	if err != nil {
		return nil, err
	}

	err = table.Sqlx.Select(holder, query)
	if err != nil {
		return nil, err
	}

	values, err := domainCreator.GetValues(holder, table.Name)

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

func (table Table) GetAll(mutex *sync.RWMutex) (*[]utilsInterfaces.Domain, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id;", table.Name)

	manyObj, err := table.getManyObjects(query)

	if err != nil {
		return nil, err
	}

	return &manyObj, nil
}

func (table Table) GetPopular(mutex *sync.RWMutex) (*[]utilsInterfaces.Domain, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	query := fmt.Sprintf(`
	SELECT *
	FROM %s
	ORDER BY %s DESC
	LIMIT %d; `, table.Name, constants.Count_listening, constants.Top)

	manyObj, err := table.getManyObjects(query)

	if err != nil {
		return nil, err
	}

	return &manyObj, nil
}

func (table Table) GetLastId(mutex *sync.RWMutex) (uint64, error) {
	//mutex.RLock()
	//defer mutex.RUnlock()

	query := fmt.Sprintf("SELECT max(id) from %s", table.Name)

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

func (table Table) GetType(mutex *sync.RWMutex) reflect.Type {
	mutex.RLock()
	defer mutex.RUnlock()
	return reflect.TypeOf(table)
}

func (table Table) GetSize(mutex *sync.RWMutex) (uint64, error) {
	//mutex.RLock()
	//defer mutex.RUnlock()

	query := fmt.Sprintf("SELECT count(*) From %s;", table.Name)

	size := uint64(0)
	if err := table.Sqlx.Get(&size, query); err != nil {
		os.Exit(1)
	}

	return size, nil
}

// ----------------------------------------------------------------------
