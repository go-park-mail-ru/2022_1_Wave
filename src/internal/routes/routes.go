package routes

import (
	"github.com/go-park-mail-ru/2022_1_Wave/api"
	docs "github.com/go-park-mail-ru/2022_1_Wave/docs"
	"github.com/go-park-mail-ru/2022_1_Wave/middleware"
	"github.com/gorilla/mux"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

// albums
func SetAlbumsRoutes(router *mux.Router) {
	router.HandleFunc(api.GetAllAlbumsUrl, api.GetAlbums).Methods(http.MethodGet)
	router.HandleFunc(api.CreateAlbumUrl, api.CreateAlbum).Methods(http.MethodPost)
	router.HandleFunc(api.UpdateAlbumUrl, api.UpdateAlbum).Methods(http.MethodPut)
	router.HandleFunc(api.GetAlbumUrl, api.GetAlbum).Methods(http.MethodGet)
	router.HandleFunc(api.GetPopularAlbumsUrl, api.GetPopularAlbums).Methods(http.MethodGet)
	router.HandleFunc(api.DeleteAlbumUrl, api.DeleteAlbum).Methods(http.MethodDelete)
}

// artists
func SetArtistsRoutes(router *mux.Router) {
	router.HandleFunc(api.GetAllArtistsUrl, api.GetArtists).Methods(http.MethodGet)
	router.HandleFunc(api.CreateArtistUrl, api.CreateArtist).Methods(http.MethodPost)
	router.HandleFunc(api.UpdateArtistUrl, api.UpdateArtist).Methods(http.MethodPut)
	router.HandleFunc(api.GetArtistUrl, api.GetArtist).Methods(http.MethodGet)
	router.HandleFunc(api.GetPopularArtistsUrl, api.GetPopularArtists).Methods(http.MethodGet)
	router.HandleFunc(api.DeleteArtistUrl, api.DeleteArtist).Methods(http.MethodDelete)
}

// songs
func SetSongsRoutes(router *mux.Router) {
	router.HandleFunc(api.GetAllTracksUrl, api.GetTracks).Methods(http.MethodGet)
	router.HandleFunc(api.CreateTrackUrl, api.CreateTrack).Methods(http.MethodPost)
	router.HandleFunc(api.UpdateTrackUrl, api.UpdateTrack).Methods(http.MethodPut)
	router.HandleFunc(api.GetTrackUrl, api.GetTrack).Methods(http.MethodGet)
	router.HandleFunc(api.GetPopularTracksUrl, api.GetPopularTracks).Methods(http.MethodGet)
	router.HandleFunc(api.DeleteTrackUrl, api.DeleteTrack).Methods(http.MethodDelete)
}

// auth
func SetAuthRoutes(router *mux.Router) {
	router.HandleFunc(api.LoginUrl, middleware.CSRF(middleware.NotAuth(api.Login))).Methods(http.MethodPost)
	router.HandleFunc(api.LogoutUrl, middleware.CSRF(middleware.Auth(api.Logout))).Methods(http.MethodPost)
	router.HandleFunc(api.SignUpUrl, middleware.CSRF(middleware.NotAuth(api.SignUp))).Methods(http.MethodPost)
	router.HandleFunc(api.GetUserUrl, api.GetUser).Methods(http.MethodGet)
	router.HandleFunc(api.GetSelfUserUrl, middleware.CSRF(middleware.Auth(api.GetSelfUser))).Methods(http.MethodGet)
	router.HandleFunc(api.GetCSRFAuthUrl, api.GetCSRF).Methods(http.MethodGet)
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
		api.GetStaticUrl,
		http.FileServer(http.Dir("./static")),
	)
	router.Handle(api.GetStaticUrl, staticHandler)
}
