package main

import (
	_ "github.com/go-park-mail-ru/2022_1_Wave/docs"
	"github.com/go-park-mail-ru/2022_1_Wave/init/router"
	"github.com/go-park-mail-ru/2022_1_Wave/init/storage"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/interfaces"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/storage/local"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"reflect"
)

// ConfigFilename config
const ConfigFilename = "config.toml"
const port = ":5000"

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	//if err := config.LoadConfig(ConfigFilename); err != nil {
	//	log.Fatal(err)
	//} else {
	//	log.Println("config loaded successfuly: ", config.C)
	//}

	const quantity = 10
	localStorage := utilsInterfaces.GlobalStorageInterface(structStorageLocal.LocalStorage{})

	if err := storage.InitStorage(quantity, &localStorage); err != nil {
		e.Logger.Fatal("error to init storage type <%v>, err: %v\n", reflect.TypeOf(localStorage), err)
	}

	e.Logger.Printf("Success init local storage type <%v>\n", reflect.TypeOf(localStorage))
	e.Logger.Printf("Artists: %v", localStorage.GetArtistRepoLen())
	e.Logger.Printf("Albums: %v", localStorage.GetAlbumRepoLen())
	e.Logger.Printf("Tracks: %v", localStorage.GetArtistRepoLen())

	router.Router(e)

	e.Logger.Warnf("start listening on %s", port)
	err := e.Start("0.0.0.0:5000")
	if err != nil {
		e.Logger.Errorf("server error: %s", err)
	}

	e.Logger.Warnf("shutdown")

}
