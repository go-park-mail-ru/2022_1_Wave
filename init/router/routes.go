package router

import (
	_ "github.com/go-park-mail-ru/2022_1_Wave/docs"
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	albumDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/delivery/http"
	albumCoverDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/albumCover/delivery/http"
	artistDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/artist/delivery/http"
	authHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/auth/delivery/http"
	trackDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/track/delivery/http"
	userHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/user/delivery/http"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

func Router(e *echo.Echo) {

	api := e.Group(apiPrefix)
	v1 := api.Group(v1Prefix)

	logger.GlobalLogger.Logrus.Warnln("api version:", v1Prefix)

	SetAlbumsRoutes(v1)
	logger.GlobalLogger.Logrus.Warnln("setting albums routes")

	SetAlbumCoversRoutes(v1)
	logger.GlobalLogger.Logrus.Warnln("setting album covers routes")

	SetArtistsRoutes(v1)
	logger.GlobalLogger.Logrus.Warnln("setting artists routes")

	SetTracksRoutes(v1)
	logger.GlobalLogger.Logrus.Warnln("setting tracks routes")

	SetDocsPath(v1)
	logger.GlobalLogger.Logrus.Warnln("setting docs routes")

	SetStaticHandle(v1)
	logger.GlobalLogger.Logrus.Warnln("setting static routes")

	SetAuthRoutes(v1)
	logger.GlobalLogger.Logrus.Warnln("setting auth routes")

	SetUserRoutes(v1)
	logger.GlobalLogger.Logrus.Warnln("setting user routes")
}

// SetAlbumsRoutes albums
func SetAlbumsRoutes(apiVersion *echo.Group) {
	albumRoutes := apiVersion.Group(albumsPrefix)

	albumRoutes.GET(idEchoPattern, albumDeliveryHttp.Get)
	albumRoutes.GET(locate, albumDeliveryHttp.GetAll)
	albumRoutes.POST(locate, albumDeliveryHttp.Create)
	albumRoutes.PUT(locate, albumDeliveryHttp.Update)
	albumRoutes.GET(popularPrefix, albumDeliveryHttp.GetPopular)
	albumRoutes.DELETE(idEchoPattern, albumDeliveryHttp.Delete)
}

// SetAlbumCoversRoutes albumCovers
func SetAlbumCoversRoutes(apiVersion *echo.Group) {
	albumRoutes := apiVersion.Group(albumCoversPrefix)
	albumRoutes.GET(idEchoPattern, albumCoverDeliveryHttp.Get)
	albumRoutes.GET(locate, albumCoverDeliveryHttp.GetAll)
	albumRoutes.POST(locate, albumCoverDeliveryHttp.Create)
	albumRoutes.PUT(locate, albumCoverDeliveryHttp.Update)
	albumRoutes.DELETE(idEchoPattern, albumCoverDeliveryHttp.Delete)
}

// SetArtistsRoutes artists
func SetArtistsRoutes(apiVersion *echo.Group) {
	artistRoutes := apiVersion.Group(artistsPrefix)

	artistRoutes.GET(idEchoPattern, artistDeliveryHttp.Get)
	artistRoutes.GET(locate, artistDeliveryHttp.GetAll)
	artistRoutes.POST(locate, artistDeliveryHttp.Create)
	artistRoutes.PUT(locate, artistDeliveryHttp.Update)
	artistRoutes.GET(popularPrefix, artistDeliveryHttp.GetPopular)
	artistRoutes.GET(idEchoPattern+popularPrefix, artistDeliveryHttp.GetPopularTracks)
	artistRoutes.DELETE(idEchoPattern, artistDeliveryHttp.Delete)
}

// SetTracksRoutes songs
func SetTracksRoutes(apiVersion *echo.Group) {
	trackRoutes := apiVersion.Group(tracksPrefix)

	trackRoutes.GET(idEchoPattern, trackDeliveryHttp.Get)
	trackRoutes.GET(locate, trackDeliveryHttp.GetAll)
	trackRoutes.POST(locate, trackDeliveryHttp.Create)
	trackRoutes.PUT(locate, trackDeliveryHttp.Update)
	trackRoutes.GET(popularPrefix, trackDeliveryHttp.GetPopular)
	trackRoutes.DELETE(idEchoPattern, trackDeliveryHttp.Delete)
}

func SetUserRoutes(apiVersion *echo.Group) {
	userRoutes := apiVersion.Group(usersPrefix)

	userRoutes.GET("/:id", userHttp.Handler.GetUser)
	userRoutes.GET("/self", userHttp.Handler.GetSelfUser, authHttp.M.Auth, authHttp.M.CSRF)

	userRoutes.PATCH("/self", userHttp.Handler.UpdateSelfUser, authHttp.M.Auth, authHttp.M.CSRF)
	userRoutes.PATCH("/upload_avatar", userHttp.Handler.UploadAvatar, authHttp.M.Auth, authHttp.M.CSRF)
}

// InitAuthModule auth
func SetAuthRoutes(apiVersion *echo.Group) {
	apiVersion.POST(loginPrefix, authHttp.Handler.Login, authHttp.M.IsSession, authHttp.M.CSRF)
	apiVersion.POST(logoutPrefix, authHttp.Handler.Logout, authHttp.M.IsSession, authHttp.M.CSRF)
	apiVersion.POST(signUpPrefix, authHttp.Handler.SignUp, authHttp.M.IsSession, authHttp.M.CSRF)
	apiVersion.GET(getCSRFPrefix, authHttp.Handler.GetCSRF)
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
