package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jasanfarah/mindslide-go-api/services"
)

func AgendaToPresentationHandler(c *gin.Context) {
	presentationCreated := services.AgendaPresentationCreator(c)
	if presentationCreated.Slides == nil {
		c.JSON(400, gin.H{"error": "Could not generate Presentation"})
		return
	}

	c.JSON(200, presentationCreated)

}

func CreatePresentationHandler(c *gin.Context) {
	presentationCreated := services.PresentationCreator(c)
	if presentationCreated.Slides == nil {
		c.JSON(400, gin.H{"error": "Could not generate Presentation"})
		return
	}

	c.JSON(200, presentationCreated)

}
