package main

import (
	"github.com/go-park-mail-ru/2022_1_Wave/api"
	"github.com/go-park-mail-ru/2022_1_Wave/config"
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

// auth urls
const (
	loginUrl       = "/" + apiSuffix + currentApiVersion + "login"
	logoutUrl      = "/" + apiSuffix + currentApiVersion + "logout"
	signUpUrl      = "/" + apiSuffix + currentApiVersion + "signup"
	getUserUrl     = "/" + apiSuffix + currentApiVersion + "users/" + idSuffix
	getSelfUserUrl = "/" + apiSuffix + currentApiVersion + "users/" + "self"
	getCSRFAuthUrl = "/" + apiSuffix + currentApiVersion + "get_csrf"
)

const CONFIG_FILENAME = "config.toml"

func main() {
	var err error
	if err = config.LoadConfig(CONFIG_FILENAME); err != nil {
		log.Fatal(err)
	} else {
		log.Println("config loaded successfuly: ", config.C)
	}

	router := mux.NewRouter()

	// albums
	router.HandleFunc(createAlbumUrl, api.CreateAlbum).Methods(http.MethodPost)
	router.HandleFunc(updateAlbumUrl, api.UpdateAlbum).Methods(http.MethodPut)
	router.HandleFunc(getAlbumUrl, api.GetAlbum).Methods(http.MethodGet)
	router.HandleFunc(deleteAlbumUrl, api.DeleteAlbum).Methods(http.MethodDelete)

	// artists
	router.HandleFunc(createArtistUrl, api.CreateArtist).Methods(http.MethodPost)
	router.HandleFunc(updateArtistUrl, api.UpdateArtist).Methods(http.MethodPut)
	router.HandleFunc(getArtistUrl, api.GetArtist).Methods(http.MethodGet)
	router.HandleFunc(deleteArtistUrl, api.DeleteArtist).Methods(http.MethodDelete)

	//auth
	router.HandleFunc(loginUrl, middleware.CSRF(middleware.NotAuth(api.Login))).Methods(http.MethodPost)
	router.HandleFunc(logoutUrl, middleware.CSRF(middleware.Auth(api.Logout))).Methods(http.MethodPost)
	router.HandleFunc(signUpUrl, middleware.CSRF(middleware.NotAuth(api.SignUp))).Methods(http.MethodPost)
	router.HandleFunc(getUserUrl, api.GetUser).Methods(http.MethodGet)
	router.HandleFunc(getSelfUserUrl, middleware.CSRF(middleware.Auth(api.GetSelfUser))).Methods(http.MethodGet)
	router.HandleFunc(getCSRFAuthUrl, api.GetCSRF).Methods(http.MethodGet)

	docs.SwaggerInfo.BasePath = "/"
	router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	router.PathPrefix("/docs/*any").Handler(swaggerFiles.Handler)

	log.Println("start serving :5000")
	http.ListenAndServe(":5000", router)
}
