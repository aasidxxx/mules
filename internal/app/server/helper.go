package server

func (s *Server) LoadRoute() {

	//s.router.HandleFunc("/api/home", s.GetHome).Methods("GET")

	s.router.HandleFunc("/", s.GetHome).Methods("GET")
	s.router.HandleFunc("/all_user", s.AllUser).Methods("GET")
}
