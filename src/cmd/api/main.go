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
	createArtistUrl = "/" + currentApiVersion + artistsSuffix
	updateArtistUrl = createArtistUrl
	getArtistUrl    = "/" + currentApiVersion + artistsSuffix + idSuffix
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

	log.Println("start serving :8000")
	http.ListenAndServe(":8000", router)
}
