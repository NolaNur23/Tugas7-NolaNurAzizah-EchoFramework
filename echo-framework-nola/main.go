package main

import (
	db "echo-framework-nola/db"
	routes "echo-framework-nola/routes"
)

func main() {

	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":3000"))
}
