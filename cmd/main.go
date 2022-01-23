// go get github.com/urfave/cli/v2

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/marcboeker/mega/app"
)

func main() {
	app, err := app.Initialize()
	if err != nil {
		log.Fatalf("could not initialize app: %s\n", err)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("shutting down...")
		app.Shutdown()
		os.Exit(1)
	}()

	if err := app.Run(); err != nil {
		log.Fatalf("could not run app: %s\n", err)
	}
}
