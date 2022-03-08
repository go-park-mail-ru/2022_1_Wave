package db

import (
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/db/models"
	"log"
	"math"
	"math/rand"
	"sort"
	"sync"
)

type AlbumRep interface {
	Insert(album *models.Album) error
	Update(album *models.Album) error
	Delete(id uint64) error
	SelectByID(id uint64) (*models.Album, error)
	GetAllAlbums() (*[]models.Album, error)
	GetPopularAlbums() (*[]models.Album, error)
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
	GetPopularArtists() (*[]models.Artist, error)
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

func randomRune() string {
	return string('a' + rune(rand.Intn('z'-'a'+1)))
}

func (storage *globalStorage) InitStorage() {
	const quantity = 50
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()

	albums := make([]models.Album, quantity)
	songs := make([]models.Song, quantity)
	artists := make([]models.Artist, quantity)

	const max = 10000
	// albums and artists
	for i := 0; i < quantity; i++ {
		artists[i] = models.Artist{
			Id:             uint64(i),
			Name:           randomRune(),
			Photo:          "/public/artists/photo/" + randomRune() + ".png",
			CountFollowers: uint64(rand.Int63n(max + 1)),
			CountListening: uint64(rand.Int63n(max + 1)),
		}
		albums[i] = models.Album{
			Id:             uint64(i),
			Title:          randomRune(),
			AuthorId:       uint64(rand.Int63n(quantity + 1)),
			CountLikes:     uint64(rand.Int63n(max + 1)),
			CountListening: uint64(rand.Int63n(max + 1)),
			Date:           0,
			CoverId:        uint64(rand.Int63n(max + 1)),
		}
	}

	// songs
	for i := 0; i < quantity; i++ {
		songs[i] = models.Song{
			Id:             uint64(i),
			AlbumId:        albums[rand.Intn(len(albums))].Id,
			AuthorId:       artists[rand.Intn(len(artists))].Id,
			Title:          string('a' + rune(rand.Intn('z'-'a'+1))),
			Duration:       uint64(rand.Int63n(max + 1)),
			Mp4:            "/public/songs/mp4/" + randomRune() + ".mp4",
			CountLikes:     uint64(rand.Int63n(max + 1)),
			CountListening: uint64(rand.Int63n(max + 1)),
		}
	}

	storage.ArtistStorage.Artists = artists
	storage.AlbumStorage.Albums = albums
	storage.SongStorage.Songs = songs

	log.Println("Success init local storage.")
	log.Println("Artists:", len(storage.ArtistStorage.Artists))
	log.Println("Albums:", len(storage.AlbumStorage.Albums))
	log.Println("Songs:", len(storage.SongStorage.Songs))
}

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

func (storage *albumStorage) GetPopularAlbums() (*[]models.Album, error) {
	const top = 20
	storage.Mutex.RLock()
	defer storage.Mutex.RUnlock()

	var albumsPtr = make([]*models.Album, len(storage.Albums))
	for i := 0; i < len(storage.Albums); i++ {
		albumsPtr[i] = &storage.Albums[i]
	}

	sort.SliceStable(albumsPtr, func(i int, j int) bool {
		album1 := *(albumsPtr[i])
		album2 := *(albumsPtr[j])
		return album1.CountListening > album2.CountListening
	})

	topChart := make([]models.Album, uint64(math.Min(top, float64(len(storage.Albums)))))
	for i := 0; i < len(topChart); i++ {
		topChart[i] = *albumsPtr[i]
	}

	return &topChart, nil
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

func (storage *artistStorage) GetPopularArtists() (*[]models.Artist, error) {
	const top = 20
	storage.Mutex.RLock()
	defer storage.Mutex.RUnlock()

	var artistsPtr = make([]*models.Artist, len(storage.Artists))
	for i := 0; i < len(storage.Artists); i++ {
		artistsPtr[i] = &storage.Artists[i]
	}

	sort.SliceStable(artistsPtr, func(i int, j int) bool {
		artist1 := *(artistsPtr[i])
		artist2 := *(artistsPtr[j])
		return artist1.CountListening > artist2.CountListening
	})

	topChart := make([]models.Artist, uint64(math.Min(top, float64(len(storage.Artists)))))
	for i := 0; i < len(topChart); i++ {
		topChart[i] = *artistsPtr[i]
	}
	return &topChart, nil
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
