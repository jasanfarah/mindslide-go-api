package handlers

import "github.com/jasanfarah/mindslide-go-api/models"
import "github.com/gin-gonic/gin"

func TestHandler(c *gin.Context) {
	c.JSON(200, "OK")
}

func FakeTestPresentationHandler(c *gin.Context) {
	presentation := models.Presentation{
		Slides: []models.Slide{
		  {
			Title: "Introduktion til kunstig intelligens",
			Subtitle: []string{
			  "Definition og historie af kunstig intelligens",
			  "Vigtige begreber af kunstig intelligens",
			  "Relevante eksempler på anvendelser",
			},
		  },
		  {
			Title: "Evolution af kunstig intelligens",
			Subtitle: []string{
			  "Begreber og teknikker relateret til kunstig intelligens",
			  "Omfattende problemer som løses med kunstig intelligens",
			  "Hvordan kunstig intelligens bruger maskinlæring til at forudsige, klassificere og genkende data",
			},
		  },
		  {
			Title: "Anvendelse af kunstig intelligens",
			Subtitle: []string{
			  "Grundlæggende elementer, som kunstig intelligens anvendes til at løse problemer",
			  "Hvordan kunstig intelligens bruges til at forbedre produktivitet og præcision",
			  "Brugssager i hverdagen som er muliggjort af kunstig intelligens",
			},
		  },
		  {
			Title: "Udfordringer for kunstig intelligens",
			Subtitle: []string{
			  "Væsentlige barrierer for brugen af kunstig intelligens",
			  "Betydningen af etisk og moralsk ansvar for kunstig intelligens",
			  "Hvor stor en betydning kunstig intelligens har på folks dagligdag",
			},
		  },
		  {
			Title: "Konklusion",
			Subtitle: []string{
			  "En sammenfattende gennemgang af kunstig intelligens' potentiale og udfordringer",
			  "Vigtigheden af etisk ansvarlighed, når det kommer til kunstig intelligens",
			  "Betydningen af kunstig intelligens for vores samfund i fremtiden",
			},
		  },
		},
	  }	  
	
	c.JSON(200, presentation)
}
