package handler

import (
	"github.com/johannesscr/api-gateway-example/micro/microservice"
	"github.com/johannesscr/api-gateway-example/src/includes"
	"log"
	"net/http"
)

// HandleHome is the health check for the server
func HandleHome(w http.ResponseWriter, r *http.Request) {
	msg := struct{
		Alive bool `json:"alive"`
	}{
		Alive: true,
	}
	res := includes.Resp{
		Status: 200,
		Message: "Welcome to Admin API Gateway",
		Data: msg,
	}
	res.Respond(w, r)
}

// HandleUserGet fetch a user
func HandleUserGet(w http.ResponseWriter, r *http.Request) {
	qp, _ := includes.QueryParams(w, r)
	userUUID := qp["userUuid"][0]

	ms := microservice.NewService()
	u, errors := ms.GetUser(userUUID)
	if errors != nil {
		log.Println(errors)
	}
	log.Println(u)

	msg := struct{
		Alive bool `json:"alive"`
	}{
		Alive: true,
	}
	res := includes.Resp{
		Status: 200,
		Message: "Welcome to Admin API Gateway",
		Data: msg,
	}
	res.Respond(w, r)
}