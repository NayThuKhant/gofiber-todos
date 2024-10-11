package services

import (
	"fmt"

	"github.com/naythukhant/gofiber-todos/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DATABASE *gorm.DB

func InitDatabase() {

	var err error

	datasourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		configs.Database.Username,
		configs.Database.Password,
		configs.Database.Host,
		configs.Database.Port,
		configs.Database.Database,
	)

	DATABASE, err = gorm.Open(mysql.Open(datasourceName), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
}
