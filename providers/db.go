package providers

import (
	"fmt"
	"github.com/ilhampraset/gingorm-boilerplate/config"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDb(cfg *config.Config) {
	fmt.Println(cfg.Database.Host)

	database, err := gorm.Open(cfg.Database.Connection, fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Name))
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to Database")
	}

	DB = database
}
