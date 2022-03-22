package routes

import (
	"github.com/Aibier/coinmarket-service/controllers"

	"github.com/gin-gonic/gin"
)

// TwitRoute ...
func TwitRoute(router *gin.Engine) {
	router.GET("/sync", controllers.SyncRecords())
	router.GET("/twits", controllers.GetAllTRecords())
}
