package service

type Student struct {
	ID       int
	College  string
	Facultet string
	Human
}

func (s Student) GetCollege() string {
	return s.College
}

// std := Std {ID: 1, College: "asd", Facultet:"123", model.Human{Name:"Marshall" }}

// std.GetName()
// std.GetCollege()

//open closed

//embeded - override func; || type ParamFunc  || interface{method()} for other struct

func (u *UserService) FilterUser(users []Student, fn func(Student) bool) (res []Student) {
	// users := s.userRepo.GetUsers
	for _, user := range users {
		if fn(user) {
			res = append(res, user)
		}
	}
	return res
}
