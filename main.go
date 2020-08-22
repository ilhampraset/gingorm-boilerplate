package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ilhampraset/gingorm-boilerplate/controllers"
	"github.com/ilhampraset/gingorm-boilerplate/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	models.InitDb()
}

func main() {

	app := gin.Default()
	app.GET("/people", controllers.GetPeople)
	app.GET("/people/:id", controllers.GetPerson)
	app.POST("/people", controllers.CreatePerson)
	app.PUT("/people/:id", controllers.UpdatePerson)
	app.DELETE("/people/:id", controllers.DeletePerson)
	app.Run()
}
