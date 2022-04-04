package router

import (
	docs "github.com/go-park-mail-ru/2022_1_Wave/docs"
	albumDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/delivery/http"
	artistDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/artist/delivery/http"
	trackDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/track/delivery/http"
	"github.com/gorilla/mux"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func Router() (router *mux.Router) {
	router = mux.NewRouter()
	SetAlbumsRoutes(router)
	SetArtistsRoutes(router)
	SetTracksRoutes(router)
	SetAuthRoutes(router)
	SetDocsPath(router)
	SetStaticHandle(router)
	return
}

// SetAlbumsRoutes albums
func SetAlbumsRoutes(router *mux.Router) {
	router.HandleFunc(GetAllAlbumsUrl, albumDeliveryHttp.GetAll).Methods(http.MethodGet)
	router.HandleFunc(CreateAlbumUrl, albumDeliveryHttp.Create).Methods(http.MethodPost)
	router.HandleFunc(UpdateAlbumUrl, albumDeliveryHttp.Update).Methods(http.MethodPut)
	router.HandleFunc(GetAlbumUrl, albumDeliveryHttp.Get).Methods(http.MethodGet)
	router.HandleFunc(GetPopularAlbumsUrl, albumDeliveryHttp.GetPopular).Methods(http.MethodGet)
	router.HandleFunc(DeleteAlbumUrl, albumDeliveryHttp.Delete).Methods(http.MethodDelete)
}

// SetArtistsRoutes artists
func SetArtistsRoutes(router *mux.Router) {
	router.HandleFunc(GetAllArtistsUrl, artistDeliveryHttp.GetAll).Methods(http.MethodGet)
	router.HandleFunc(CreateArtistUrl, artistDeliveryHttp.Create).Methods(http.MethodPost)
	router.HandleFunc(UpdateArtistUrl, artistDeliveryHttp.Update).Methods(http.MethodPut)
	router.HandleFunc(GetArtistUrl, artistDeliveryHttp.Get).Methods(http.MethodGet)
	router.HandleFunc(GetPopularArtistsUrl, artistDeliveryHttp.GetPopular).Methods(http.MethodGet)
	router.HandleFunc(DeleteArtistUrl, artistDeliveryHttp.Delete).Methods(http.MethodDelete)
}

// SetTracksRoutes songs
func SetTracksRoutes(router *mux.Router) {
	router.HandleFunc(GetAllTracksUrl, trackDeliveryHttp.GetAll).Methods(http.MethodGet)
	router.HandleFunc(CreateTrackUrl, trackDeliveryHttp.Create).Methods(http.MethodPost)
	router.HandleFunc(UpdateTrackUrl, trackDeliveryHttp.Update).Methods(http.MethodPut)
	router.HandleFunc(GetTrackUrl, trackDeliveryHttp.Get).Methods(http.MethodGet)
	router.HandleFunc(GetPopularTracksUrl, trackDeliveryHttp.GetPopular).Methods(http.MethodGet)
	router.HandleFunc(DeleteTrackUrl, trackDeliveryHttp.Delete).Methods(http.MethodDelete)
}

// SetAuthRoutes auth
func SetAuthRoutes(router *mux.Router) {
	//router.HandleFunc(LoginUrl, middleware.CSRF(middleware.NotAuth(Login))).Methods(http.MethodPost)
	//router.HandleFunc(LogoutUrl, middleware.CSRF(middleware.Auth(Logout))).Methods(http.MethodPost)
	//router.HandleFunc(SignUpUrl, middleware.CSRF(middleware.NotAuth(SignUp))).Methods(http.MethodPost)
	//router.HandleFunc(GetUserUrl, GetUser).Methods(http.MethodGet)
	//router.HandleFunc(GetSelfUserUrl, middleware.CSRF(middleware.Auth(GetSelfUser))).Methods(http.MethodGet)
	//router.HandleFunc(GetCSRFAuthUrl, GetCSRF).Methods(http.MethodGet)
}

// SetDocsPath docs
func SetDocsPath(router *mux.Router) {
	docs.SwaggerInfo.BasePath = "/"
	router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	router.PathPrefix("/docs/*any").Handler(swaggerFiles.Handler)
}

// SetStaticHandle static
func SetStaticHandle(router *mux.Router) {
	// /net/v1/static/img/album/123.jpg -> ./static/img/album/123.jpg
	staticHandler := http.StripPrefix(
		GetStaticUrl,
		http.FileServer(http.Dir("./static")),
	)
	router.Handle(GetStaticUrl, staticHandler)
}

// config
const (
	Proto             = "http://"
	Host              = "localhost"
	currentApiVersion = v1Prefix
	apiPath           = apiPrefix + currentApiVersion
)

// prefixes
const (
	apiPrefix     = "api/"
	v1Prefix      = "v1/"
	albumsPrefix  = "albums/"
	artistsPrefix = "artists/"
	songsPrefix   = "tracks/"
	usersPrefix   = "users/"
	AssetsPrefix  = "assets/"
)

// destinations
const (
	login   = "login"
	logout  = "logout"
	signUp  = "signup"
	getCSRF = "get_csrf"
	self    = "self"
	popular = "popular"
	id      = "{id:[0-9]+}"
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
	CreateAlbumUrl       = "/" + apiPath + albumsPrefix
	UpdateAlbumUrl       = CreateAlbumUrl
	GetAllAlbumsUrl      = "/" + apiPath + albumsPrefix
	GetAlbumUrlWithoutId = GetAllAlbumsUrl
	GetAlbumUrl          = GetAlbumUrlWithoutId + id
	GetPopularAlbumsUrl  = GetAllAlbumsUrl + popular
	DeleteAlbumUrl       = GetAlbumUrl
)

// artists urls
const (
	CreateArtistUrl       = "/" + apiPath + artistsPrefix
	UpdateArtistUrl       = CreateArtistUrl
	GetAllArtistsUrl      = "/" + apiPath + artistsPrefix
	GetArtistUrlWithoutId = GetAllAlbumsUrl
	GetArtistUrl          = GetArtistUrlWithoutId + id
	GetPopularArtistsUrl  = GetAllArtistsUrl + popular
	DeleteArtistUrl       = GetArtistUrl
)

// tracks urls
const (
	CreateTrackUrl       = "/" + apiPath + songsPrefix
	UpdateTrackUrl       = CreateTrackUrl
	GetAllTracksUrl      = "/" + apiPath + songsPrefix
	GetTrackUrlWithoutId = GetAllTracksUrl
	GetTrackUrl          = GetTrackUrlWithoutId + id
	GetPopularTracksUrl  = GetAllTracksUrl + popular
	DeleteTrackUrl       = GetTrackUrl
)

// auth urls
const (
	LoginUrl       = "/" + apiPath + login
	LogoutUrl      = "/" + apiPath + logout
	SignUpUrl      = "/" + apiPath + signUp
	GetUserUrl     = "/" + apiPath + usersPrefix + id
	GetSelfUserUrl = "/" + apiPath + usersPrefix + self
	GetCSRFAuthUrl = "/" + apiPath + getCSRF
	GetStaticUrl   = "/" + apiPath + "/static/"
)
