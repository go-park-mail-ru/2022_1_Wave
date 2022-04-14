package system

import (
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/init/router"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	AlbumUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/usecase"
	AlbumCoverUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/albumCover/usecase"
	ArtistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/artist/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/auth/usecase"
	utilsInterfaces "github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	structStoragePostgresql "github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/storage/postgresql"
	TrackUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/track/usecase"
	UserUsecase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/user/userUseCase"
	"github.com/labstack/echo/v4"
)

const local = "local"
const database = local

// вынести юзкейсы из свича
func Init(e *echo.Echo, quantity int, dataBaseType string) error {
	var initedStorage utilsInterfaces.GlobalStorageInterface
	var err error
	switch dataBaseType {
	case internal.Postgres:
		initedStorage, err = structStoragePostgresql.Postgres{
			Sqlx:           nil,
			SessionRepo:    nil,
			UserRepo:       nil,
			AlbumRepo:      nil,
			AlbumCoverRepo: nil,
			ArtistRepo:     nil,
			TrackRepo:      nil,
		}.Init(quantity)
		if err != nil {
			return err
		}
	default:
		return errors.New(internal.BadType)
	}

	al := initedStorage.GetAlbumRepo()
	ar := initedStorage.GetArtistRepo()
	alc := initedStorage.GetAlbumCoverRepo()
	sess := initedStorage.GetSessionRepo()
	us := initedStorage.GetUserRepo()
	tr := initedStorage.GetTrackRepo()

	auth := AuthUseCase.NewAuthUseCase(sess, us)
	album := AlbumUseCase.MakeAlbumUseCase(tr, ar, al, alc)
	albumCover := AlbumCoverUseCase.MakeAlbumCoverUseCase(alc)
	artist := ArtistUseCase.MakeArtistUseCase(ar)
	track := TrackUseCase.MakeTrackUseCase(tr, ar)
	user := UserUsecase.NewUserUseCase(us, sess)

	return router.Router(e, auth, album, albumCover, artist, track, user)

}
