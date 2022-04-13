package main

import (
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	_ "github.com/go-park-mail-ru/2022_1_Wave/docs"
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/init/system"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/labstack/echo/v4"
)

// ConfigFilename config
const ConfigFilename = "config.toml"

// todo вынести в конфиг
const port = ":5000"

const dbSize = 0

func main() {
	// database model, you can switch it
	dbType := internal.Postgres
	//dbType := internal.Local

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

	//env, err := storage.InitStorage(dbSize, dbType)
	//
	//if err != nil {
	//	logs.Logrus.Fatal("error to init storage type", dbType, "err:", err)
	//}

	logs.Logrus.Info("Success init storage type", dbType)

	//artistRepoLen, err := artistUseCase.UseCase.GetSize()
	//if err != nil {
	//	logs.Logrus.Fatal("Error:", err)
	//}
	//
	//albumRepoLen, err := albumUseCase.UseCase.GetSize()
	//if err != nil {
	//	logs.Logrus.Fatal("Error:", err)
	//}
	//
	//albumRepoCoverLen, err := albumCoverUseCase.UseCase.GetSize()
	//if err != nil {
	//	logs.Logrus.Fatal("Error:", err)
	//}
	//
	//trackRepoLen, err := trackUseCase.UseCase.GetSize()
	//if err != nil {
	//	logs.Logrus.Fatal("Error:", err)
	//}

	//logs.Logrus.Info("Artists:", artistRepoLen)
	//logs.Logrus.Info("Albums:", albumRepoLen)
	//logs.Logrus.Info("AlbumCovers:", albumRepoCoverLen)
	//logs.Logrus.Info("Tracks:", trackRepoLen)
	//

	if err := system.Init(e, dbSize, dbType); err != nil {
		logs.Logrus.Fatal("Error:", err)
	}

	logs.Logrus.Warn("start listening on", port)

	if err := e.Start("0.0.0.0:5000"); err != nil {
		logs.Logrus.Fatal("server error:", err)
	}

	logs.Logrus.Warn("shutdown")
}
