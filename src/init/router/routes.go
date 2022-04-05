package router

import (
	_ "github.com/go-park-mail-ru/2022_1_Wave/docs"
	albumDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/delivery/http"
	artistDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/artist/delivery/http"
	trackDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/track/delivery/http"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

func Router(e *echo.Echo) {

	api := e.Group(apiPrefix)
	v1 := api.Group(v1Prefix)

	SetAlbumsRoutes(v1)
	SetArtistsRoutes(v1)
	SetTracksRoutes(v1)
	SetAuthRoutes(v1)
	SetDocsPath(v1)
	SetStaticHandle(v1)
}

// SetAlbumsRoutes albums
func SetAlbumsRoutes(apiVersion *echo.Group) {
	albumRoutes := apiVersion.Group(albumsPrefix)

	albumRoutes.GET(idEchoPattern, albumDeliveryHttp.Get)
	albumRoutes.GET(locate, albumDeliveryHttp.GetAll)
	albumRoutes.POST(locate, albumDeliveryHttp.Create)
	albumRoutes.PUT(locate, albumDeliveryHttp.Update)
	albumRoutes.GET(popularPrefix, albumDeliveryHttp.GetPopular)
	albumRoutes.DELETE(locate, albumDeliveryHttp.Delete)
}

// SetArtistsRoutes artists
func SetArtistsRoutes(apiVersion *echo.Group) {
	artistRoutes := apiVersion.Group(artistsPrefix)

	artistRoutes.GET(idEchoPattern, artistDeliveryHttp.Get)
	artistRoutes.GET(locate, artistDeliveryHttp.GetAll)
	artistRoutes.POST(locate, artistDeliveryHttp.Create)
	artistRoutes.PUT(locate, artistDeliveryHttp.Update)
	artistRoutes.GET(popularPrefix, artistDeliveryHttp.GetPopular)
	artistRoutes.DELETE(locate, artistDeliveryHttp.Delete)
}

// SetTracksRoutes songs
func SetTracksRoutes(apiVersion *echo.Group) {
	trackRoutes := apiVersion.Group(tracksPrefix)

	trackRoutes.GET(idEchoPattern, trackDeliveryHttp.Get)
	trackRoutes.GET(locate, trackDeliveryHttp.GetAll)
	trackRoutes.POST(locate, trackDeliveryHttp.Create)
	trackRoutes.PUT(locate, trackDeliveryHttp.Update)
	trackRoutes.GET(popularPrefix, trackDeliveryHttp.GetPopular)
	trackRoutes.DELETE(locate, trackDeliveryHttp.Delete)
}

// SetAuthRoutes auth
func SetAuthRoutes(apiVersion *echo.Group) {
	//router.HandleFunc(LoginUrl, middleware.CSRF(middleware.NotAuth(Login))).Methods(http.MethodPost)
	//router.HandleFunc(LogoutUrl, middleware.CSRF(middleware.Auth(Logout))).Methods(http.MethodPost)
	//router.HandleFunc(SignUpUrl, middleware.CSRF(middleware.NotAuth(SignUp))).Methods(http.MethodPost)
	//router.HandleFunc(GetUserUrl, GetUser).Methods(http.MethodGet)
	//router.HandleFunc(GetSelfUserUrl, middleware.CSRF(middleware.Auth(GetSelfUser))).Methods(http.MethodGet)
	//router.HandleFunc(GetCSRFAuthUrl, GetCSRF).Methods(http.MethodGet)
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
	currentApiVersion = v1Locate
	apiPath           = apiLocate + currentApiVersion
)

// prefixes
const (
	apiPrefix     = "/api"
	v1Prefix      = "/v1"
	albumsPrefix  = "/albums"
	artistsPrefix = "/artists"
	tracksPrefix  = "/tracks"
	usersPrefix   = "/users"
	docsPrefix    = "/docs"
	popularPrefix = "/popular"
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

// TODO проблема с джсоном, чекнуть хендлеры

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
