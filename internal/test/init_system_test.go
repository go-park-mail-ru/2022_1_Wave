package test

import (
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/init/system"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInit(t *testing.T) {
	dbType := internal.Postgres
	dbSize := int64(10)
	e := echo.New()

	logs, err := logger.InitLogrus(":5000", dbType)
	require.NoError(t, err)

	e.Use(logs.ColoredLogMiddleware)
	e.Use(logs.JsonLogMiddleware)
	e.Logger.SetOutput(logs.Logrus.Writer())
	logs.Logrus.Debug("hello from init test")
	require.NoError(t, system.Init(e, dbSize, dbType))
}
