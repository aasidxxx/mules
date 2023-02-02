package server

func (s *Server) LoadRoute() {

	s.router.HandleFunc("/api/home", s.GetHome).Methods("GET")
}
