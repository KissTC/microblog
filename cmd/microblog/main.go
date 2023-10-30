package main

import (
	"log"
	"os"
	"os/signal"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kisstc/microblog/internal/server"
)

func main() {

	port := os.Getenv("PORT")

	serv, err := server.New(port)
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
