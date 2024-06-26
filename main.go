package main

import "github.com/gin-gonic/gin"

func GetAllStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"name": "Tony Stark",
	})
}

func main() {
	r := gin.Default()
	r.GET("/students", GetAllStudents)
	r.Run()
}
