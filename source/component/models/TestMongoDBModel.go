package models

import (
	"time"
)

type TestMongoDBModel struct {
	ID         string    `json:"id" bson:"id"`
	Name       string    `json:"name" bson:"name"`
	Surname    string    `json:"surname" bson:"surname"`
	Address    string    `json:"address" bson:"address"`
	InsertDate time.Time `json:"insert_date" bson:"insert_date"`
}

type TestMongoDBModelResponseError struct {
	StartData       time.Time `json:"startData"`
	EndData         time.Time `json:"endData"`
	Message         string    `json:"message"`
	CountRowSuccess int       `json:"countRowSuccess"`
}

type TestMongoDBModelResponse struct {
	StartData time.Time     `json:"startData"`
	EndData   time.Time     `json:"endData"`
	Message   []interface{} `json:"message"`
}
