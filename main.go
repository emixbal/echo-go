package main

import (
	"echo-go/db"
	"echo-go/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":1234"))
}
