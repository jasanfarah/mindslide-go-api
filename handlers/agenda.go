package handlers

import (
"github.com/gin-gonic/gin"
"github.com/jasanfarah/mindslide-go-api/services"
)


func CreateAgendaHandler(c *gin.Context) {
	presentation := services.AgendaCreator(c)
	if presentation.Slides == nil {
		c.JSON(400, gin.H{"error": "Could not generate agenda"})
		return
	}
	c.JSON(200, presentation)
}