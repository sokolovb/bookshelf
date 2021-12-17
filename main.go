package main

import (
	"flag"

	"github.com/sokolovb/bookshelf/rest"
)

func main() {
	port := flag.String("port", "8080", "Listen server port")
	flag.Parse()

	server := rest.NewServer(*port)
	server.Start()
}
