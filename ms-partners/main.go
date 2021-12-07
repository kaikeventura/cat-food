package main

import "github.com/kaikeventura/cat-food/ms-partners/configuration/server"

func main() {
	startServer()
}

func startServer() {
	serverGin := server.NewServer()
	serverGin.Run()
}
