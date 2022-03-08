package db

import (
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/db/models"
	"math"
	"sort"
	"sync"
)

type AlbumRep interface {
	Insert(album *models.Album) error
	Update(album *models.Album) error
	Delete(id uint64) error
	SelectByID(id uint64) (*models.Album, error)
	GetAllAlbums() (*[]models.Album, error)
	//SelectByParam(count uint64, from uint64) ([]*models.Album, error)
	//SelectByTitle(title string) (*models.Album, error)
	//SelectByAuthor(author string) (*[]models.Album, error)
}

type ArtistRep interface {
	Insert(arist *models.Artist) error
	Update(arist *models.Artist) error
	Delete(id uint64) error
	SelectByID(id uint64) (*models.Artist, error)
	GetAllArtists() (*[]models.Artist, error)
	//SelectByParam(count uint64, from uint64) ([]*models.Album, error)
	//SelectByTitle(title string) (*models.Album, error)
	//SelectByAuthor(author string) (*[]models.Album, error)
}

type SongRep interface {
	Insert(song *models.Song) error
	Update(song *models.Song) error
	Delete(id uint64) error
	SelectByID(id uint64) (*models.Song, error)
	GetAllSongs() (*[]models.Song, error)
	GetPopularSongs() (*[]models.Song, error)
	//SelectByParam(count uint64, from uint64) ([]*models.Album, error)
	//SelectByTitle(title string) (*models.Album, error)
	//SelectByAuthor(author string) (*[]models.Album, error)
}

type albumStorage struct {
	Albums []models.Album `json:"albums"`
	Mutex  sync.RWMutex   `json:"mutex"`
}

type artistStorage struct {
	Artists []models.Artist `json:"artists"`
	Mutex   sync.RWMutex    `json:"mutex"`
}

type songStorage struct {
	Songs []models.Song `json:"songs"`
	Mutex sync.RWMutex  `json:"mutex"`
}

type globalStorage struct {
	AlbumStorage  albumStorage  `json:"albumStorage"`
	ArtistStorage artistStorage `json:"artistStorage"`
	SongStorage   songStorage   `json:"songStorage"`
	Mutex         sync.RWMutex  `json:"mutex"`
}

var Storage = globalStorage{}

// ------------------------------------------------------------------

func (storage *albumStorage) Insert(album *models.Album) error {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	storage.Albums = append(storage.Albums, *album)
	return nil
}

func (storage *albumStorage) Update(album *models.Album) error {
	albumFromDB, err := storage.SelectByID(album.Id)
	if err != nil {
		return err
	}
	*albumFromDB = *album
	return nil
}

func (storage *albumStorage) Delete(id uint64) error {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	if id+1 > uint64(len(storage.Albums)) {
		return errors.New(IndexOutOfRange)
	}

	storage.Albums = append(storage.Albums[:id], storage.Albums[id+1:]...)
	for i := id; i < uint64(len(storage.Albums)); i++ {
		storage.Albums[i].Id = i
	}
	return nil
}

func (storage *albumStorage) SelectByID(id uint64) (*models.Album, error) {
	storage.Mutex.RLock()
	defer storage.Mutex.RUnlock()
	if id+1 > uint64(len(storage.Albums)) {
		return nil, errors.New(IndexOutOfRange)
	}
	return &storage.Albums[id], nil
}

func (storage *albumStorage) GetAllAlbums() (*[]models.Album, error) {
	storage.Mutex.RLock()
	defer storage.Mutex.RUnlock()
	return &storage.Albums, nil
}

// ------------------------------------------------------------------

func (storage *artistStorage) Insert(artist *models.Artist) error {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	storage.Artists = append(storage.Artists, *artist)
	return nil
}

func (storage *artistStorage) Update(artist *models.Artist) error {
	artistFromDB, err := storage.SelectByID(artist.Id)
	if err != nil {
		return err
	}
	*artistFromDB = *artist
	return nil
}

func (storage *artistStorage) Delete(id uint64) error {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	if id+1 > uint64(len(storage.Artists)) {
		return errors.New(IndexOutOfRange)
	}

	storage.Artists = append(storage.Artists[:id], storage.Artists[id+1:]...)
	for i := id; i < uint64(len(storage.Artists)); i++ {
		storage.Artists[i].Id = i
	}
	return nil
}

func (storage *artistStorage) SelectByID(id uint64) (*models.Artist, error) {
	storage.Mutex.RLock()
	defer storage.Mutex.RUnlock()
	if id+1 > uint64(len(storage.Artists)) {
		return nil, errors.New(IndexOutOfRange)
	}
	return &storage.Artists[id], nil
}

func (storage *artistStorage) GetAllArtists() (*[]models.Artist, error) {
	storage.Mutex.RLock()
	defer storage.Mutex.RUnlock()
	return &storage.Artists, nil
}

// ------------------------------------------------------------------

func (storage *songStorage) Insert(song *models.Song) error {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	storage.Songs = append(storage.Songs, *song)
	return nil
}

func (storage *songStorage) Update(song *models.Song) error {
	songFromDB, err := storage.SelectByID(song.Id)
	if err != nil {
		return err
	}
	*songFromDB = *song
	return nil
}

func (storage *songStorage) Delete(id uint64) error {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	if id+1 > uint64(len(storage.Songs)) {
		return errors.New(IndexOutOfRange)
	}

	storage.Songs = append(storage.Songs[:id], storage.Songs[id+1:]...)
	for i := id; i < uint64(len(storage.Songs)); i++ {
		storage.Songs[i].Id = i
	}
	return nil
}

func (storage *songStorage) SelectByID(id uint64) (*models.Song, error) {
	storage.Mutex.RLock()
	defer storage.Mutex.RUnlock()
	if id+1 > uint64(len(storage.Songs)) {
		return nil, errors.New(IndexOutOfRange)
	}
	return &storage.Songs[id], nil
}

func (storage *songStorage) GetAllSongs() (*[]models.Song, error) {
	storage.Mutex.RLock()
	defer storage.Mutex.RUnlock()
	return &storage.Songs, nil
}

func (storage *songStorage) GetPopularSongs() (*[]models.Song, error) {
	const top = 20
	storage.Mutex.RLock()
	defer storage.Mutex.RUnlock()

	var songsPtr = make([]*models.Song, len(storage.Songs))
	for i := 0; i < len(storage.Songs); i++ {
		songsPtr[i] = &storage.Songs[i]
	}

	sort.SliceStable(songsPtr, func(i int, j int) bool {
		song1 := *(songsPtr[i])
		song2 := *(songsPtr[j])
		return song1.CountListening > song2.CountListening
	})

	topChart := make([]models.Song, uint64(math.Min(top, float64(len(storage.Songs)))))
	for i := 0; i < len(topChart); i++ {
		topChart[i] = *songsPtr[i]
	}

	return &topChart, nil
}
