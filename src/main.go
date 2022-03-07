package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2022_1_Wave/api"
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	docs "github.com/go-park-mail-ru/2022_1_Wave/docs"
	"github.com/go-park-mail-ru/2022_1_Wave/service"
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
