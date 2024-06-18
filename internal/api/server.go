// internal/api/server.go

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/koray.oezdemir/go-back-end/internal/service"
)

type Server struct {
	userController *UserController
}

func NewServer(userService service.UserService) *Server {
	userController := NewUserController(userService)
	return &Server{
		userController: userController,
	}
}

func (s *Server) SetupRoutes(router *gin.Engine) {
	s.userController.RegisterRoutes(router)
	// Füge hier weitere Controller hinzu
}



