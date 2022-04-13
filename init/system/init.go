package system

import (
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/init/router"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	AlbumUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/usecase"
	AlbumCoverUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/albumCover/usecase"
	ArtistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/artist/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/auth/usecase"
	structStoragePostgresql "github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/storage/postgresql"
	TrackUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/track/usecase"
	UserUsecase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/user/usecase"
	"github.com/labstack/echo/v4"
)

const local = "local"
const database = local

//func initRepo(domainType reflect.Type, initedStorage utilsInterfaces.GlobalStorageInterface) {
//	var err error
//
//	var handler utilsInterfaces.HandlerInterface
//	var useCase utilsInterfaces.UseCaseInterface
//
//	// set usecase
//	switch domainType {
//	case domain.AlbumDomainType:
//
//		useCase, err = concreteUseCase.SetRepo(*initedStorage.GetAlbumRepo())
//		if err != nil {
//			log.Fatal(err)
//		}
//	case domain.AlbumCoverDomainType:
//		useCase, err = concreteUseCase.SetRepo(*initedStorage.GetAlbumCoverRepo())
//		if err != nil {
//			log.Fatal(err)
//		}
//	case domain.ArtistDomainType:
//		useCase, err = concreteUseCase.SetRepo(*initedStorage.GetArtistRepo())
//		if err != nil {
//			log.Fatal(err)
//		}
//	case domain.TrackDomainType:
//		useCase, err = concreteUseCase.SetRepo(*initedStorage.GetTrackRepo())
//		if err != nil {
//			log.Fatal(err)
//		}
//	default:
//		log.Fatal(internal.BadType)
//	}
//
//	// set handler and model
//	handler, err = concreteHandler.SetModel(domainType)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	handler, err = handler.SetUseCase(concreteUseCase)
//
//	*concreteUseCase = useCase.(structsUseCase.UseCase)
//	*concreteHandler = handler.(structsDeliveryHttp.Handler)
//}

func Init(e *echo.Echo, quantity int, dataBaseType string) error {
	switch dataBaseType {
	case internal.Postgres:
		storage, err := structStoragePostgresql.Postgres{
			Sqlx:           nil,
			SessionRepo:    nil,
			UserRepo:       nil,
			AlbumRepo:      nil,
			AlbumCoverRepo: nil,
			ArtistRepo:     nil,
			TrackRepo:      nil,
		}.Init(quantity)
		casted := storage.(structStoragePostgresql.Postgres)
		if err != nil {
			return err
		}

		auth := AuthUseCase.NewAuthUseCase(*casted.SessionRepo, *casted.UserRepo)
		album := AlbumUseCase.MakeAlbumUseCase(*casted.TrackRepo, *casted.ArtistRepo, *casted.AlbumRepo, *casted.AlbumCoverRepo)
		albumCover := AlbumCoverUseCase.MakeAlbumCoverUseCase(*casted.AlbumCoverRepo)
		artist := ArtistUseCase.MakeArtistUseCase(*casted.ArtistRepo)
		track := TrackUseCase.MakeTrackUseCase(*casted.TrackRepo, *casted.ArtistRepo)
		user := UserUsecase.NewUserUseCase(*casted.UserRepo, *casted.SessionRepo)

		return router.Router(e, auth, album, albumCover, artist, track, user)

	default:
		return errors.New(internal.BadType)
	}
}
