package main

import (
	"github.com/go-park-mail-ru/2022_1_Wave/api"
	docs "github.com/go-park-mail-ru/2022_1_Wave/docs"
	"github.com/go-park-mail-ru/2022_1_Wave/middleware"
	"github.com/gorilla/mux"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

const (
	v1            = "v1/"
	albumsSuffix  = "albums/"
	artistsSuffix = "artists/"
	idSuffix      = "{id:[0-9]+}"
)

//config
const (
	apiSuffix         = "api/"
	currentApiVersion = v1
)

// albums urls
const (
	createAlbumUrl = "/" + apiSuffix + currentApiVersion + albumsSuffix
	updateAlbumUrl = createAlbumUrl
	getAlbumUrl    = "/" + apiSuffix + currentApiVersion + albumsSuffix + idSuffix
	deleteAlbumUrl = getAlbumUrl
)

// artists urls
const (
	createArtistUrl = "/" + apiSuffix + currentApiVersion + artistsSuffix
	updateArtistUrl = createArtistUrl
	getArtistUrl    = "/" + apiSuffix + currentApiVersion + artistsSuffix + idSuffix
	deleteArtistUrl = getArtistUrl
)

func main() {
	router := mux.NewRouter()

	// albums
	router.HandleFunc(createAlbumUrl, api.CreateAlbum)
	router.HandleFunc(updateAlbumUrl, api.UpdateAlbum)
	router.HandleFunc(getAlbumUrl, api.GetAlbum)
	router.HandleFunc(deleteAlbumUrl, api.DeleteAlbum)

	// artists
	router.HandleFunc(createArtistUrl, api.CreateArtist)
	router.HandleFunc(updateArtistUrl, api.UpdateArtist)
	router.HandleFunc(getArtistUrl, api.GetArtist)
	router.HandleFunc(deleteArtistUrl, api.DeleteArtist)

	//auth
	router.HandleFunc("/api/v1/login", middleware.CSRF(middleware.NotAuth(api.Login))).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/logout", middleware.CSRF(middleware.Auth(api.Logout))).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/signup", middleware.CSRF(middleware.NotAuth(api.SignUp))).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/users/{id:[0-9]+}", api.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/users/self", middleware.CSRF(middleware.Auth(api.GetSelfUser))).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/get_csrf", api.GetCSRF).Methods(http.MethodGet)

	docs.SwaggerInfo.BasePath = "/"
	router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	router.PathPrefix("/docs/*any").Handler(swaggerFiles.Handler)

	log.Println("start serving :5000")
	http.ListenAndServe(":5000", router)
}
