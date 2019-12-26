package infrastructure

import (
	"github.com/glojuara/rest-api/domain"
	_ "github.com/glojuara/rest-api/domain"
)

type IPersonRepository interface {
	// Create(person domain.Person) domain.Person
	Find() []domain.Person
	FindById(id string) domain.Person
	// Delete(id string)
}

type MockIPersonRepository struct {
}

func mockPeople() []domain.Person {
	alessandra := domain.Person{Firstname: "Alessandra"}
	guilherme := domain.Person{Firstname: "Guilherme"}
	people := make([]domain.Person, 0)
	people = append(people, alessandra)
	people = append(people, guilherme)
	return people
}

func (r MockIPersonRepository) Find() []domain.Person {
	people := mockPeople()
	return people
}

func (r MockIPersonRepository) FindById(id string) domain.Person {
	people := mockPeople()
	for _, p := range people {
		if p.ID == id {
			return p
		}

	}

	return domain.Person{}
}
