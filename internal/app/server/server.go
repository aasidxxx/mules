package server

import (
	"fmt"
	"github.com/aasidxxx/mules/storage"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	port    string
	router  *mux.Router
	Storage *storage.Storage
}

func (s *Server) Start() error {

	log.Println("Запуск веб-сервера на http://127.0.0.1:4001")

	//--- загружаем роуты
	s.LoadRoute()

	return http.ListenAndServe(":4001", s.router)

	//log.Fatal(err)

}

func NewServer() *Server {

	//--- подтягиваем  storage
	t := storage.New()
	t.Open()
	fmt.Println(t.UserRepository)

	//--- подтягиваем объект юзеров
	usr := storage.NewUserInit(t)

	t.UserRepository = usr

	return &Server{
		port:    ":8080",
		router:  mux.NewRouter(),
		Storage: t,
	}
}
