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

//func home(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("mu test 222"))
//}

func (s *Server) Start() error {

	//mux := http.NewServeMux()
	//mux.HandleFunc("/", home)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	//--- загружаем роуты
	s.LoadRoute()

	return http.ListenAndServe(":4000", s.router)

	//log.Fatal(err)

}

func NewServer() *Server {
	return &Server{
		port:   ":8080",
		router: mux.NewRouter(),
	}
}
