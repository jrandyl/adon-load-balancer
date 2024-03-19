package main

import (
	"log"

	"github.com/jrandyl/adon-load-balancer/server"
)

func main() {
	err := server.Start(":443")
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
