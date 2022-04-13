package TrackUseCase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	utilsInterfaces "github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	"reflect"
)

//var UseCase structsUseCase.UseCase

type TrackUseCase struct {
	TrackRepo  utilsInterfaces.TrackRepoInterface
	ArtistRepo utilsInterfaces.ArtistRepoInterface
}

func MakeTrackUseCase(track utilsInterfaces.TrackRepoInterface, artist utilsInterfaces.ArtistRepoInterface) TrackUseCase {
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

	coverPath, err := track.CreatePathById(internal.PngFormat, track.AlbumId)
	if err != nil {
		return nil, err
	}

	srcPath, err := track.CreatePath(internal.Mp3Format)
	if err != nil {
		return nil, err
	}

	return &domain.TrackDataTransfer{
		Id:         track.Id,
		Title:      track.Title,
		Artist:     artist.Name,
		Cover:      coverPath,
		Src:        srcPath,
		Likes:      track.CountLikes,
		Listenings: track.CountListening,
		Duration:   track.Duration,
	}, nil
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

func (useCase TrackUseCase) GetType() reflect.Type {
	return reflect.TypeOf(domain.Track{})
}

//func (useCase TrackUseCase) GetRepo() (utilsInterfaces.RepoInterface, error) {
//	return useCase.TrackRepo, nil
//}

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
