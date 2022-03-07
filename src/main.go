package main

import (
	"github.com/NNKulickov/wave.music_backend/api"
	"github.com/NNKulickov/wave.music_backend/config"
	docs "github.com/NNKulickov/wave.music_backend/docs"
	"github.com/NNKulickov/wave.music_backend/middleware"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
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

	router := mux.NewRouter()
	router.HandleFunc("/login/", api.Login).Methods(http.MethodPost)
	router.HandleFunc("/logout/", middleware.Session(api.Logout)).Methods(http.MethodGet)
	router.HandleFunc("/signup/", api.SignUp).Methods(http.MethodPost)

	http.ListenAndServe(":80", router)

	/*router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	if err = router.Run(":5000"); err != nil {
		log.Fatal(err)
	}*/
}
