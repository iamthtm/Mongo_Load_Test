package provider

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"example/testmongdb/component/models"
	"example/testmongdb/component/tool/connect"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testMongoDBProvider struct {
	app *fiber.App
}

func NewTestMongoDBProvider(app *fiber.App) *testMongoDBProvider {
	return &testMongoDBProvider{
		app: app,
	}
}

func (t *testMongoDBProvider) GetTestMongoDB() (*models.TestMongoDBModel, error) {
	db := connect.MI.DB.Collection(viper.GetString(`mongoDB.collectionName`))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//Find the document
	findOptions := options.Find()

	// Construct the filter for the query
	filter := bson.M{}
	setBatchSize, _ := strconv.Atoi(viper.GetString(`mongoDB.SetBatchSize`))
	findOptions.SetBatchSize(int32(setBatchSize))
	findResult, err := db.Find(ctx, filter, findOptions)
	if err != nil {
		fmt.Println("findResult err", err)
	}
	defer findResult.Close(ctx)
	var results models.TestMongoDBModel

	if err = findResult.All(ctx, &results); err != nil {
		fmt.Println("results err", err)
	}
	return &results, nil
}

// InsertManyMongoDB is a function
func (t *testMongoDBProvider) InsertsMongoGB(data []interface{}) ([]interface{}, error) {
	db := connect.MI.DB.Collection(viper.GetString(`mongoDB.collectionName`))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	status, err := db.InsertMany(ctx, data)
	if err != nil {
		fmt.Println("InsertMany err", err)
	}
	return status.InsertedIDs, nil
}

// insert One
func (t *testMongoDBProvider) InsertOneMongoGB(data models.TestMongoDBModel) (interface{}, error) {
	db := connect.MI.DB.Collection(viper.GetString(`mongoDB.collectionName`))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	status, err := db.InsertOne(ctx, data)
	if err != nil {
		fmt.Println("InsertOne err", err)
	}
	return status.InsertedID, nil
}
