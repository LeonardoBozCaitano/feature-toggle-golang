package server

func (s *Server) routes() {
	s.Router.HandleFunc("/feature", s.HandleFeatureGetAll()).Methods("GET")
	s.Router.HandleFunc("/feature", s.HandleFeatureInsert()).Methods("POST")
	s.Router.HandleFunc("/feature/{name}/{client}", s.HandleFeatureClientVerification()).Methods("GET")
}
