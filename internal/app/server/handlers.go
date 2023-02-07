package server

import (
	"fmt"
	"net/http"
)

func (s *Server) GetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("----home ")
	w.Write([]byte("корневой роут"))
}

func (s *Server) AllUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("----home ")
	w.Write([]byte("корневой роут"))
}
