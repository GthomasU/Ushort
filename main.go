package main

import (
	"Ushort/api"
	_ "Ushort/docs"
	"os"
	"os/signal"
	"syscall"
)

// @title Ushort API
// @version 1.0
// @description This is a Ushort API documentation.
// @contact.name API Support
// @contact.email gabrielthomas300@gmail.com
func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	server := api.NewServer()
	go server.StartListening()
	<-signalChan

}
