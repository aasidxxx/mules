package server

import (
	"log"
	"net/http"
)

type Server struct {
	port string
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("mu test 222"))
}

func (s *Server) Start() error {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")

	return http.ListenAndServe(":4000", mux)

	//log.Fatal(err)

}

func NewServer() *Server {
	return &Server{
		port: ":8080",
	}
}
