package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Server struct {
		Port int `yaml:"port"`
	}
}

type Server struct {
	Router *mux.Router
	Config *ServerConfig
}

/** Init server configuration */
func (s *Server) Init() {
	s.Router = mux.NewRouter()
	staticFiles := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	s.Router.PathPrefix("/static/").Handler(staticFiles)

	yamlFile, err := ioutil.ReadFile("./configs/properties.yaml")
	if err != nil {
		panic("Error no configuration file found")
	}

	err = yaml.Unmarshal(yamlFile, &s.Config)
	if err != nil {
		panic("Error reading configuration file")
	}
}

/** Start server */
func (s *Server) Start() {
	log.Println("Start server on", s.Config.Server.Port)
	var addr = fmt.Sprintf(":%d", s.Config.Server.Port)
	http.Handle("/", s.Router)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func (s *Server) GetMapping(path string, f func(http.ResponseWriter, *http.Request)) {
	s.Router.NewRoute().Path(path).HandlerFunc(f).Methods("GET")
}

func (s *Server) PostMpaaing(path string, f func(http.ResponseWriter, *http.Request)) {
	s.Router.NewRoute().Path(path).HandlerFunc(f).Methods("POST")
}

func (s *Server) PutMapping(path string, f func(http.ResponseWriter, *http.Request)) {
	s.Router.NewRoute().Path(path).HandlerFunc(f).Methods("PUT")
}

func (s *Server) DeleteMapping(path string, f func(http.ResponseWriter, *http.Request)) {
	s.Router.NewRoute().Path(path).HandlerFunc(f).Methods("DELETE")
}

func (s *Server) PatchMapping(path string, f func(http.ResponseWriter, *http.Request)) {
	s.Router.NewRoute().Path(path).HandlerFunc(f).Methods("PATCH")
}
