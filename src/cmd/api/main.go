package main

import (
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	_ "github.com/go-park-mail-ru/2022_1_Wave/docs"
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/init/router"
	"github.com/go-park-mail-ru/2022_1_Wave/init/storage"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	albumUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/usecase"
	albumCoverUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/albumCover/usecase"
	artistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/artist/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	trackUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/track/usecase"
	"github.com/labstack/echo/v4"
)

// ConfigFilename config
const ConfigFilename = "config.toml"
const port = ":5000"

const dbSize = 10

func main() {
	// database model, you can switch it
	dbType := internal.Postgres
	//database := internal.Local

	e := echo.New()

	//host, _ := os.Hostname()
	logs, err := logger.InitLogrus(port, dbType)
	if err != nil {
		e.Logger.Fatalf("error to init logrus:", err)
	}

	e.Use(logs.ColoredLogMiddleware)
	e.Use(logs.JsonLogMiddleware)
	e.Logger.SetOutput(logs.Logrus.Writer())

	if err := config.LoadConfig(ConfigFilename); err != nil {
		logs.Logrus.Fatal("error to load config:", err)
	}
	logs.Logrus.Info("config loaded successful")

	if err := storage.InitStorage(dbSize, dbType); err != nil {
		logs.Logrus.Fatal("error to init storage type", dbType, "err:", err)
	}

	logs.Logrus.Info("Success init storage type", dbType)

	artistRepoLen, err := artistUseCase.UseCase.GetSize(domain.ArtistMutex)
	if err != nil {
		logs.Logrus.Fatal("Error:", err)
	}

	albumRepoLen, err := albumUseCase.UseCase.GetSize(domain.AlbumMutex)
	if err != nil {
		logs.Logrus.Fatal("Error:", err)
	}

	albumRepoCoverLen, err := albumCoverUseCase.UseCase.GetSize(domain.AlbumCoverMutex)
	if err != nil {
		logs.Logrus.Fatal("Error:", err)
	}

	trackRepoLen, err := trackUseCase.UseCase.GetSize(domain.TrackMutex)
	if err != nil {
		logs.Logrus.Fatal("Error:", err)
	}

	logs.Logrus.Info("Artists:", artistRepoLen)
	logs.Logrus.Info("Albums:", albumRepoLen)
	logs.Logrus.Info("AlbumCovers:", albumRepoCoverLen)
	logs.Logrus.Info("Tracks:", trackRepoLen)

	router.Router(e)

	logs.Logrus.Warn("start listening on", port)

	if err := e.Start("0.0.0.0:5000"); err != nil {
		logs.Logrus.Fatal("server error:", err)
	}

	logs.Logrus.Warn("shutdown")
}
