package service

import (
	"inert/model"
	"inert/repository"
	"log"
)

type UserService struct {
	userRepo repository.UserRepositoryInterface
}

func InitUserService(r repository.UserRepositoryInterface) *UserService {
	return &UserService{userRepo: r}
}

func (u *UserService) AddUser(m model.Inert) (int, error) {
	TestWorker()
	log.Println("call add user service", m)
	u.userRepo.AddUserDb()
	return 1, nil
}

func (u *UserService) GetFilteredUser() {
	type paramFn func(Student) bool
	h := Human{}
	h.SetName("asd")
	//1 change; 2 !change
	// stds := u.userRepo.GetStudents()
	listUsers := []Student{
		{1, "Mit", "sad", h},
		{2, "Aiu", "dasd", Human{}},
	}

	param := paramFn(func(s Student) bool {
		return s.College == "Mit"
	})
	resFil := u.FilterUser(listUsers, param)

	log.Print(resFil, 2)

}

// func (s *UserService) GetUser(m model.Inert) {
// 	log.Print("call get user service")
// 	s.userRepo.GetUser()
// }
