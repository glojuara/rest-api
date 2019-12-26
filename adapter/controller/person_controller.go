package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	_ "github.com/glojuara/rest-api/domain"
	infra "github.com/glojuara/rest-api/infrastructure"
	"github.com/gorilla/mux"
)

type personController struct {
	router     *mux.Router
	repository infra.IPersonRepository
}

var instance *personController
var once sync.Once

func GetInstance(router *mux.Router, repository infra.IPersonRepository) *personController {
	once.Do(func() {
		instance = &personController{router: router, repository: repository}
		instance.setup()
	})
	return instance
}

func (c *personController) setup() {
	c.router.HandleFunc("/contato", c.getPeople).Methods("GET")
	c.router.HandleFunc("/contato/{id}", c.getPerson).Methods("GET")
	c.router.HandleFunc("/contato/{id}", c.createPerson).Methods("POST")
	c.router.HandleFunc("/contato/{id}", c.deletePerson).Methods("DELETE")

	people := c.repository.Find()
	fmt.Println(people)

}

func (c *personController) getPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(c.repository.Find())
}
func (c *personController) getPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(c.repository.FindById(params["id"]))
}
func (c *personController) createPerson(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("TODO: implments func ")
}
func (c *personController) deletePerson(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("TODO: implments func ")
}
