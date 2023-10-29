package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/kisstc/microblog/internal/server"
)

func main() {
	serv, err := server.New("8080")
	if err != nil {
		log.Fatal(err)
	}

	// start the server
	go serv.Start()

	// wait for an interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// si todo sale bien, cerrar
	serv.Close()
}
