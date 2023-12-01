package main

import (
	"4nc3str4l/my-blockchain/server"
	"flag"
	"log"
)

func init() {
	log.SetPrefix("🔗 Blockchain: ")
}

func main() {
	port := flag.Uint("port", 5000, "TCP Port Number for Blockchain Server")
	flag.Parse()
	log.Println(*port)

	app := server.NewBlockchainServer(uint16(*port))
	app.Run()
}
