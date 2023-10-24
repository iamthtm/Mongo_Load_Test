package service

import (
	"example/testmongdb/component/models"
	"example/testmongdb/component/provider"

	"github.com/gofiber/fiber/v2"
)

type TestMongoDBService struct {
	Client *fiber.App
}

// TestMongoDBService is a call provider and return string success or error
func (t *TestMongoDBService) InsertsMongoGB(data []interface{}) ([]interface{}, error) {
	TestMongoDBProvider := provider.NewTestMongoDBProvider(t.Client)
	return TestMongoDBProvider.InsertsMongoGB(data)
}

func (t *TestMongoDBService) GetMongoDB() (*models.TestMongoDBModel, error) {
	TestMongoDBProvider := provider.NewTestMongoDBProvider(t.Client)
	return TestMongoDBProvider.GetTestMongoDB()
}

func (t *TestMongoDBService) InsertOneMongoGB(data models.TestMongoDBModel) (interface{}, error) {
	TestMongoDBProvider := provider.NewTestMongoDBProvider(t.Client)
	return TestMongoDBProvider.InsertOneMongoGB(data)
}
