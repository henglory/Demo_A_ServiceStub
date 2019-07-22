package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/henglory/Demo_ServiceStubby/server"
	"github.com/henglory/Demo_ServiceStubby/service"
)

func main() {

	s := service.Service{}

	log.Println("Start server")
	serv := server.NewServer(s)
	serv.Start()
	defer serv.Close()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Server exiting")
}
