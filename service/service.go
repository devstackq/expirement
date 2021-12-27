package service

import (
	"inert/model"
	"inert/repository"
)

type Services struct {
	UserService *UserService
	PostService *PostService
}

type UserServiceInterface interface {
	GetUser()
	AddUser(model.Inert) (int, error)
}
type PostServiceInterface interface {
	GetPost()
	AddPost()
}

//dependency injection - любую базу могу
func InitServices(r repository.Repositories) *Services {
	return &Services{
		UserService: InitUserService(r.UserRepo),
		PostService: InitPostService(r.PostRepo),
	}
}
