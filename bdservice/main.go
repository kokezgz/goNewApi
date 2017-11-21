package main

import (
	"log"

	"github.com/Koke/BC/bdservice/api"
)

func main() {
	log.Println("Starting Server...")

	var controller api.Controller
	controller.StartServer()
}
