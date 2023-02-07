package main

import (
	"fmt"
	"github.com/aasidxxx/mules/internal/app/server"
	"github.com/aasidxxx/mules/storage"
	"log"
)

var (
	id   int
	name string
)

func main() {
	fmt.Println("mu test")

	t := storage.New()
	t.Open()
	fmt.Println(t.UserRepository)

	//подтягиваем объект юзеров
	usr := storage.NewUserInit(t)
	//fmt.Println(usr)

	t.UserRepository = usr

	allUser, _ := t.UserRepository.SelectAll()

	fmt.Println("+++++!!!!+++++")
	fmt.Println(allUser)

	for _, u := range allUser {
		fmt.Println(u)
	}

	server := server.NewServer()
	log.Fatal(server.Start())
}
