package app

func (s *Server) RegisterRoutes() {
	s.GetMapping("/", s.HandleIndex())
}
