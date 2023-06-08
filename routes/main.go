package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jasanfarah/mindslide-go-api/handlers"
)

func SetupRoutes( router *gin.Engine){
	router.POST("/generateAgenda", handlers.CreateAgendaHandler)
	router.GET("/test", handlers.TestHandler)
	router.POST("/test-generate-presentation", handlers.FakeTestPresentationHandler)

	router.POST("/generatePresentationFromAgenda", handlers.AgendaToPresentationHandler)
	router.POST("/generatePresentation", handlers.CreatePresentationHandler)

}