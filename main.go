package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Serares/file-converter/config"
	"github.com/Serares/file-converter/handlers"
	"github.com/Serares/file-converter/routers"
	"github.com/gorilla/mux"
)

func main() {
	config := config.LoadConfig()

	router := mux.NewRouter()

	uploadHandler := handlers.NewUploadHandler()

	routers.RegisterUploadRoutes(router, uploadHandler)

	fmt.Println("Server started on port", config.Port)
	log.Println("http server started on port", config.Port)
	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(config.Port), router))
}
