package main

import (
	"catering-api/app/routes"
	"catering-api/helpers"
)

func main() {
	db := helpers.DatabaseConnect()

	server := routes.RouteService(db)

	server.Logger.Fatal(server.Start(":8002"))
}