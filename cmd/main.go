package main

import (
	"github.com/Aibier/coinmarket-service/configs"
	"github.com/Aibier/coinmarket-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//run database
	configs.ConnectDB()
	//routes
	routes.TwitRoute(router)

	err := router.Run("localhost:8000")
	if err != nil {
		return 
	}
}
