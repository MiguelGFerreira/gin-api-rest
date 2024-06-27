package main

import (
	"github.com/MiguelGFerreira/gin-api-rest/models"
	"github.com/MiguelGFerreira/gin-api-rest/routes"
)

func main() {
	models.Heroes = []models.Hero{
		{Name: "Tony Stark", Ssn: "524-99-7106", Power: "Money"},
		{Name: "Steve Rogers", Ssn: "694-16-5643", Power: "Bravery"},
	}
	routes.HandleRequest()
}
