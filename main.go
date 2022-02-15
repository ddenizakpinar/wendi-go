package main

import (
	"net/http"

	"github.com/ddenizakpinar/wendi-go/controllers"
	"github.com/ddenizakpinar/wendi-go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Monster Dictionary")
	})

	r.GET("/monsters", controllers.FindMonsters)
	r.GET("/monsters/:id", controllers.FindMonster)
	r.POST("/monsters", controllers.CreateMonster)
	r.PATCH("/monsters/:id", controllers.UpdateMonster)
	r.DELETE("/monsters/:id", controllers.DeleteMonster)

	r.Run()
}
