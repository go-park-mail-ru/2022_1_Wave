package main

import (
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	"github.com/go-park-mail-ru/2022_1_Wave/db"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/routes"
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

	const quantity = 10
	db.Storage.InitStorage(quantity)

	router := mux.NewRouter()

	routes.SetAlbumsRoutes(router)
	routes.SetArtistsRoutes(router)
	routes.SetSongsRoutes(router)
	routes.SetAuthRoutes(router)
	routes.SetDocsPath(router)
	routes.SetStaticHandle(router)

	log.Println("start serving :5000")
	http.ListenAndServe(":5000", router)
}
