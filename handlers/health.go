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
			Title:    "Introduktion til Elon Musk",
			Subtitle: []string{
			  "Elon Musk er en innovativ opfinder og entreprenør, der er kendt for at starte firmaer som Tesla Motors, SpaceX, Neuralink og The Boring Company.",
			  "Elon Musk har visioner der strækker sig ud over vores nuværende teknologi og åbner nye veje for bæredygtighed, innovation og den nyeste teknologi, der vil forandre vores fremtid.",
			},
		  },
		  {
			Title:    "Tesla og SpaceX",
			Subtitle: []string{
			  "Elon Musks langsigtede mål er at udfordre hele industrien ved hjælp af Tesla og SpaceX.",
			  "Tesla og SpaceX har fortsat visionen om at skabe bæredygtige køretøjer, der reducerer vores miljøpåvirkning og skaber et mere grønt og rent miljø.",
			  "Tesla og SpaceX arbejder sammen for et fælles mål om at skabe en kolonisering af Mars, som er en stor del af Elons Musks drøm om at demokratisere adgangen til rummet.",
			},
		  },
		},
	  }
	  
	
	c.JSON(200, presentation)
}