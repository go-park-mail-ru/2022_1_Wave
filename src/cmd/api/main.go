package main

import (
	"github.com/go-park-mail-ru/2022_1_Wave/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	v1 = "v1"
)

func main() {
	albumRouter := mux.NewRouter()
	api := &api.Album{}

	albumRouter.HandleFunc("/v1/album.create", api.CreateAlbum)

	log.Println("start serving :8000")
	http.ListenAndServe(":8000", albumRouter)
}
