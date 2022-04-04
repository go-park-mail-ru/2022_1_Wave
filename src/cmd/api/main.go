package main

import (
	"github.com/go-park-mail-ru/2022_1_Wave/init/router"
	"github.com/go-park-mail-ru/2022_1_Wave/init/storage"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/interfaces"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/storage/local"
	"log"
	"net/http"
)

// ConfigFilename config
const ConfigFilename = "config.toml"
const port = ":5000"

func main() {
	//if err := config.LoadConfig(ConfigFilename); err != nil {
	//	log.Fatal(err)
	//} else {
	//	log.Println("config loaded successfuly: ", config.C)
	//}

	const quantity = 10
	localStorage := utilsInterfaces.GlobalStorageInterface(structStorageLocal.LocalStorage{})
	storage.InitStorage(quantity, &localStorage)

	rout := router.Router()

	log.Println("start serving :5000")

	err := http.ListenAndServe(port, rout)
	if err != nil {
		log.Fatal(err)
	}
}
