package main

import (
	"echo-api/db"
	"echo-api/router"
)

func main() {
	db.NewData()

	e := router.Init()

	e.Logger.Fatal(e.Start(":1324"))
}
