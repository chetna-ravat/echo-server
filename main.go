package main

import (
	"flag"
	"log"

	"github.com/chetna-ravat/echo-server/config"
	"github.com/chetna-ravat/echo-server/server"
)

func setupFlag() {

	flag.StringVar(&config.Host, "host", "0.0.0.0", "host ip address of the echo server")
	flag.IntVar(&config.Port, "port", 5200, "port on which echo server is listening")
	flag.Parse()
}

func main() {
	setupFlag()
	log.Println("Starting the echo server ...")
	server.RunSyncServer()
}
