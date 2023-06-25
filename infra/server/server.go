package server

import (
	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/gin-api/infra/server/routes"
)

type Server struct {
	Host string
	Port int
	gin *gin.Engine
}

func NewRouter(port int, host string) *Server {
	return &Server{
		Host: host,
		Port: port,
	}
}

func (s *Server) setupRouter() {
	r := gin.Default()
	routes.InitUser(r)
	s.gin = r
}

func (s *Server) Start() {
	s.setupRouter()
	s.gin.Run()
}