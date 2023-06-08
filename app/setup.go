package app

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jasanfarah/mindslide-go-api/routes"
)



func SetupAPI() {
	fmt.Println("Starting server...")

	router := gin.Default()
	routes.SetupRoutes(router)

	router.Use(cors.Default())

	router.Run()
	
}
