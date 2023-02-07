package main

import (
	"fmt"
	"github.com/aasidxxx/mules/internal/app/server"
	"log"
)

var (
	id   int
	name string
)

func main() {
	fmt.Println("mu test")

	server := server.NewServer()

	//--- тест все юзеры из базы
	fmt.Println(server.Storage.UserRepository.SelectAll())

	log.Fatal(server.Start())
}
