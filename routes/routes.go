package routes

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.Static("/js", "./js")
	r.Static("/icons", "./icons")
	r.GET("/", controllers.ShowIndexPage)
	r.GET("/:id", controllers.VerificaSitePorId)
	r.POST("/", controllers.VerificaSites)
	r.DELETE("/remove/:id", func(ctx *gin.Context) {
		controllers.DeleteSiteFromList(ctx)
	})
	r.Run()
}
