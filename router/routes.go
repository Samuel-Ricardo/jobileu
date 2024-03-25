package router

import (
	"github.com/Samuel-Ricardo/jobileu/handler"
	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "api/v1"

	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
	}
}
