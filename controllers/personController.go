package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilhampraset/gingorm-boilerplate/models"
	. "github.com/ilhampraset/gingorm-boilerplate/providers"
	. "github.com/ilhampraset/gingorm-boilerplate/services"
	"log"
)

// type Server struct {
// 	personService
// }

func GetPeople(personService *PersonService) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		people := personService.FindAll()
		c.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"result":  people,
		})
	}
	return gin.HandlerFunc(fn)

}

func GetPerson(c *gin.Context) {
	id := c.Params.ByName("id")

	var person models.Person

	if err := DB.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"result":  person,
		})
	}
}

func CreatePerson(c *gin.Context) {
	var person models.Person
	c.BindJSON(&person)

	DB.Create(&person)
	c.JSON(200, gin.H{
		"code":    201,
		"message": "create success",
	})
}

func UpdatePerson(c *gin.Context) {
	var person models.Person
	id := c.Params.ByName("id")
	if err := DB.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
	}
	c.BindJSON(&person)
	DB.Save(&person)
	c.JSON(200, person)
}

func DeletePerson(c *gin.Context) {
	var person models.Person
	id := c.Params.ByName("id")

	if err := DB.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
	}

	DB.Delete(&person)
	c.JSON(200, person)
}
