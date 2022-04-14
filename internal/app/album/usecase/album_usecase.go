package AlbumUseCase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	TrackUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/track/usecase"
	"reflect"
)

type AlbumUseCase struct {
	TrackRepo      domain.TrackRepo
	ArtistRepo     domain.ArtistRepo
	AlbumRepo      domain.AlbumRepo
	AlbumCoverRepo domain.AlbumCoverRepo
}

type AlbumUseCaseInterface interface {
	CastToDTO(album domain.Album, trackUseCase TrackUseCase.TrackUseCase) (*domain.AlbumDataTransfer, error)
	GetAll(track TrackUseCase.TrackUseCase) ([]domain.AlbumDataTransfer, error)
	GetLastId() (id int, err error)
	Create(dom domain.Album) error
	Update(dom domain.Album) error
	Delete(id int) error
	GetById(track TrackUseCase.TrackUseCase, id int) (*domain.AlbumDataTransfer, error)
	GetPopular(track TrackUseCase.TrackUseCase) ([]domain.AlbumDataTransfer, error)
	GetAlbumsFromArtist(artist int, track TrackUseCase.TrackUseCase) ([]domain.AlbumDataTransfer, error)
	GetSize() (int, error)
}

func MakeAlbumUseCase(track domain.TrackRepo,
	artist domain.ArtistRepo,
	album domain.AlbumRepo,
	albumCover domain.AlbumCoverRepo) AlbumUseCase {
	return AlbumUseCase{
		TrackRepo:      track,
		ArtistRepo:     artist,
		AlbumRepo:      album,
		AlbumCoverRepo: albumCover,
	}
}

func (useCase AlbumUseCase) CastToDTO(album domain.Album, trackUseCase TrackUseCase.TrackUseCase) (*domain.AlbumDataTransfer, error) {
	artist, err := useCase.ArtistRepo.SelectByID(album.ArtistId)
	if err != nil {
		return nil, err
	}

	coverPath, err := album.CreatePath(internal.PngFormat)
	if err != nil {
		return nil, err
	}

	tracks, err := useCase.TrackRepo.GetTracksFromAlbum(album.Id)
	tracksDto := make([]domain.TrackDataTransfer, len(tracks))
	for idx, obj := range tracks {
		dto, err := trackUseCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		tracksDto[idx] = *dto
	}

	return &domain.AlbumDataTransfer{
		Id:     album.Id,
		Title:  album.Title,
		Artist: artist.Name,
		Cover:  coverPath,
		Tracks: tracksDto,
	}, nil
}

func (useCase AlbumUseCase) GetAll(track TrackUseCase.TrackUseCase) ([]domain.AlbumDataTransfer, error) {
	albums, err := useCase.AlbumRepo.GetAll()
	if err != nil {
		return nil, err
	}

	dto := make([]domain.AlbumDataTransfer, len(albums))

	for idx, obj := range albums {
		data, err := useCase.CastToDTO(obj, track)
		if err != nil {
			return nil, err
		}
		dto[idx] = *data
	}

	return dto, nil
}

func (useCase AlbumUseCase) GetLastId() (id int, err error) {
	return useCase.AlbumRepo.GetLastId()
}

func (useCase AlbumUseCase) Create(dom domain.Album) error {
	return useCase.AlbumRepo.Insert(dom)
}

func (useCase AlbumUseCase) Update(dom domain.Album) error {
	return useCase.AlbumRepo.Update(dom)
}

func (useCase AlbumUseCase) Delete(id int) error {
	return useCase.AlbumRepo.Delete(id)
}

func (useCase AlbumUseCase) GetById(track TrackUseCase.TrackUseCase, id int) (*domain.AlbumDataTransfer, error) {
	album, err := useCase.AlbumRepo.SelectByID(id)
	if err != nil {
		return nil, err
	}
	dto, err := useCase.CastToDTO(*album, track)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (useCase AlbumUseCase) GetPopular(track TrackUseCase.TrackUseCase) ([]domain.AlbumDataTransfer, error) {
	albums, err := useCase.AlbumRepo.GetPopular()
	if err != nil {
		return nil, err
	}

	dto := make([]domain.AlbumDataTransfer, len(albums))

	for idx, obj := range albums {
		data, err := useCase.CastToDTO(obj, track)
		if err != nil {
			return nil, err
		}
		dto[idx] = *data
	}

	return dto, nil
}

func (useCase AlbumUseCase) GetType() reflect.Type {
	return reflect.TypeOf(domain.Album{})
}

func (useCase AlbumUseCase) GetAlbumsFromArtist(artist int, track TrackUseCase.TrackUseCase) ([]domain.AlbumDataTransfer, error) {
	albums, err := useCase.AlbumRepo.GetAlbumsFromArtist(artist)
	if err != nil {
		return nil, err
	}

	dto := make([]domain.AlbumDataTransfer, len(albums))

	for idx, obj := range albums {
		data, err := useCase.CastToDTO(obj, track)
		if err != nil {
			return nil, err
		}
		dto[idx] = *data
	}

	return dto, nil
}

//func (useCase ArtistUseCase) SetRepo(Repo domain.RepoInterface) error {
//	useCase.Repo) = Repo
//	return useCase, nil
//}

func (useCase AlbumUseCase) GetSize() (int, error) {
	return useCase.AlbumRepo.GetSize()
}
