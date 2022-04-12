package storage

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	albumDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/delivery/http"
	albumUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/usecase"
	albumCoverDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/albumCover/delivery/http"
	albumCoverUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/albumCover/usecase"
	artistDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/artist/delivery/http"
	artistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/artist/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/delivery/http"
	structStorageLocal "github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/storage/local"
	structStoragePostgresql "github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/storage/postgresql"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/usecase"
	trackDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/track/delivery/http"
	trackUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/track/usecase"
	"github.com/sirupsen/logrus"
	"log"
	"reflect"
	"sync"
)

const local = "local"
const database = local

func initRepo(domainType reflect.Type, concreteUseCase *structsUseCase.UseCase, concreteHandler *structsDeliveryHttp.Handler, initedStorage utilsInterfaces.GlobalStorageInterface, mutex *sync.RWMutex) {
	var err error

	var handler utilsInterfaces.HandlerInterface
	var useCase utilsInterfaces.UseCaseInterface

	// set usecase
	switch domainType {
	case domain.AlbumDomainType:
		useCase, err = concreteUseCase.SetRepo(*initedStorage.GetAlbumRepo(), mutex)
		if err != nil {
			log.Fatal(err)
		}
	case domain.AlbumCoverDomainType:
		useCase, err = concreteUseCase.SetRepo(*initedStorage.GetAlbumCoverRepo(), mutex)
		if err != nil {
			log.Fatal(err)
		}
	case domain.ArtistDomainType:
		useCase, err = concreteUseCase.SetRepo(*initedStorage.GetArtistRepo(), mutex)
		if err != nil {
			log.Fatal(err)
		}
	case domain.TrackDomainType:
		useCase, err = concreteUseCase.SetRepo(*initedStorage.GetTrackRepo(), mutex)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal(internal.BadType)
	}

	// set handler and model
	handler, err = concreteHandler.SetModel(domainType)
	if err != nil {
		log.Fatal(err)
	}

	handler, err = handler.SetUseCase(concreteUseCase, mutex)

	*concreteUseCase = useCase.(structsUseCase.UseCase)
	*concreteHandler = handler.(structsDeliveryHttp.Handler)
}

func InitStorage(quantity int, dataBaseType string) error {

	var storage utilsInterfaces.GlobalStorageInterface
	switch dataBaseType {
	case internal.Postgres:
		storage = structStoragePostgresql.Postgres{}
	case internal.Local:
		storage = structStorageLocal.LocalStorage{}
	default:
		logrus.Fatal(internal.BadType)
	}

	initedStorage, err := storage.Init(quantity)

	if err != nil {
		logrus.Fatal("error to init database")
		return err
	}

	// albums
	initRepo(domain.AlbumDomainType, &albumUseCase.UseCase, &albumDeliveryHttp.Handler, initedStorage, domain.AlbumMutex)

	// albums covers
	initRepo(domain.AlbumCoverDomainType, &albumCoverUseCase.UseCase, &albumCoverDeliveryHttp.Handler, initedStorage, domain.AlbumCoverMutex)

	// artists
	initRepo(domain.ArtistDomainType, &artistUseCase.UseCase, &artistDeliveryHttp.Handler, initedStorage, domain.ArtistMutex)

	// tracks
	initRepo(domain.TrackDomainType, &trackUseCase.UseCase, &trackDeliveryHttp.Handler, initedStorage, domain.TrackMutex)

	storage = initedStorage

	return nil
}
