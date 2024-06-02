package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "webapi/cmd/config"
	application "webapi/internal/app"
)

func main() {
	// Noop 3

	app := application.NewApp()
	// Listen to interrupt signals for graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go app.Run()

	// Wait for an interrupt signal
	<-interrupt
	log.Println("Shutting down the server...")
}
