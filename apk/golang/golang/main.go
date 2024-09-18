package main

import "golang/routes"

func main() {

	e := routes.Init()
	e.Logger.Fatal(e.Start(":911"))

}
