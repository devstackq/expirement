package repository

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryInterface interface {
	AddUserDb()
	GetUser()
}
type PostRepositoryInterface interface {
	AddPostDb()
}

type Repositories struct {
	UserRepo UserRepositoryInterface
	PostRepo PostRepositoryInterface
}

//any db; di, dependency inversion

func InitSqlRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		UserRepo: InitUserRepositorySql(db),
		PostRepo: InitPostRepositorySql(db),
	}
}

func InitMongoRepositories(db *mongo.Client) *Repositories {
	return &Repositories{
		UserRepo: InitUserRepositoryMongo(db),
		PostRepo: InitPostRepositoryMongo(db),
	}
}
