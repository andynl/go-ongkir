package main

import (
	"log"
	"net/http"
	"runtime"

	"github.com/andynl/go-ongkir/routes"
	"github.com/gorilla/handlers"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	port := "4321"

	headersCORS := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Connection", "Accept", "Upgrade-Insecure-Requests"})
	originsCORS := handlers.AllowedOrigins([]string{"*"})
	methodsCORS := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(originsCORS, headersCORS, methodsCORS)(routes.Router)))
}
