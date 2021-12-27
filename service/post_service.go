package service

import (
	"inert/model"
	"inert/repository"
	"log"
)

type PostService struct {
	postRepo repository.PostRepositoryInterface
}

func InitPostService(r repository.PostRepositoryInterface) *PostService {
	return &PostService{postRepo: r}
}

func (s *PostService) AddPost() {
	s.postRepo.AddPostDb()
	log.Println("call add post service", s)
}
func (s *PostService) GetUser(m model.Inert) {
	log.Print("call get post service")
	// s.postRepo.GetPost()
}
