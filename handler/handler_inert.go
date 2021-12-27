package handler

import (
	"encoding/json"
	"inert/model"
	"inert/service"
	"io/ioutil"
	"log"
	"net/http"
)

type Handler struct {
	Service service.Services
}

func InitInertHandler(service service.Services) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	var in model.Inert
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		log.Println("read all err")
		return
	}

	err = json.Unmarshal(b, &in)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		log.Println("json err handler")
		return
	}

	log.Println(in, "read obj")
	lid, err := h.Service.UserService.AddUser(in)
	log.Println(lid, err, "service res")
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err.Error())
		log.Println("service err")
		return
	}
	log.Print("ok handler side")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(lid)

}
