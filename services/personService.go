package service

import (
	"fmt"
	. "github.com/ilhampraset/gingorm-boilerplate/models"
	. "github.com/ilhampraset/gingorm-boilerplate/repositories"
)

type PersonService struct {
	repository *PersonRepository
}

func (service *PersonService) FindAll() []*Person {
	fmt.Println(service.repository.FindAll())
	return service.repository.FindAll()
}

func NewPersonService(repository *PersonRepository) *PersonService {
	return &PersonService{repository: repository}
}
