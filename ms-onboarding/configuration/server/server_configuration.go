package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/cat-food/ms-onboarding/configuration/server/routes"
	"log"
	"os"
)

type Server struct {
	port string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port: os.Getenv("SERVER_PORT"),
		server: gin.Default(),
	}
}

func (server *Server) Run() {
	router := routes.ConfigurationRouter(server.server)

	log.Print("Server is running at port: ", server.port)
	log.Fatal(router.Run(":" + server.port))
}
