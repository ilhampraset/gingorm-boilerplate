package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/ilhampraset/gingorm-boilerplate/config"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jinzhu/gorm"
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
	// provider.InitDb(&cfg)

}
func initDB(cfg *config.Config) *gorm.DB {

	db, err := gorm.Open(cfg.Database.Connection,
		fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
			cfg.Database.Username,
			cfg.Database.Password,
			cfg.Database.Host,
			cfg.Database.Name))

	if err != nil {
		fmt.Println(err)
	}
	return db
}
func main() {
	db := initDB(&cfg)

	defer db.Close()
	personController := InitPersonAPI(db)

	app := gin.Default()

	app.GET("/people", personController.GetPeople)
	app.GET("/people/:id", personController.GetPerson)
	app.POST("/people", personController.CreatePerson)
	app.PATCH("/people/:id", personController.UpdatePerson)
	app.DELETE("/people/:id", personController.DeletePerson)
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
