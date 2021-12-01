package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/cat-food/ms-onboarding/configuration/server/routes"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/adapters/inbound/controllers"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/application/services"
	"log"
)

type Server struct {
	port string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port: "5000",
		server: gin.Default(),
	}
}

func (server *Server) Run() {
	dependencyInjection()
	router := routes.ConfigurationRouter(server.server)

	log.Print("Server is running at port: ", server.port)
	log.Fatal(router.Run(":" + server.port))
}

func dependencyInjection()  {
	controllers.ConstructUserController(services.ConstructUserService())
}
