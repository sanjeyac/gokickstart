package app

import "fmt"

func (s *Server) RegisterRoutes() {
	fmt.Println("Registering routes")
	s.GetMapping("/", s.HandleIndex())
}
