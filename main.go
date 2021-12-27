package main

import (
	"inert/handler"
	"inert/repository"
	"inert/service"
	"log"
	"net/http"
)

func main() {
	// repoMongo := repository.InitMongoRepositories(nil)
	repoSql := repository.InitSqlRepositories(nil)
	s := service.InitServices(*repoSql)
	s.UserService.GetFilteredUser() // for test, embedding

	hr := handler.InitInertHandler(*s)

	http.HandleFunc("/add", hr.AddUser)
	// http.HandleFunc("update", hr)
	log.Print("start")
	http.ListenAndServe(":6969", nil)
}
