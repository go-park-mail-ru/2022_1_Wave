package TrackUseCase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
)

//var UseCase structsUseCase.UseCase

type TrackUseCase struct {
	TrackRepo  domain.TrackRepo
	ArtistRepo domain.ArtistRepo
}
type TrackUseCaseInterface interface {
	CastToDTO(track domain.Track) (*domain.TrackDataTransfer, error)
	GetAll() ([]domain.TrackDataTransfer, error)
	GetLastId() (id int, err error)
	Create(dom domain.Track) error
	Update(dom domain.Track) error
	Delete(id int) error
	GetById(id int) (*domain.TrackDataTransfer, error)
	GetPopular() ([]domain.TrackDataTransfer, error)
	GetTracksFromAlbum(albumId int) ([]domain.TrackDataTransfer, error)
	GetPopularTracksFromArtist(artistId int) ([]domain.TrackDataTransfer, error)
	GetSize() (int, error)
}

func MakeTrackUseCase(track domain.TrackRepo, artist domain.ArtistRepo) TrackUseCase {
	return TrackUseCase{
		TrackRepo:  track,
		ArtistRepo: artist,
	}
}

func (useCase TrackUseCase) CastToDTO(track domain.Track) (*domain.TrackDataTransfer, error) {
	artist, err := useCase.ArtistRepo.SelectByID(track.ArtistId)
	if err != nil {
		return nil, err
	}

	trackDto, err := track.CastToDtoWithoutArtistName()
	if err != nil {
		return nil, err
	}

	trackDto.Artist = artist.Name
	return trackDto, nil
}

func (useCase TrackUseCase) GetAll() ([]domain.TrackDataTransfer, error) {
	tracks, err := useCase.TrackRepo.GetAll()

	if err != nil {
		return nil, err
	}

	dto := make([]domain.TrackDataTransfer, len(tracks))

	for idx, obj := range tracks {
		result, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = *result
	}

	return dto, nil
}

func (useCase TrackUseCase) GetLastId() (id int, err error) {
	return useCase.TrackRepo.GetLastId()
}

func (useCase TrackUseCase) Create(dom domain.Track) error {
	return useCase.TrackRepo.Insert(dom)
}

func (useCase TrackUseCase) Update(dom domain.Track) error {
	return useCase.TrackRepo.Update(dom)
}

func (useCase TrackUseCase) Delete(id int) error {
	return useCase.TrackRepo.Delete(id)
}

func (useCase TrackUseCase) GetById(id int) (*domain.TrackDataTransfer, error) {
	track, err := useCase.TrackRepo.SelectByID(id)
	if err != nil {
		return nil, err
	}

	dto, err := useCase.CastToDTO(*track)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (useCase TrackUseCase) GetPopular() ([]domain.TrackDataTransfer, error) {
	popular, err := useCase.TrackRepo.GetPopular()

	if err != nil {
		return nil, err
	}

	dto := make([]domain.TrackDataTransfer, len(popular))

	for idx, obj := range popular {
		result, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = *result
	}

	return dto, nil
}

func (useCase TrackUseCase) GetTracksFromAlbum(albumId int) ([]domain.TrackDataTransfer, error) {
	tracks, err := useCase.TrackRepo.GetTracksFromAlbum(albumId)
	if err != nil {
		return nil, err
	}

	dto := make([]domain.TrackDataTransfer, len(tracks))

	for idx, obj := range tracks {
		result, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = *result
	}

	return dto, nil
}

func (useCase TrackUseCase) GetPopularTracksFromArtist(artistId int) ([]domain.TrackDataTransfer, error) {
	tracks, err := useCase.TrackRepo.GetPopularTracksFromArtist(artistId)
	if err != nil {
		return nil, err
	}

	dto := make([]domain.TrackDataTransfer, len(tracks))

	for idx, obj := range tracks {
		result, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = *result
	}

	return dto, nil
}

func (useCase TrackUseCase) GetSize() (int, error) {
	return useCase.TrackRepo.GetSize()
}
