package main

import (
	"github.com/kaikeventura/cat-food/ms-onboarding/configuration/container"
	"github.com/kaikeventura/cat-food/ms-onboarding/configuration/database"
	"github.com/kaikeventura/cat-food/ms-onboarding/configuration/server"
)

func main() {
	startDatabase()
	startDependencyInjection()
	startServer()
}

func startDatabase()  {
	database.RunDatabase()
}

func startDependencyInjection() {
	container.RunDependencyInjection()
}

func startServer()  {
	serverGin := server.NewServer()
	serverGin.Run()
}
