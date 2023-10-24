package controller

import (
	"example/testmongdb/component/service"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

// redisController is a struct
type redisController struct {
	redisService service.RedisService
}

// RedisController is a route endpoint
func RedisController(app *fiber.App, connectRedis *redis.Client) {

	redisController := redisController{
		redisService: service.RedisService{
			Redis: connectRedis,
		},
	}

	app.Get("/redis", redisController.GetRedis)
}

// GetRedis is a function
func (r *redisController) GetRedis(c *fiber.Ctx) error {
	data, _ := r.redisService.GetRedis()
	return c.JSON(fiber.Map{
		"message": data,
	})
}
