package main

import (
	"github.com/google/wire"
	c "github.com/ilhampraset/gingorm-boilerplate/controllers"
	r "github.com/ilhampraset/gingorm-boilerplate/repositories"
	s "github.com/ilhampraset/gingorm-boilerplate/services"
	"github.com/jinzhu/gorm"
)

func initPersonAPI(db *gorm.DB) c.PersonController {
	wire.Build(r.ProvidePersonRepository, s.ProvidePersonService, c.ProvidePersonController)

	return c.PersonController{}
}
