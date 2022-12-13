package main

import (
	"workshop-management/db"
	"workshop-management/server"
)

func main() {

	db.Connect()
	db.GetShelfs()

	server.PrintTeste("opaaa")
	server.Start()

}
