package main

import (
	"github.com/kaikeventura/cat-food/ms-onboarding/configuration/database"
	"github.com/kaikeventura/cat-food/ms-onboarding/configuration/server"
)

func main() {
	startDatabase()
	startServer()
}

func startDatabase()  {
	database.RunDatabase()
}

func startServer()  {
	serverGin := server.NewServer()
	serverGin.Run()
}
