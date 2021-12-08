package main

import (
	"github.com/kaikeventura/cat-food/ms-partner/configuration/container"
	"github.com/kaikeventura/cat-food/ms-partner/configuration/database"
	"github.com/kaikeventura/cat-food/ms-partner/configuration/server"
)

func main() {
	startDatabase()
	startDependencyInjection()
	startServer()
}

func startDatabase() {
	database.RunDatabase()
}

func startServer() {
	serverGin := server.NewServer()
	serverGin.Run()
}

func startDependencyInjection() {
	container.RunDependencyInjection()
}
