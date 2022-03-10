package api

import (
	docs "github.com/go-park-mail-ru/2022_1_Wave/docs"
	"github.com/go-park-mail-ru/2022_1_Wave/middleware"
	"github.com/gorilla/mux"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func InitRouter() (router *mux.Router) {
	router = mux.NewRouter()
	SetAlbumsRoutes(router)
	SetArtistsRoutes(router)
	SetTracksRoutes(router)
	SetAuthRoutes(router)
	SetDocsPath(router)
	SetStaticHandle(router)
	return
}

// albums
func SetAlbumsRoutes(router *mux.Router) {
	router.HandleFunc(GetAllAlbumsUrl, GetAlbums).Methods(http.MethodGet)
	router.HandleFunc(CreateAlbumUrl, CreateAlbum).Methods(http.MethodPost)
	router.HandleFunc(UpdateAlbumUrl, UpdateAlbum).Methods(http.MethodPut)
	router.HandleFunc(GetAlbumUrl, GetAlbum).Methods(http.MethodGet)
	router.HandleFunc(GetPopularAlbumsUrl, GetPopularAlbums).Methods(http.MethodGet)
	router.HandleFunc(DeleteAlbumUrl, DeleteAlbum).Methods(http.MethodDelete)
}

// artists
func SetArtistsRoutes(router *mux.Router) {
	router.HandleFunc(GetAllArtistsUrl, GetArtists).Methods(http.MethodGet)
	router.HandleFunc(CreateArtistUrl, CreateArtist).Methods(http.MethodPost)
	router.HandleFunc(UpdateArtistUrl, UpdateArtist).Methods(http.MethodPut)
	router.HandleFunc(GetArtistUrl, GetArtist).Methods(http.MethodGet)
	router.HandleFunc(GetPopularArtistsUrl, GetPopularArtists).Methods(http.MethodGet)
	router.HandleFunc(DeleteArtistUrl, DeleteArtist).Methods(http.MethodDelete)
}

// songs
func SetTracksRoutes(router *mux.Router) {
	router.HandleFunc(GetAllTracksUrl, GetTracks).Methods(http.MethodGet)
	router.HandleFunc(CreateTrackUrl, CreateTrack).Methods(http.MethodPost)
	router.HandleFunc(UpdateTrackUrl, UpdateTrack).Methods(http.MethodPut)
	router.HandleFunc(GetTrackUrl, GetTrack).Methods(http.MethodGet)
	router.HandleFunc(GetPopularTracksUrl, GetPopularTracks).Methods(http.MethodGet)
	router.HandleFunc(DeleteTrackUrl, DeleteTrack).Methods(http.MethodDelete)
}

// auth
func SetAuthRoutes(router *mux.Router) {
	router.HandleFunc(LoginUrl, middleware.CSRF(middleware.NotAuth(Login))).Methods(http.MethodPost)
	router.HandleFunc(LogoutUrl, middleware.CSRF(middleware.Auth(Logout))).Methods(http.MethodPost)
	router.HandleFunc(SignUpUrl, middleware.CSRF(middleware.NotAuth(SignUp))).Methods(http.MethodPost)
	router.HandleFunc(GetUserUrl, GetUser).Methods(http.MethodGet)
	router.HandleFunc(GetSelfUserUrl, middleware.CSRF(middleware.Auth(GetSelfUser))).Methods(http.MethodGet)
	router.HandleFunc(GetCSRFAuthUrl, GetCSRF).Methods(http.MethodGet)
}

// docs
func SetDocsPath(router *mux.Router) {
	docs.SwaggerInfo.BasePath = "/"
	router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	router.PathPrefix("/docs/*any").Handler(swaggerFiles.Handler)
}

// static
func SetStaticHandle(router *mux.Router) {
	// /net/v1/static/img/album/123.jpg -> ./static/img/album/123.jpg
	staticHandler := http.StripPrefix(
		GetStaticUrl,
		http.FileServer(http.Dir("./static")),
	)
	router.Handle(GetStaticUrl, staticHandler)
}
