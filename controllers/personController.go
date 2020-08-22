package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilhampraset/gingorm-boilerplate/models"
	"log"
)

func GetPeople(c *gin.Context) {
	var people []models.Person

	if err := models.DB.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"result":  people,
		})
	}
}

func GetPerson(c *gin.Context) {
	id := c.Params.ByName("id")

	var person models.Person

	if err := models.DB.Where("id = ?", id).First(&person).Error; err != nil {
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

	models.DB.Create(&person)
	c.JSON(200, gin.H{
		"code":    201,
		"message": "create success",
	})
}

func UpdatePerson(c *gin.Context) {
	var person models.Person
	id := c.Params.ByName("id")
	if err := models.DB.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
	}
	c.BindJSON(&person)
	models.DB.Save(&person)
	c.JSON(200, person)
}

func DeletePerson(c *gin.Context) {
	var person models.Person
	id := c.Params.ByName("id")

	if err := models.DB.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
	}

	models.DB.Delete(&person)
	c.JSON(200, person)
}
