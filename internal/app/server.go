package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ServerConfig struct {
	Server struct {
		Port int `yaml:"port"`
	}
	Datasource struct {
		SqliteDb string `yaml:"sqlite-db"`
	}
}

type Server struct {
	Router *mux.Router
	Config *ServerConfig
	Db     *gorm.DB
}

/** Init server configuration */
func (s *Server) Init() {
	log.Println("Setup routing")
	s.Router = mux.NewRouter()
	staticFiles := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	s.Router.PathPrefix("/static/").Handler(staticFiles)

	log.Println("Reading application properties")
	yamlFile, err := ioutil.ReadFile("./configs/properties.yaml")
	if err != nil {
		log.Fatalln(err)
		panic("Error no configuration file found")
	}

	err = yaml.Unmarshal(yamlFile, &s.Config)
	if err != nil {
		log.Fatalln(err)
		panic("Error reading configuration file")
	}

	// can be changed with a different datbase here
	log.Println("Opening database", s.Config.Datasource.SqliteDb)
	if len(s.Config.Datasource.SqliteDb) == 0 {
		panic("Missing database in properties")
	}

	s.Db, err = gorm.Open(sqlite.Open(s.Config.Datasource.SqliteDb), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		panic("Error connecting to database")
	}

}

/** Start server */
func (s *Server) Start() {
	log.Println("Start server on", s.Config.Server.Port)
	var addr = fmt.Sprintf(":%d", s.Config.Server.Port)
	http.Handle("/", s.Router)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func (s *Server) RequestMapping(path string, method string, f func(http.ResponseWriter, *http.Request)) {
	s.Router.NewRoute().Path(path).HandlerFunc(f).Methods(method)
}

func (s *Server) GetMapping(path string, f func(http.ResponseWriter, *http.Request)) {
	s.RequestMapping(path, "GET", f)
}

func (s *Server) PostMpaaing(path string, f func(http.ResponseWriter, *http.Request)) {
	s.RequestMapping(path, "POST", f)
}

func (s *Server) PutMapping(path string, f func(http.ResponseWriter, *http.Request)) {
	s.RequestMapping(path, "PUT", f)
}

func (s *Server) DeleteMapping(path string, f func(http.ResponseWriter, *http.Request)) {
	s.RequestMapping(path, "DELETE", f)
}

func (s *Server) PatchMapping(path string, f func(http.ResponseWriter, *http.Request)) {
	s.RequestMapping(path, "PATCH", f)
}
