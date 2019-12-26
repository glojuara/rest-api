package main

import (
	"log"
	"net/http"

	"github.com/glojuara/rest-api/adapter/controller"
	infra "github.com/glojuara/rest-api/infrastructure"
	"github.com/gorilla/mux"
)

// função principal
func main() {
	router := mux.NewRouter()
	controller.GetInstance(router, infra.MockIPersonRepository{})
	log.Fatal(http.ListenAndServe(":8000", router))
}
