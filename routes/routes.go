package routes

import (
	"github.com/MiguelGFerreira/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/heroes", controllers.GetAllHeroes)
	r.GET("/:name", controllers.Greetings)
	r.POST("/heroes", controllers.CreateHero)
	r.Run()
}
