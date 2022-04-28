package router

import (
	_ "github.com/go-park-mail-ru/2022_1_Wave/docs"
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	albumDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/album/delivery/http"
	AlbumUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/album/useCase"
	artistDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/artist/delivery/http"
	ArtistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/artist/useCase"
	auth_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/auth"
	authHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/auth/delivery/http"
	auth_middleware "github.com/go-park-mail-ru/2022_1_Wave/internal/auth/delivery/http/middleware"
	gatewayDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/gateway/delivery/http"
	trackDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/track/delivery/http"
	TrackUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/track/useCase"
	user_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/user"
	userHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/user/delivery/http"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

func Router(e *echo.Echo,
	auth auth_domain.AuthUseCase,
	album AlbumUseCase.UseCase,
	artist ArtistUseCase.UseCase,
	track TrackUseCase.UseCase,
	user user_domain.UserUseCase) error {

	api := e.Group(apiPrefix)
	v1 := api.Group(v1Prefix)
	albumHandler := albumDeliveryHttp.MakeHandler(album, user)
	artistHandler := artistDeliveryHttp.MakeHandler(artist, track, user)
	trackHandler := trackDeliveryHttp.MakeHandler(track, user)
	authHandler := authHttp.MakeHandler(auth)
	userHandler := userHttp.MakeHandler(user)
	gatewayHandler := gatewayDeliveryHttp.MakeHandler(album, artist, track)

	m := auth_middleware.InitMiddleware(auth)

	logger.GlobalLogger.Logrus.Warnln("api version:", v1Prefix)

	SetAlbumsRoutes(v1, albumHandler)
	logger.GlobalLogger.Logrus.Warnln("setting albums routes")

	SetArtistsRoutes(v1, artistHandler)
	logger.GlobalLogger.Logrus.Warnln("setting artists routes")

	SetTracksRoutes(v1, trackHandler)
	logger.GlobalLogger.Logrus.Warnln("setting tracks routes")

	SetGatewayRoutes(v1, gatewayHandler)
	logger.GlobalLogger.Logrus.Warnln("setting gateway routes")

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
	albumRoutes.GET(favoritesPrefix, handler.GetFavorites)
	albumRoutes.POST(favoritesPrefix+idEchoPattern, handler.AddToFavorites)
	albumRoutes.DELETE(favoritesPrefix+idEchoPattern, handler.RemoveFromFavorites)
	albumRoutes.DELETE(idEchoPattern, handler.Delete)

	coverRoutes := apiVersion.Group(albumCoversPrefix)
	coverRoutes.GET(idEchoPattern, handler.GetCover)
	coverRoutes.GET(locate, handler.GetAllCovers)
	coverRoutes.POST(locate, handler.CreateCover)
	coverRoutes.PUT(locate, handler.UpdateCover)
	coverRoutes.DELETE(idEchoPattern, handler.DeleteCover)
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
	artistRoutes.GET(favoritesPrefix, handler.GetFavorites)
	artistRoutes.POST(favoritesPrefix+idEchoPattern, handler.AddToFavorites)
	artistRoutes.DELETE(favoritesPrefix+idEchoPattern, handler.RemoveFromFavorites)
	artistRoutes.DELETE(idEchoPattern, handler.Delete)
}

// SetTracksRoutes tracks
func SetTracksRoutes(apiVersion *echo.Group, handler trackDeliveryHttp.Handler) {
	trackRoutes := apiVersion.Group(tracksPrefix)

	trackRoutes.GET(idEchoPattern, handler.Get)
	trackRoutes.GET(locate, handler.GetAll)
	trackRoutes.POST(locate, handler.Create)
	trackRoutes.PUT(locate, handler.Update)
	trackRoutes.GET(popularPrefix, handler.GetPopular)
	trackRoutes.GET(favoritesPrefix, handler.GetFavorites)
	trackRoutes.POST(favoritesPrefix+idEchoPattern, handler.AddToFavorites)
	trackRoutes.DELETE(favoritesPrefix+idEchoPattern, handler.RemoveFromFavorites)
	trackRoutes.DELETE(idEchoPattern, handler.Delete)
	trackRoutes.PUT(likePrefix+idEchoPattern, handler.Like)
	trackRoutes.PUT(listenPrefix+idEchoPattern, handler.Listen)
}

// SetGatewayRoutes songs
func SetGatewayRoutes(apiVersion *echo.Group, handler gatewayDeliveryHttp.Handler) {
	searchRoutes := apiVersion.Group(searchPrefix)
	searchRoutes.GET(strEchoPattern, handler.Search)
}

func SetUserRoutes(apiVersion *echo.Group, handler userHttp.UserHandler, m *auth_middleware.HttpMiddleware) {
	userRoutes := apiVersion.Group(usersPrefix)

	userRoutes.GET("/:id", handler.GetUser)
	userRoutes.GET("/self", handler.GetSelfUser, m.Auth, m.CSRF)

	userRoutes.PATCH("/self", handler.UpdateSelfUser, m.Auth, m.CSRF)
	userRoutes.PATCH("/upload_avatar", handler.UploadAvatar, m.Auth, m.CSRF)
}

// InitAuthModule auth
func SetAuthRoutes(apiVersion *echo.Group, handler authHttp.AuthHandler, m *auth_middleware.HttpMiddleware) {
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
	searchPrefix      = "/search"
	docsPrefix        = "/docs"
	popularPrefix     = "/popular"
	favoritesPrefix   = "/favorites"
	likePrefix        = "/like"
	listenPrefix      = "/listen"
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
	login          = "login"
	logout         = "logout"
	signUp         = "signup"
	getCSRF        = "get_csrf"
	self           = "self"
	popular        = "popular"
	idMuxPattern   = "{id:[0-9]+}"
	idEchoPattern  = "/:id"
	strEchoPattern = "/:toFind"
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
