package main

import (
	"golang/db"
	"golang/routes"
)

func main() {

	db.Init()

	e := routes.Init()
	e.Logger.Fatal(e.Start(":911"))

}
