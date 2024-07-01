package main

import (
	"github.com/MiguelGFerreira/gin-api-rest/database"
	"github.com/MiguelGFerreira/gin-api-rest/routes"
)

func main() {
	database.ConnectDatabase()
	routes.HandleRequest()
}
