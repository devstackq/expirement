package model

import "time"

type Inert struct {
	ID           int            `json:"id"`
	Name         string         `json:"name"`
	Age          int            `json:"age"`
	Birthday     time.Time      `json:"birthday"`
	Participants map[int]string `json:"participants"`
}

type Test struct {
	Inert
}

// type Student struct {
// 	ID       int
// 	College  string
// 	Facultet string
// 	Human
// }

// func (h *Human) GetName() string {
// 	return h.Name
// }
// func (h *Human) SetName(name string) {
// 	h.Name = name
// }
