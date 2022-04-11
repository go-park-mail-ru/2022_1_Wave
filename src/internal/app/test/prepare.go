package test

import (
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/init/storage"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	albumDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/delivery/http"
	albumCoverDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/albumCover/delivery/http"
	artistDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/artist/delivery/http"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	utilsInterfaces "github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	trackDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/track/delivery/http"
	"sync"
)

var tester HandlerTester

type AlbumTestCreator struct{}

func (creator AlbumTestCreator) PrepareOneTestDomain() utilsInterfaces.Domain {
	return domain.Album{
		Id:             1,
		Title:          "testedAlbum",
		ArtistId:       7,
		CountLikes:     500,
		CountListening: 10000,
		Date:           0,
	}
}

type AlbumCoverTestCreator struct{}

func (creator AlbumCoverTestCreator) PrepareOneTestDomain() utilsInterfaces.Domain {
	return domain.AlbumCover{
		Id:     3,
		Title:  "testedAlbum",
		Quote:  "tested quote",
		IsDark: true,
	}
}

type ArtistTestCreator struct{}

func (creator ArtistTestCreator) PrepareOneTestDomain() utilsInterfaces.Domain {
	return domain.Artist{
		Id:             5,
		Name:           "testArtist",
		CountFollowers: 1355,
		CountListening: 5000,
	}
}

type TrackTestCreator struct{}

func (creator TrackTestCreator) PrepareOneTestDomain() utilsInterfaces.Domain {
	return domain.Track{
		Id:             7,
		AlbumId:        uint64(5),
		ArtistId:       3,
		Title:          "testTrack",
		Duration:       300,
		CountLikes:     5050,
		CountListening: 228,
	}
}

var Mutex sync.Mutex

func InitTestDb(kindOf string, dataBaseType string) error {
	const testDataBaseSize = 20

	err := storage.InitStorage(testDataBaseSize, dataBaseType)
	if err != nil {
		return err
	}

	tester = HandlerTester{}

	switch kindOf {
	case constants.Album:
		tester, err = tester.SetHandler(albumDeliveryHttp.Handler)
	case constants.AlbumCover:
		tester, err = tester.SetHandler(albumCoverDeliveryHttp.Handler)
	case constants.Artist:
		tester, err = tester.SetHandler(artistDeliveryHttp.Handler)
	case constants.Track:
		tester, err = tester.SetHandler(trackDeliveryHttp.Handler)
	default:
		return errors.New(constants.BadType)
	}

	return err
}

//func PreparePostgres(kindOf string) (utilsInterfaces.GlobalStorageInterface, error) {
//	return initTestDb(kindOf)
//}
//
//func PrepareLocal(kindOf string) (utilsInterfaces.GlobalStorageInterface, error)) {
//	return initTestDb(kindOf)
//}

//var Albums = []utilsInterfaces.Domain{
//	domain.Album{
//		Id:             0,
//		Title:          "album0",
//		ArtistId:       0,
//		CountLikes:     500,
//		CountListening: 5000,
//		Date:           0,
//	},
//	domain.Album{
//		Id:             1,
//		Title:          "album1",
//		ArtistId:       1,
//		CountLikes:     600,
//		CountListening: 6000,
//		Date:           0,
//	},
//	domain.Album{
//		Id:             2,
//		Title:          "album2",
//		ArtistId:       2,
//		CountLikes:     1500,
//		CountListening: 15000,
//		Date:           0,
//	},
//	domain.Album{
//		Id:             3,
//		Title:          "album3",
//		ArtistId:       3,
//		CountLikes:     4500,
//		CountListening: 45000,
//		Date:           0,
//	},
//	domain.Album{
//		Id:             4,
//		Title:          "album4",
//		ArtistId:       4,
//		CountLikes:     4500,
//		CountListening: 45000,
//		Date:           0,
//	},
//}
//
//var AlbumCovers = []utilsInterfaces.Domain{
//	domain.AlbumCover{
//		Id:     0,
//		Title:  "albumCoverTitle0",
//		Quote:  "albumCoverQuote0",
//		IsDark: false,
//	},
//	domain.AlbumCover{
//		Id:     1,
//		Title:  "albumCoverTitle1",
//		Quote:  "albumCoverQuote1",
//		IsDark: false,
//	},
//	domain.AlbumCover{
//		Id:     2,
//		Title:  "albumCoverTitle2",
//		Quote:  "albumCoverQuote2",
//		IsDark: false,
//	},
//	domain.AlbumCover{
//		Id:     3,
//		Title:  "albumCoverTitle3",
//		Quote:  "albumCoverQuote3",
//		IsDark: false,
//	},
//	domain.AlbumCover{
//		Id:     4,
//		Title:  "albumCoverTitle4",
//		Quote:  "albumCoverQuote4",
//		IsDark: false,
//	},
//}
//
//var Artists = []utilsInterfaces.Domain{
//	domain.Artist{
//		Id:             0,
//		Name:           "artist0",
//		CountFollowers: 500,
//		CountListening: 5000,
//	},
//	domain.Artist{
//		Id:             1,
//		Name:           "artist1",
//		CountFollowers: 1500,
//		CountListening: 15000,
//	},
//	domain.Artist{
//		Id:             2,
//		Name:           "artist2",
//		CountFollowers: 2500,
//		CountListening: 25000,
//	},
//	domain.Artist{
//		Id:             3,
//		Name:           "artist3",
//		CountFollowers: 3500,
//		CountListening: 35000,
//	},
//	domain.Artist{
//		Id:             4,
//		Name:           "artist4",
//		CountFollowers: 4500,
//		CountListening: 45000,
//	},
//}
//
//var Tracks = []utilsInterfaces.Domain{
//	domain.Track{
//		Id:             0,
//		AlbumId:        nil,
//		ArtistId:       0,
//		Title:          "track0",
//		Duration:       0,
//		CountLikes:     500,
//		CountListening: 5000,
//	},
//	domain.Track{
//		Id:             1,
//		Title:          "track1",
//		AlbumId:        1,
//		ArtistId:       1,
//		CountLikes:     600,
//		CountListening: 6000,
//		Duration:       0,
//	},
//	domain.Track{
//		Id:             2,
//		Title:          "track2",
//		ArtistId:       2,
//		CountLikes:     1500,
//		CountListening: 15000,
//		Duration:       0,
//	},
//	domain.Track{
//		Id:             3,
//		Title:          "track3",
//		ArtistId:       3,
//		CountLikes:     4500,
//		CountListening: 45000,
//		Duration:       0,
//	},
//	domain.Track{
//		Id:             4,
//		Title:          "track4",
//		ArtistId:       4,
//		CountLikes:     4500,
//		CountListening: 45000,
//		Duration:       0,
//	},
//}
//
//func InitAlbumTestDb(t *testing.T) error {
//	//const testDataBaseSize = 20
//
//	//ctrl := gomock.NewController(t)
//	//defer ctrl.Finish()
//
//	//storageMock := mocks.NewMockGlobalStorageInterface(ctrl)
//	//albumRepoMock := mocks.NewMockRepoInterface(ctrl)
//	//artistRepoMock := mocks.NewMockRepoInterface(ctrl)
//	//albumCoverRepoMock := mocks.NewMockRepoInterface(ctrl)
//	//trackRepoMock := mocks.NewMockRepoInterface(ctrl)
//
//	//useCaseMock := mocks.NewMockUseCaseInterface(ctrl)
//	//handlerMock := mocks.NewMockHandlerInterface(ctrl)
//	//
//	//handlerMock.EXPECT().GetModel().Return(domain.AlbumDomainType, nil)
//	//handlerMock.EXPECT().GetUseCase().Return(useCaseMock, nil)
//	//handlerMock.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil)
//
//	//storageMock.EXPECT().Init(5).Return(structStorageLocal.LocalStorage{
//	//	AlbumRepo: structRepoLocal.Repo{
//	//		Domains: Albums,
//	//	},
//	//	AlbumCoverRepo: structRepoLocal.Repo{
//	//		Domains: AlbumCovers,
//	//	},
//	//	ArtistRepo: structRepoLocal.Repo{
//	//		Domains: Artists,
//	//	},
//	//	TrackRepo: structRepoLocal.Repo{
//	//		Domains: Tracks,
//	//	},
//	//})
//
//	var err error
//
//	var useCase utilsInterfaces.UseCaseInterface
//
//	useCase, err = structsUseCase.UseCase{}.SetRepo(structRepoLocal.Repo{Domains: Albums}, domain.AlbumMutex)
//	albumUseCase.UseCase = useCase.(structsUseCase.UseCase)
//	if err != nil {
//		return err
//	}
//
//	useCase, err = structsUseCase.UseCase{}.SetRepo(structRepoLocal.Repo{Domains: AlbumCovers}, domain.AlbumCoverMutex)
//	albumCoverUseCase.UseCase = useCase.(structsUseCase.UseCase)
//	if err != nil {
//		return err
//	}
//
//	useCase, err = structsUseCase.UseCase{}.SetRepo(structRepoLocal.Repo{Domains: Artists}, domain.ArtistMutex)
//	artistUseCase.UseCase = useCase.(structsUseCase.UseCase)
//	if err != nil {
//		return err
//	}
//
//	useCase, err = structsUseCase.UseCase{}.SetRepo(structRepoLocal.Repo{Domains: Tracks}, domain.TrackMutex)
//	trackUseCase.UseCase = useCase.(structsUseCase.UseCase)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
