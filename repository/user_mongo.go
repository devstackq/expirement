package repository

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryMongo struct {
	mongoDb *mongo.Client
}

func InitUserRepositoryMongo(db *mongo.Client) UserRepositoryInterface {
	return &UserRepositoryMongo{mongoDb: db}
}

func (m *UserRepositoryMongo) AddUserDb() {
	log.Print("repo layer, add user Mongo")
}

func (m *UserRepositoryMongo) GetUser() {
	log.Print("call mongo get user")
}
