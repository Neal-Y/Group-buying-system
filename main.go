package main

import (
	"log"
	"shopping-cart/infrastructure"
	"shopping-cart/route"
)

func main() {

	dbErr := infrastructure.InitMySQL()
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	_, err := route.InitGinServer()
	if err != nil {
		log.Fatal(err)
	}
}
