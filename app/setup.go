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
	router.Use(cors.Default())

	routes.SetupRoutes(router)


	router.Run()
	
}
