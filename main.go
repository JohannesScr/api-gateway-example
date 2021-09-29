package main

import (
	"github.com/johannesscr/api-gateway-example/src/includes"
	"github.com/johannesscr/api-gateway-example/src/server"
	"log"
	"net/http"
)

func main() {
	env := includes.Env{}
	env.Load(".env")


	s := server.NewServer()
	log.Fatal(http.ListenAndServe(":8080", s.Router))
}
