package main

import (
	"database/sql"
	"github.com/NNKulickov/wave.music_backend/api"
	"github.com/NNKulickov/wave.music_backend/config"
	docs "github.com/NNKulickov/wave.music_backend/docs"
	"github.com/NNKulickov/wave.music_backend/service"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

const CONFIG_FILENAME = "config.toml"

func main() {

	var err error
	if err = config.LoadConfig(CONFIG_FILENAME); err != nil {
		log.Fatal(err)
	} else {
		log.Println("config loaded successfuly: ", config.C)
	}
	docs.SwaggerInfo.BasePath = "/api"

	router := gin.Default()
	api.DefineRoutes(router)

	if service.DB, err = sql.Open("postgres", config.C.DBConnectionString); err != nil {
		log.Fatal(err)
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	if err = router.Run(":5000"); err != nil {
		log.Fatal(err)
	}

}
