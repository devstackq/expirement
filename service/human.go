package service

type Human struct {
	id   int
	name string
	age  int
}

// type Human model.Human

func (h *Human) GetName() string {
	return h.name
}

func (h *Human) SetName(name string) {
	h.name = name
}

// func (h *Human) Random() int {
// 	return 10
// }

//base class A; child class B - have access variable and method
