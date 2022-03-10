package main

import (
	"github.com/go-park-mail-ru/2022_1_Wave/api"
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	"github.com/go-park-mail-ru/2022_1_Wave/db"
	"log"
	"net/http"
)

//config
const CONFIG_FILENAME = "config.toml"
const PATH_TO_STATIC = "./static"
const port = ":5000"

func main() {
	if err := config.LoadConfig(CONFIG_FILENAME); err != nil {
		log.Fatal(err)
	} else {
		log.Println("config loaded successfuly: ", config.C)
	}

	const quantity = 10
	db.Storage.InitStorage(quantity)

	router := api.InitRouter()

	log.Println("start serving :5000")

	http.ListenAndServe(port, router)
}
