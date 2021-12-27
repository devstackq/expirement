package repository

import (
	"database/sql"
	"log"
)

type UserRepositorySql struct {
	db *sql.DB
}

func InitUserRepositorySql(db *sql.DB) UserRepositoryInterface {
	return &UserRepositorySql{db: db}
}

func (p *UserRepositorySql) AddUserDb() {
	log.Print("Sql repo layer, add user")
}
func (m *UserRepositorySql) GetUser() {
	log.Print("call mongo get user")
}
