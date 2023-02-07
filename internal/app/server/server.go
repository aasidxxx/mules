package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	port   string
	router *mux.Router
}

func (s *Server) Start() error {

	log.Println("Запуск веб-сервера на http://127.0.0.1:4001")

	//--- загружаем роуты
	s.LoadRoute()

	return http.ListenAndServe(":4001", s.router)

	//log.Fatal(err)

}

func NewServer() *Server {
	return &Server{
		port:   ":8080",
		router: mux.NewRouter(),
	}
}
