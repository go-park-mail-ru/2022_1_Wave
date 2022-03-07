package main

import (
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	docs "github.com/go-park-mail-ru/2022_1_Wave/docs"
	_ "github.com/lib/pq"
	"log"
)

const CONFIG_FILENAME = "../config/config.toml"

func main() {
	var err error
	if err = config.LoadConfig(CONFIG_FILENAME); err != nil {
		log.Fatal(err)
	} else {
		log.Println("config loaded successfuly: ", config.C)
	}
	docs.SwaggerInfo.BasePath = "/api"

	/*
		if service.DB, err = sql.Open("postgres", config.C.DBConnectionString); err != nil {
			log.Fatal(err)
		}
	*/

	/*router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	if err = router.Run(":5000"); err != nil {
		log.Fatal(err)
	}*/
}
