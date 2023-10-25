package main

import (
	"catering-api/app/routes"
	"catering-api/helpers"
)

func main() {

	helpers.ConnectAWS()

	db := helpers.DatabaseConnect()

	server := routes.RouteService(db)

	server.Logger.Fatal(server.Start(":8002"))
}