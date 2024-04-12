package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("server is starting...")
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	log.Println("server is ready to listen and serve")
	<-stop
	log.Println("server is shutting down...")
}
