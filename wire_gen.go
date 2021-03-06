// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/ilhampraset/gingorm-boilerplate/controllers"
	"github.com/ilhampraset/gingorm-boilerplate/repositories"
	"github.com/ilhampraset/gingorm-boilerplate/services"
	"github.com/jinzhu/gorm"
)

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Injectors from wire.go:

func InitPersonAPI(db *gorm.DB) controllers.PersonController {
	personRepository := repository.ProvidePersonRepository(db)
	personService := service.ProvidePersonService(personRepository)
	personController := controllers.ProvidePersonController(personService)
	return personController
}
