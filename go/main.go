package main

import (
	"./app/db"
	"./app/router"
)

func main() {

	db.Connect()
	defer db.Close()

	router.Router()
}
