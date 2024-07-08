package main

import (
	"Ushort/api"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	server := api.NewServer()
	go server.StartListening()
	<-signalChan

}
