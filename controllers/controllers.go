package controllers

import (
	"net/http"

	"github.com/MiguelGFerreira/gin-api-rest/database"
	"github.com/MiguelGFerreira/gin-api-rest/models"
	"github.com/gin-gonic/gin"
)

func GetAllHeroes(c *gin.Context) {
	var heroes []models.Hero
	database.DB.Find(&heroes)
	c.JSON(200, heroes)
}

func Greetings(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"Api says: ": name + " is a legend",
	})
}

func CreateHero(c *gin.Context) {
	var hero models.Hero
	if err := c.ShouldBindJSON(&hero); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	database.DB.Create(&hero)
	c.JSON(http.StatusOK, hero)
}

func GetHeroById(c *gin.Context) {
	var hero models.Hero
	id := c.Params.ByName("id")
	database.DB.First(&hero, id)

	if hero.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Hero not found",
		})
		return
	}

	c.JSON(http.StatusOK, hero)
}
