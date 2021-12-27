package repository

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepositoryMongo struct {
	mongoDb *mongo.Client
}

func InitPostRepositoryMongo(db *mongo.Client) PostRepositoryInterface {
	return &PostRepositoryMongo{mongoDb: db}
}

func (p *PostRepositoryMongo) AddPostDb() {
	log.Print("repo layer, add post Mongo")
}
