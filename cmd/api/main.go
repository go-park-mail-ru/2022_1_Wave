package main

import (
	echoprometheus "github.com/globocom/echo-prometheus"
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	_ "github.com/go-park-mail-ru/2022_1_Wave/docs"
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/init/system"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// ConfigFilename config
const ConfigFilename = "config.toml"

// todo вынести в конфиг
const port = ":5000"

const randomGeneratedDbSize = 0

func skipUrls(metricMiddleware echo.MiddlewareFunc) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Path() != "/metrics" {
				newNext := metricMiddleware(next)
				return newNext(c)
			} else {
				return next(c)
			}
		}
	}
}

func main() {
	dbType := internal.Postgres
	e := echo.New()
	//p := prometheus.NewPrometheus("echo", nil)
	//p.Use(e)
	e.Use(skipUrls(echoprometheus.MetricsMiddleware()))

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
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

	if err := system.Init(e, randomGeneratedDbSize, dbType); err != nil {
		logs.Logrus.Fatal("Error:", err)
	}

	logs.Logrus.Info("Success init storage type", dbType)
	logs.Logrus.Warn("start listening on", port)

	if err := e.Start("0.0.0.0:5000"); err != nil {
		logs.Logrus.Fatal("server error:", err)
	}

	logs.Logrus.Warn("shutdown")
}
