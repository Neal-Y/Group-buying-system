package main

import (
	"log"
	"shopping-cart/config"
	"shopping-cart/infrastructure"
	"shopping-cart/route"
)

func main() {

	config.LoadConfig()

	dbErr := infrastructure.InitMySQL()
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	_, err := route.InitGinServer()
	if err != nil {
		log.Fatal(err)
	}
}
