package main

import (
	"flag"
	"log"
)

func init() {
	log.SetPrefix("🔗 Blockchain: ")
}

func main() {
	log.Println("Hello!")
	port := flag.Uint("port", 5000, "TCP Port Number for Blockchain Server")
	flag.Parse()
	log.Println(*port)
}
