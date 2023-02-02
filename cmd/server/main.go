package main

import (
	"fmt"
	"github.com/aasidxxx/mules/internal/app/server"
	"log"
)

func main() {
	fmt.Println("mu test")

	server := server.NewServer()

	log.Fatal(server.Start())

}
