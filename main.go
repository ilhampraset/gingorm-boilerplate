package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ilhampraset/gingorm-boilerplate/config"
	"github.com/ilhampraset/gingorm-boilerplate/controllers"
	provider "github.com/ilhampraset/gingorm-boilerplate/providers"
	. "github.com/ilhampraset/gingorm-boilerplate/repositories"
	. "github.com/ilhampraset/gingorm-boilerplate/services"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"os"
)

type Args struct {
	ConfigPath string
}

var cfg config.Config

func init() {

	args := ProcessArgs(&cfg)

	if err := cleanenv.ReadConfig(args.ConfigPath, &cfg); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	provider.InitDb(&cfg)

}

func main() {
	db, err := provider.ConnectDatabase(&cfg)

	if err != nil {
		panic(err)
	}

	personRepository := NewPersonRepository(db)

	personService := NewPersonService(personRepository)
	app := gin.Default()

	app.GET("/people", controllers.GetPeople(personService))
	app.GET("/people/:id", controllers.GetPerson)
	app.POST("/people", controllers.CreatePerson)
	app.PUT("/people/:id", controllers.UpdatePerson)
	app.DELETE("/people/:id", controllers.DeletePerson)
	app.Run()
}

// ProcessArgs processes and handles CLI arguments
func ProcessArgs(cfg *config.Config) Args {
	var a Args

	f := flag.NewFlagSet("Example server", 1)
	f.StringVar(&a.ConfigPath, "c", ".env", "Path to configuration file")

	fu := f.Usage
	f.Usage = func() {
		fu()
		envHelp, _ := cleanenv.GetDescription(cfg, nil)

		fmt.Fprintln(f.Output())
		fmt.Fprintln(f.Output(), envHelp)
	}

	f.Parse(os.Args[1:])
	return a
}
