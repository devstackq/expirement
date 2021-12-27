package repository

import (
	"database/sql"
	"log"
)

type PostRepositorySql struct {
	db *sql.DB
}

func InitPostRepositorySql(db *sql.DB) PostRepositoryInterface {
	return &PostRepositorySql{db: db}
}

func (p *PostRepositorySql) AddPostDb() {
	log.Print("repo layer, add post Sql")
}
