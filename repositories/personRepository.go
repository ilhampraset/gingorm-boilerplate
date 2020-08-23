package repository

import (
	. "github.com/ilhampraset/gingorm-boilerplate/models"
	"github.com/jinzhu/gorm"
)

type PersonRepository struct {
	db *gorm.DB
}

func (repository *PersonRepository) FindAll() []*Person {
	people := []*Person{}

	repository.db.Find(&people)
	return people
}

func NewPersonRepository(database *gorm.DB) *PersonRepository {
	return &PersonRepository{db: database}
}
