package repository

import (
	m "github.com/ilhampraset/gingorm-boilerplate/models"
	"github.com/jinzhu/gorm"
)

type PersonRepository struct {
	Db *gorm.DB
}

func ProvidePersonRepository(db *gorm.DB) *PersonRepository {
	return &PersonRepository{Db: db}
}

func (p *PersonRepository) FindAll() ([]*m.Person, error) {
	var people []*m.Person

	p.Db.Find(&people)
	return people, nil
}

func (p *PersonRepository) FindByID(id uint) (m.Person, error) {
	var person m.Person
	p.Db.First(&person, id)
	return person, nil
}

func (p *PersonRepository) Save(person m.Person) (m.Person, error) {
	p.Db.Save(&person)
	return person, nil
}

func (p *PersonRepository) Delete(person m.Person) {
	p.Db.Delete(person)
}
