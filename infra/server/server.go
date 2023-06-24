package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/Kimoto-Norihiro/gin-api/repositories"
	"github.com/Kimoto-Norihiro/gin-api/handlers"
	"github.com/Kimoto-Norihiro/gin-api/usecases"
)

type Server struct {
	DB *gorm.DB
	Host string
	Port int
	gin *gin.Engine
}

func NewRouter(port int, host string, db *gorm.DB) *Server {
	return &Server{
		DB: db,
		Host: host,
		Port: port,
	}
}

func (s *Server) setupRouter() {
	r := gin.Default()
	repo := repositories.NewUserRepository(s.DB)
	usecase := usecases.NewUserUsecase(repo)
	userHandler := handlers.NewUserHandler(usecase)

	r.POST("/users", userHandler.CreateUser)
	s.gin = r
}

func (s *Server) Start() {
	s.setupRouter()
	s.gin.Run()
}