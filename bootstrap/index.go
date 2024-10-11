package bootstrap

import (
	"github.com/naythukhant/gofiber-todos/configs"
	"github.com/naythukhant/gofiber-todos/services"
)

func Bootstrap() {
	configs.Init()
	services.InitDatabase()
}
