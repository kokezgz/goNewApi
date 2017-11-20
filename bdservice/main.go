package main

import (
	"log"

	"./pkg/api"
)

func main() {
	log.Println("Starting Server...")

	var controller api.Controller
	controller.StartServer()
}
