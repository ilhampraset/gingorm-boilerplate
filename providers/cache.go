package providers

import (
	"fmt"
	. "github.com/ilhampraset/gingorm-boilerplate/config"
	"github.com/jinzhu/gorm"
)

func ConnectDatabase(cfg *Config) (*gorm.DB, error) {
	return gorm.Open(cfg.Database.Connection,
		fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
			cfg.Database.Username,
			cfg.Database.Password,
			cfg.Database.Host,
			cfg.Database.Name))

}
