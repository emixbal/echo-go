package main

import (
	"echo-go/routes"
)

func main() {
	e := routes.Init()
	e.Logger.Fatal(e.Start(":1234"))
}
