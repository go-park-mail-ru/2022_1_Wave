package main

import (
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	"github.com/go-park-mail-ru/2022_1_Wave/db"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/net"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//config
const CONFIG_FILENAME = "config.toml"
const PATH_TO_STATIC = "./static"

func main() {
	if err := config.LoadConfig(CONFIG_FILENAME); err != nil {
		log.Fatal(err)
	} else {
		log.Println("config loaded successfuly: ", config.C)
	}

	db.Storage.InitStorage()

	router := mux.NewRouter()

	net.SetAlbumsRoutes(router)
	net.SetArtistsRoutes(router)
	net.SetSongsRoutes(router)
	net.SetAuthRoutes(router)
	net.SetDocsPath(router)
	net.SetStaticHandle(router)

	log.Println("start serving :5000")
	http.ListenAndServe(":5000", router)
}
