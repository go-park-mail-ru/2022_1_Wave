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
	router.HandleFunc("/v1/login", middleware.CSRF(middleware.NotAuth(api.Login))).Methods(http.MethodPost)
	router.HandleFunc("/v1/logout", middleware.CSRF(middleware.Auth(api.Logout))).Methods(http.MethodGet)
	router.HandleFunc("/v1/signup", middleware.CSRF(middleware.NotAuth(api.SignUp))).Methods(http.MethodPost)
	router.HandleFunc("/v1/users/{id:[0-9]+}", api.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/v1/users/self", middleware.CSRF(middleware.Auth(api.GetSelfUser))).Methods(http.MethodGet)
	router.HandleFunc("/v1/get_csrf", api.GetCSRF).Methods(http.MethodGet)

	http.ListenAndServe(":1234", router)

	/*router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	if err = router.Run(":5000"); err != nil {
		log.Fatal(err)
	}*/
}
