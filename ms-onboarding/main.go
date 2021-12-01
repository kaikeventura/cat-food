package main

import "github.com/kaikeventura/cat-food/ms-onboarding/configuration/server"

func main() {
	serverGin := server.NewServer()
	serverGin.Run()
}
