package main

import (
	"github.com/kaikeventura/cat-food/ms-partner/configuration/container"
	"github.com/kaikeventura/cat-food/ms-partner/configuration/server"
)

func main() {
	startDependencyInjection()
	startServer()
}

func startServer() {
	serverGin := server.NewServer()
	serverGin.Run()
}

func startDependencyInjection()  {
	container.RunDependencyInjection()
}
