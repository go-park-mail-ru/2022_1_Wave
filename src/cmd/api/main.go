package main

import (
	"github.com/go-park-mail-ru/2022_1_Wave/api"
	"github.com/gorilla/mux"
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
	createAlbumUrl = "/" + currentApiVersion + albumsSuffix
	updateAlbumUrl = createAlbumUrl
	getAlbumUrl    = "/" + currentApiVersion + albumsSuffix + idSuffix
	deleteAlbumUrl = getAlbumUrl
)

// artists urls
const (
	createArtistUrl  = "/" + apiSuffix + currentApiVersion + artistsSuffix
	updateArtistUrl  = createArtistUrl
	getAllArtistsUrl = "/" + apiSuffix + currentApiVersion + artistsSuffix
	getArtistUrl     = "/" + apiSuffix + currentApiVersion + artistsSuffix + idSuffix
	deleteArtistUrl  = getArtistUrl
)

func main() {
	router := mux.NewRouter()

	// albums
	router.HandleFunc(createAlbumUrl, api.CreateAlbum)
	router.HandleFunc(updateAlbumUrl, api.UpdateAlbum)
	router.HandleFunc(getAlbumUrl, api.GetAlbum)
	router.HandleFunc(deleteAlbumUrl, api.DeleteAlbum)

	// artists
	router.HandleFunc(getAllArtistsUrl, api.GetArtists).Methods(http.MethodGet)
	router.HandleFunc(createArtistUrl, api.CreateArtist).Methods(http.MethodPost)
	router.HandleFunc(updateArtistUrl, api.UpdateArtist).Methods(http.MethodPut)
	router.HandleFunc(getArtistUrl, api.GetArtist).Methods(http.MethodGet)
	router.HandleFunc(deleteArtistUrl, api.DeleteArtist).Methods(http.MethodDelete)

	log.Println("start serving :8000")
	http.ListenAndServe(":8000", router)
}
