package net

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
	router.HandleFunc(getAllAlbumsUrl, api.GetAlbums).Methods(http.MethodGet)
	router.HandleFunc(createAlbumUrl, api.CreateAlbum).Methods(http.MethodPost)
	router.HandleFunc(updateAlbumUrl, api.UpdateAlbum).Methods(http.MethodPut)
	router.HandleFunc(getAlbumUrl, api.GetAlbum).Methods(http.MethodGet)
	router.HandleFunc(getPopularAlbumsUrl, api.GetPopularAlbums).Methods(http.MethodGet)
	router.HandleFunc(deleteAlbumUrl, api.DeleteAlbum).Methods(http.MethodDelete)
}

// artists
func SetArtistsRoutes(router *mux.Router) {
	router.HandleFunc(getAllArtistsUrl, api.GetArtists).Methods(http.MethodGet)
	router.HandleFunc(createArtistUrl, api.CreateArtist).Methods(http.MethodPost)
	router.HandleFunc(updateArtistUrl, api.UpdateArtist).Methods(http.MethodPut)
	router.HandleFunc(getArtistUrl, api.GetArtist).Methods(http.MethodGet)
	router.HandleFunc(getPopularArtistsUrl, api.GetPopularArtists).Methods(http.MethodGet)
	router.HandleFunc(deleteArtistUrl, api.DeleteArtist).Methods(http.MethodDelete)
}

// songs
func SetSongsRoutes(router *mux.Router) {
	router.HandleFunc(getAllSongsUrl, api.GetSongs).Methods(http.MethodGet)
	router.HandleFunc(createSongUrl, api.CreateSong).Methods(http.MethodPost)
	router.HandleFunc(updateSongUrl, api.UpdateSong).Methods(http.MethodPut)
	router.HandleFunc(getSongUrl, api.GetSong).Methods(http.MethodGet)
	router.HandleFunc(getPopularSongsUrl, api.GetPopularSongs).Methods(http.MethodGet)
	router.HandleFunc(deleteSongUrl, api.DeleteSong).Methods(http.MethodDelete)
}

// auth
func SetAuthRoutes(router *mux.Router) {
	router.HandleFunc(loginUrl, middleware.CSRF(middleware.NotAuth(api.Login))).Methods(http.MethodPost)
	router.HandleFunc(logoutUrl, middleware.CSRF(middleware.Auth(api.Logout))).Methods(http.MethodPost)
	router.HandleFunc(signUpUrl, middleware.CSRF(middleware.NotAuth(api.SignUp))).Methods(http.MethodPost)
	router.HandleFunc(getUserUrl, api.GetUser).Methods(http.MethodGet)
	router.HandleFunc(getSelfUserUrl, middleware.CSRF(middleware.Auth(api.GetSelfUser))).Methods(http.MethodGet)
	router.HandleFunc(getCSRFAuthUrl, api.GetCSRF).Methods(http.MethodGet)
}

// docs
func SetDocsPath(router *mux.Router) {
	docs.SwaggerInfo.BasePath = "/"
	router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	router.PathPrefix("/docs/*any").Handler(swaggerFiles.Handler)
}

// static
func SetStaticHandle(router *mux.Router) {
	// /api/v1/static/img/album/123.jpg -> ./static/img/album/123.jpg
	staticHandler := http.StripPrefix(
		getStaticUrl,
		http.FileServer(http.Dir("./static")),
	)
	router.Handle(getStaticUrl, staticHandler)
}
