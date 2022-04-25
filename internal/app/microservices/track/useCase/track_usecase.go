package TrackUseCase

import (
	"context"
	AlbumPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/repository"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/common/commonProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/track/trackProto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TrackUseCase struct {
	TrackRepo  *domain.TrackRepo
	ArtistRepo *domain.ArtistRepo
	trackProto.UnimplementedTrackUseCaseServer
}

func MakeTrackService(track domain.TrackRepo, artist domain.ArtistRepo) TrackUseCase {
	return TrackUseCase{
		TrackRepo:  &track,
		ArtistRepo: &artist,
	}
}

func (useCase TrackUseCase) CastToDTO(track *trackProto.Track) (*trackProto.TrackDataTransfer, error) {
	artist, err := (*useCase.ArtistRepo).SelectByID(track.ArtistId)
	if err != nil {
		return nil, err
	}

	trackDto, err := AlbumPostgres.CastTrackToDtoWithoutArtistName(track)
	if err != nil {
		return nil, err
	}

	trackDto.Artist = artist.Name
	return trackDto, nil
}

func (useCase TrackUseCase) GetAll(context.Context, *emptypb.Empty) (*trackProto.TracksResponse, error) {
	tracks, err := (*useCase.TrackRepo).GetAll()

	if err != nil {
		return nil, err
	}

	dto := make([]*trackProto.TrackDataTransfer, len(tracks))

	for idx, obj := range tracks {
		result, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}

	return &trackProto.TracksResponse{Tracks: dto}, nil
}

func (useCase TrackUseCase) GetLastId(context.Context, *emptypb.Empty) (*commonProto.IntResponse, error) {
	id, err := (*useCase.TrackRepo).GetLastId()
	if err != nil {
		return nil, err
	}

	return &commonProto.IntResponse{Data: id}, nil
}

func (useCase TrackUseCase) Create(ctx context.Context, dom *trackProto.Track) (*emptypb.Empty, error) {
	err := (*useCase.TrackRepo).Create(dom)
	return &emptypb.Empty{}, err
}

func (useCase TrackUseCase) Update(ctx context.Context, dom *trackProto.Track) (*emptypb.Empty, error) {
	err := (*useCase.TrackRepo).Update(dom)
	return &emptypb.Empty{}, err
}

func (useCase TrackUseCase) Delete(ctx context.Context, data *commonProto.IdArg) (*emptypb.Empty, error) {
	err := (*useCase.TrackRepo).Delete(data.Id)
	return &emptypb.Empty{}, err
}

func (useCase TrackUseCase) GetById(ctx context.Context, data *commonProto.IdArg) (*trackProto.TrackDataTransfer, error) {
	track, err := (*useCase.TrackRepo).SelectByID(data.Id)
	if err != nil {
		return nil, err
	}
	dto, err := useCase.CastToDTO(track)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (useCase TrackUseCase) GetPopular(context.Context, *emptypb.Empty) (*trackProto.TracksResponse, error) {
	tracks, err := (*useCase.TrackRepo).GetPopular()
	if err != nil {
		return nil, err
	}

	dto := make([]*trackProto.TrackDataTransfer, len(tracks))

	for i := 0; i < len(tracks); i++ {
		data, err := useCase.CastToDTO(tracks[i])
		if err != nil {
			return nil, err
		}
		dto[i] = data
	}

	return &trackProto.TracksResponse{Tracks: dto}, nil
}

func (useCase TrackUseCase) GetTracksFromAlbum(ctx context.Context, data *commonProto.IdArg) (*trackProto.TracksResponse, error) {
	tracks, err := (*useCase.TrackRepo).GetTracksFromAlbum(data.Id)
	if err != nil {
		return nil, err
	}

	dto := make([]*trackProto.TrackDataTransfer, len(tracks))

	for idx, obj := range tracks {
		result, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}

	return &trackProto.TracksResponse{Tracks: dto}, nil
}

func (useCase TrackUseCase) GetPopularTracksFromArtist(ctx context.Context, data *commonProto.IdArg) (*trackProto.TracksResponse, error) {
	tracks, err := (*useCase.TrackRepo).GetPopularTracksFromArtist(data.Id)
	if err != nil {
		return nil, err
	}

	dto := make([]*trackProto.TrackDataTransfer, len(tracks))

	for idx, obj := range tracks {
		result, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}

	return &trackProto.TracksResponse{Tracks: dto}, nil
}

func (useCase TrackUseCase) GetSize(context.Context, *emptypb.Empty) (*commonProto.IntResponse, error) {
	size, err := (*useCase.TrackRepo).GetSize()
	return &commonProto.IntResponse{Data: size}, err
}
