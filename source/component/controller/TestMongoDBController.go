package controller

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"

	"example/testmongdb/component/models"
	"example/testmongdb/component/service"
	"math/rand"
)

type testMongoDBController struct {
	TestMongoDBService service.TestMongoDBService
}

func TestMongoDBController(app *fiber.App) {
	testMongoDBController := testMongoDBController{
		TestMongoDBService: service.TestMongoDBService{
			Client: app,
		},
	}

	GroupURL := app.Group(viper.GetString(`app.context`))
	GroupURL.Get("/insertAll", testMongoDBController.InsertAllMongoGB)
	GroupURL.Get("/insert", testMongoDBController.InsertMongoGB)
	GroupURL.Get("/get", testMongoDBController.GetMongoDB)
}

func (r *testMongoDBController) InsertAllMongoGB(c *fiber.Ctx) error {
	startData := time.Now()
	insertManyBatchSize, _ := strconv.Atoi(c.Query("InsertManyBatchSize"))
	serverName := c.Query("ServerName")
	//check query is empty
	if insertManyBatchSize == 0 {
		insertManyBatchSize, _ = strconv.Atoi(viper.GetString(`mongoDB.InsertManyBatchSize`))
	}

	datas := []interface{}{}
	for i := 0; i < insertManyBatchSize; i++ {
		var data = models.TestMongoDBModel{
			ID:         fmt.Sprintln(i) + serverName,
			Name:       randomString(20),
			Surname:    randomString(20),
			Address:    randomString(256),
			InsertDate: time.Now(),
		}
		datas = append(datas, data)
	}
	status, err := r.TestMongoDBService.InsertsMongoGB(datas)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		resp := models.TestMongoDBModelResponseError{
			StartData: startData,
			EndData:   time.Now(),
			Message:   err.Error(),
		}
		return c.JSON(resp)
	}
	resp := models.TestMongoDBModelResponse{
		StartData: startData,
		EndData:   time.Now(),
		Message:   status,
	}
	return c.JSON(resp)

}

func (r *testMongoDBController) InsertMongoGB(c *fiber.Ctx) error {
	startData := time.Now()
	insertManyBatchSize, _ := strconv.Atoi(c.Query("InsertManyBatchSize"))
	serverName := c.Query("ServerName")
	//check query is empty
	if insertManyBatchSize == 0 {
		insertManyBatchSize, _ = strconv.Atoi(viper.GetString(`mongoDB.InsertManyBatchSize`))
	}
	status := []interface{}{}

	for i := 0; i < insertManyBatchSize; i++ {
		data := models.TestMongoDBModel{
			ID:         fmt.Sprintln(i) + serverName,
			Name:       randomString(20),
			Surname:    randomString(20),
			Address:    randomString(256),
			InsertDate: time.Now(),
		}

		datas, err := r.TestMongoDBService.InsertOneMongoGB(data)
		status = append(status, datas)
		if err != nil {
			c.SendStatus(fiber.StatusBadRequest)
			resp := models.TestMongoDBModelResponseError{
				StartData:       startData,
				EndData:         time.Now(),
				Message:         err.Error(),
				CountRowSuccess: i + 1,
			}
			return c.JSON(resp)
		}
	}
	res := models.TestMongoDBModelResponse{
		StartData: startData,
		EndData:   time.Now(),
		Message:   status,
	}

	return c.JSON(res)

}

func (r *testMongoDBController) GetMongoDB(c *fiber.Ctx) error {
	startData := time.Now()
	datas, err := r.TestMongoDBService.GetMongoDB()
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		resp := models.TestMongoDBModelResponseError{
			StartData: startData,
			EndData:   time.Now(),
			Message:   err.Error(),
		}
		return c.JSON(resp)
	}
	return c.JSON(datas)
}

// randomString is a function
func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
