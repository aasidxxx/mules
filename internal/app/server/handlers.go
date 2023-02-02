package server

import (
	"fmt"
	"net/http"
)

func (s *Server) GetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("----home ")
	w.Write([]byte("mu test 222"))

}
