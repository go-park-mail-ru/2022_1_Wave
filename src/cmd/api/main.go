package main

import (
	_ "github.com/go-park-mail-ru/2022_1_Wave/docs"
	"github.com/go-park-mail-ru/2022_1_Wave/init/router"
	"github.com/go-park-mail-ru/2022_1_Wave/init/storage"
	utilsInterfaces "github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/interfaces"
	structStoragePostgresql "github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/storage/postgresql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"reflect"
)

// ConfigFilename config
const ConfigFilename = "config.toml"
const port = ":5000"

const dbSize = 10

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	//if err := config.LoadConfig(ConfigFilename); err != nil {
	//	log.Fatal(err)
	//} else {
	//	log.Println("config loaded successfuly: ", config.C)
	//}

	globalStorage := utilsInterfaces.GlobalStorageInterface(structStoragePostgresql.Postgres{})
	//globalStorage := utilsInterfaces.GlobalStorageInterface(structStorageLocal.LocalStorage{})

	if err := storage.InitStorage(dbSize, &globalStorage); err != nil {
		e.Logger.Fatal("error to init storage type <%v>, err: %v", reflect.TypeOf(globalStorage), err)
	}

	e.Logger.Printf("Success init storage type <%v>", reflect.TypeOf(globalStorage))

	artistRepoLen, err := globalStorage.GetArtistRepoLen()
	if err != nil {
		e.Logger.Fatalf("Error: %v", err)
	}

	albumRepoLen, err := globalStorage.GetAlbumRepoLen()
	if err != nil {
		e.Logger.Fatalf("Error: %v", err)
	}

	albumRepoCoverLen, err := globalStorage.GetAlbumCoverRepoLen()
	if err != nil {
		e.Logger.Fatalf("Error: %v", err)
	}

	trackRepoLen, err := globalStorage.GetTrackRepoLen()
	if err != nil {
		e.Logger.Fatalf("Error: %v", err)
	}

	e.Logger.Printf("Artists: %v", artistRepoLen)
	e.Logger.Printf("Albums: %v", albumRepoLen)
	e.Logger.Printf("AlbumCovers: %v", albumRepoCoverLen)
	e.Logger.Printf("Tracks: %v", trackRepoLen)

	router.Router(e)

	e.Logger.Warnf("start listening on %s", port)

	if err := e.Start("0.0.0.0:5000"); err != nil {
		e.Logger.Errorf("server error: %s", err)
	}

	e.Logger.Warnf("shutdown")

}
