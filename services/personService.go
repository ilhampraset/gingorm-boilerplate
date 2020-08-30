package service

import (
	"fmt"
	m "github.com/ilhampraset/gingorm-boilerplate/models"
	. "github.com/ilhampraset/gingorm-boilerplate/repositories"
	"reflect"
)

type PersonService struct {
	repository *PersonRepository
}

func ProvidePersonService(p *PersonRepository) *PersonService {
	return &PersonService{repository: p}
}

func (s *PersonService) FindAll() ([]*m.Person, error) {

	people, err := s.repository.FindAll()
	if err != nil {
		error.Error(err)
	} else {
		var c m.Person
		field := reflect.TypeOf(c)
		for _, val := range people {
			fmt.Println(field.Field(0).Name, ":", people[0].ID)
			fmt.Println(field.Field(1).Name, ":", val.Firstname)
			fmt.Println(field.Field(2).Name, ":", val.Lastname)
			if val.Lastname == "praset" {
				val.Firstname = "celeng"
			}
			if val.ID == 2 {
				val.Lastname = "yolooo"
			}

		}

	}

	return people, nil
}

func (s *PersonService) FindByID(id uint) m.Person {
	person, _ := s.repository.FindByID(id)
	return person
}

func (s *PersonService) Save(person m.Person) m.Person {
	person, _ = s.repository.Save(person)
	return person
}

func (s *PersonService) Delete(person m.Person) {
	s.repository.Delete(person)
}
