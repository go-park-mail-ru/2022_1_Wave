package router

import (
	_ "github.com/go-park-mail-ru/2022_1_Wave/docs"
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	albumDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/delivery/http"
	AlbumUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/usecase"
	albumCoverDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/albumCover/delivery/http"
	albumCoverUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/albumCover/usecase"
	artistDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/artist/delivery/http"
	ArtistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/artist/usecase"
	authHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/auth/delivery/http"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/auth/delivery/http/http_middleware"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	trackDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/track/delivery/http"
	TrackUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/track/usecase"
	userHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/user/delivery/http"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

func Router(e *echo.Echo,
	auth domain.AuthUseCase,
	album AlbumUseCase.AlbumUseCase,
	albumCover albumCoverUseCase.AlbumCoverUseCase,
	artist ArtistUseCase.ArtistUseCase,
	track TrackUseCase.TrackUseCase,
	user domain.UserUseCase) error {

	api := e.Group(apiPrefix)
	v1 := api.Group(v1Prefix)

	albumHandler := albumDeliveryHttp.MakeHandler(album, track)
	albumCoverHandler := albumCoverDeliveryHttp.MakeHandler(albumCover)
	artistHandler := artistDeliveryHttp.MakeHandler(artist, album, track)
	trackHandler := trackDeliveryHttp.MakeHandler(artist, track)
	authHandler := authHttp.MakeHandler(auth)
	userHandler := userHttp.MakeHandler(user)
	m := http_middleware.InitMiddleware(auth)

	logger.GlobalLogger.Logrus.Warnln("api version:", v1Prefix)

	SetAlbumsRoutes(v1, albumHandler)
	logger.GlobalLogger.Logrus.Warnln("setting albums routes")

	SetAlbumCoversRoutes(v1, albumCoverHandler)
	logger.GlobalLogger.Logrus.Warnln("setting album covers routes")

	SetArtistsRoutes(v1, artistHandler)
	logger.GlobalLogger.Logrus.Warnln("setting artists routes")

	SetTracksRoutes(v1, trackHandler)
	logger.GlobalLogger.Logrus.Warnln("setting tracks routes")

	SetDocsPath(v1)
	logger.GlobalLogger.Logrus.Warnln("setting docs routes")

	SetStaticHandle(v1)
	logger.GlobalLogger.Logrus.Warnln("setting static routes")

	SetAuthRoutes(v1, authHandler, m)
	logger.GlobalLogger.Logrus.Warnln("setting auth routes")

	SetUserRoutes(v1, userHandler, m)
	logger.GlobalLogger.Logrus.Warnln("setting user routes")

	return nil
}

// SetAlbumsRoutes albums
func SetAlbumsRoutes(apiVersion *echo.Group, handler albumDeliveryHttp.Handler) {
	albumRoutes := apiVersion.Group(albumsPrefix)

	albumRoutes.GET(idEchoPattern, handler.Get)
	albumRoutes.GET(locate, handler.GetAll)
	albumRoutes.POST(locate, handler.Create)
	albumRoutes.PUT(locate, handler.Update)
	albumRoutes.GET(popularPrefix, handler.GetPopular)
	albumRoutes.DELETE(idEchoPattern, handler.Delete)
}

// SetAlbumCoversRoutes albumCovers
func SetAlbumCoversRoutes(apiVersion *echo.Group, handler albumCoverDeliveryHttp.Handler) {
	albumRoutes := apiVersion.Group(albumCoversPrefix)
	albumRoutes.GET(idEchoPattern, handler.Get)
	albumRoutes.GET(locate, handler.GetAll)
	albumRoutes.POST(locate, handler.Create)
	albumRoutes.PUT(locate, handler.Update)
	albumRoutes.DELETE(idEchoPattern, handler.Delete)
}

// SetArtistsRoutes artists
func SetArtistsRoutes(apiVersion *echo.Group, handler artistDeliveryHttp.Handler) {
	artistRoutes := apiVersion.Group(artistsPrefix)

	artistRoutes.GET(idEchoPattern, handler.Get)
	artistRoutes.GET(locate, handler.GetAll)
	artistRoutes.POST(locate, handler.Create)
	artistRoutes.PUT(locate, handler.Update)
	artistRoutes.GET(popularPrefix, handler.GetPopular)
	artistRoutes.GET(idEchoPattern+popularPrefix, handler.GetPopularTracks)
	artistRoutes.DELETE(idEchoPattern, handler.Delete)
}

// SetTracksRoutes songs
func SetTracksRoutes(apiVersion *echo.Group, handler trackDeliveryHttp.Handler) {
	trackRoutes := apiVersion.Group(tracksPrefix)

	trackRoutes.GET(idEchoPattern, handler.Get)
	trackRoutes.GET(locate, handler.GetAll)
	trackRoutes.POST(locate, handler.Create)
	trackRoutes.PUT(locate, handler.Update)
	trackRoutes.GET(popularPrefix, handler.GetPopular)
	trackRoutes.DELETE(idEchoPattern, handler.Delete)
}

func SetUserRoutes(apiVersion *echo.Group, handler userHttp.UserHandler, m *http_middleware.HttpMiddleware) {
	userRoutes := apiVersion.Group(usersPrefix)

	userRoutes.GET("/:id", handler.GetUser)
	userRoutes.GET("/self", handler.GetSelfUser, m.Auth, m.CSRF)

	userRoutes.PATCH("/self", handler.UpdateSelfUser, m.Auth, m.CSRF)
	userRoutes.PATCH("/upload_avatar", handler.UploadAvatar, m.Auth, m.CSRF)
}

// InitAuthModule auth
func SetAuthRoutes(apiVersion *echo.Group, handler authHttp.AuthHandler, m *http_middleware.HttpMiddleware) {
	apiVersion.POST(loginPrefix, handler.Login, m.IsSession, m.CSRF)
	apiVersion.POST(logoutPrefix, handler.Logout, m.IsSession, m.CSRF)
	apiVersion.POST(signUpPrefix, handler.SignUp, m.IsSession, m.CSRF)
	apiVersion.GET(getCSRFPrefix, handler.GetCSRF)
}

// SetDocsPath docs
func SetDocsPath(apiVersion *echo.Group) {
	docRoutes := apiVersion.Group(docsPrefix)
	docRoutes.GET(locate+"*", echoSwagger.WrapHandler)
}

// SetStaticHandle static
func SetStaticHandle(apiVersion *echo.Group) {
	// /net/v1/static/img/album/123.jpg -> ./static/img/album/123.jpg
	//staticHandler := http.StripPrefix(
	//	GetStaticUrl,
	//	http.FileServer(http.Dir("./static")),
	//)
	//router.Handle(GetStaticUrl, staticHandler)
}

// config
const (
	Proto             = "http://"
	Host              = "localhost"
	redisDefaultPort  = "6379"
	currentApiVersion = v1Locate
	apiPath           = apiLocate + currentApiVersion
)

// prefixes
const (
	apiPrefix         = "/api"
	v1Prefix          = "/v1"
	albumsPrefix      = "/albums"
	albumCoversPrefix = "/albumCovers"
	artistsPrefix     = "/artists"
	tracksPrefix      = "/tracks"
	usersPrefix       = "/users"
	docsPrefix        = "/docs"
	popularPrefix     = "/popular"
	loginPrefix       = "/login"
	logoutPrefix      = "/logout"
	signUpPrefix      = "/signup"
	getCSRFPrefix     = "/get_csrf"
)

const (
	locate        = "/"
	apiLocate     = "api/"
	v1Locate      = "v1/"
	albumsLocate  = "albums/"
	artistsLocate = "artists/"
	tracksLocate  = "tracks/"
	usersLocate   = "users/"
	AssetsPrefix  = "assets/"
)

// destinations
const (
	login         = "login"
	logout        = "logout"
	signUp        = "signup"
	getCSRF       = "get_csrf"
	self          = "self"
	popular       = "popular"
	idMuxPattern  = "{id:[0-9]+}"
	idEchoPattern = "/:id"
)

// words
const (
	Get    = "Get"
	Update = "Update"
	Create = "Create"
	Delete = "Delete"
)

// albums urls
const (
	CreateAlbumUrl       = "/" + apiPath + albumsLocate
	UpdateAlbumUrl       = CreateAlbumUrl
	GetAllAlbumsUrl      = "/" + apiPath + albumsLocate
	GetAlbumUrlWithoutId = GetAllAlbumsUrl
	GetAlbumUrl          = GetAlbumUrlWithoutId + idMuxPattern
	GetPopularAlbumsUrl  = GetAllAlbumsUrl + popular
	DeleteAlbumUrl       = GetAlbumUrl
)

// artists urls
const (
	CreateArtistUrl       = "/" + apiPath + artistsLocate
	UpdateArtistUrl       = CreateArtistUrl
	GetAllArtistsUrl      = "/" + apiPath + artistsLocate
	GetArtistUrlWithoutId = GetAllAlbumsUrl
	GetArtistUrl          = GetArtistUrlWithoutId + idMuxPattern
	GetPopularArtistsUrl  = GetAllArtistsUrl + popular
	DeleteArtistUrl       = GetArtistUrl
)

// tracks urls
const (
	CreateTrackUrl       = "/" + apiPath + tracksLocate
	UpdateTrackUrl       = CreateTrackUrl
	GetAllTracksUrl      = "/" + apiPath + tracksLocate
	GetTrackUrlWithoutId = GetAllTracksUrl
	GetTrackUrl          = GetTrackUrlWithoutId + idMuxPattern
	GetPopularTracksUrl  = GetAllTracksUrl + popular
	DeleteTrackUrl       = GetTrackUrl
)

// auth urls
const (
	LoginUrl       = "/" + apiPath + login
	LogoutUrl      = "/" + apiPath + logout
	SignUpUrl      = "/" + apiPath + signUp
	GetUserUrl     = "/" + apiPath + usersLocate + idMuxPattern
	GetSelfUserUrl = "/" + apiPath + usersLocate + self
	GetCSRFAuthUrl = "/" + apiPath + getCSRF
	GetStaticUrl   = "/" + apiPath + "/static/"
)
