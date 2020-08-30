package controllers

import (
	"github.com/gin-gonic/gin"
	m "github.com/ilhampraset/gingorm-boilerplate/models"
	. "github.com/ilhampraset/gingorm-boilerplate/services"
	"log"
	"net/http"
	"strconv"
)

type PersonController struct {
	service *PersonService
}

func ProvidePersonController(p *PersonService) PersonController {

	return PersonController{service: p}
}
func (s *PersonController) GetPeople(c *gin.Context) {

	people, err := s.service.FindAll()
	if err != nil {
		c.JSON(404, gin.H{
			"code":    404,
			"message": "error",
			"error":   err,
		})

	} else {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"result":  people,
		})

	}

}

func (s *PersonController) GetPerson(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	person := s.service.FindByID(uint(id))

	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"result":  person,
	})
}

func (p *PersonController) CreatePerson(c *gin.Context) {
	var person m.Person
	err := c.ShouldBindJSON(&person)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	} else {
		p.service.Save(person)
		c.JSON(200, gin.H{
			"code":    201,
			"message": "create success",
		})
	}

}

func (p *PersonController) UpdatePerson(c *gin.Context) {
	var mp m.Person
	err := c.ShouldBindJSON(&mp)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Params.ByName("id"))
	person := p.service.FindByID(uint(id))
	if person == (m.Person{}) {
		c.Status(http.StatusBadRequest)
		return
	}
	person.Firstname = mp.Firstname
	person.Lastname = mp.Lastname
	p.service.Save(person)
	c.JSON(200, gin.H{
		"code":    201,
		"message": "update success",
	})
}

func (p *PersonController) DeletePerson(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	person := p.service.FindByID(uint(id))
	if person == (m.Person{}) {
		c.Status(http.StatusBadRequest)
		return
	}
	p.service.Delete(person)
	c.JSON(200, person)
}
