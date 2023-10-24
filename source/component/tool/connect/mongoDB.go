package connect

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var MI MongoInstance

// function connect mongodb and return client
func ConnectMongoDB() {

	client, _ := mongo.NewClient(options.Client().ApplyURI(viper.GetString(`mongoDB.url`)))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_ = client.Connect(ctx)

	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Print("Error connecting to MongoDB: ", err)
	}

	fmt.Println("Database connected!")

	MI = MongoInstance{
		Client: client,
		DB:     client.Database(viper.GetString(`mongoDB.dbName`)),
	}

}
