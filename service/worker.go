package service

import (
	"log"
)

// type Wrk model.Worker
type Worker struct {
	ID       int
	Position string
	Company  string
	Human
}

func (w *Worker) GetCompany() string {
	return w.Company
}
func (w *Worker) SetCompany(company string) {
	w.Company = company
}

//override parent method, but i can use other method and field parent struct

func (w *Worker) SetName(name string) {
	w.Human.name = name
	//own realize, but not change parent methods,
}

//inherit parent struct, use him vars and methods
//open/closed - override child method

func TestWorker() {
	//model.Human{ID: 0, Name: "Noise", Age: 29}

	wrk := &Worker{}
	wrk.SetCompany("dream team")
	wrk.SetName("child own method name")
	// wrk.Human.SetName("Bill")

	n := wrk.Human.GetName() //parent method, access parent field
	c := wrk.GetCompany()    //own metod

	log.Println(n, c, "worker", wrk)
}
