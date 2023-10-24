package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/spf13/viper"

	//Controller
	"example/testmongdb/component/controller"

	//redis
	"example/testmongdb/component/tool/connect"
)

// function read file config.json
func init() {
	viper.SetConfigFile(`settings/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	app := fiber.New()

	app.Use(
		compress.New(),
		cors.New(),
		recover.New(),
	)
	//init mongodb
	connect.ConnectMongoDB()
	controller.TestMongoDBController(app)

	app.Listen(fmt.Sprintf(`:%s`, viper.GetString(`app.port`)))
}
