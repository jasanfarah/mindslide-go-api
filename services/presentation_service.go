package services

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jasanfarah/mindslide-go-api/models"
	"github.com/jasanfarah/mindslide-go-api/utils"
)

func AgendaCreator(c *gin.Context) models.Presentation {
	const agenda = `
	 Lav en agenda for præsentationen om %s på dansk
	- Returner agendaen i et JSON format
	- Giv titler til hver slide og subtitler til hver slide
	- Lav 2 slides
	- FOLLOW THIS STRUCTURE STRICTLY
	EXAMPLE:
	{
	  "slides": [
		{
		  "title": "Kunstig intelligens",
		  "subtitle": [
			"Hvad er kunstig intelligens?",
			"Hvordan bruges kunstig intelligens?",
			"Hvad er fremtiden for kunstig intelligens?"
		  ]
		},
		{
		  "title": "Historisk oversigt",
		  "subtitle": [
			"Udviklingen af kunstig intelligens gennem årene",
			"Vigtige milepæle i kunstig intelligens",
			"Indflydelse på forskellige brancher og samfundet"
		  ]
		},
		{
		  "title": "Konklusion",
		  "subtitle": [
			"Opsummering af nøglepunkter om kunstig intelligens",
			"Potentialet og udfordringerne ved at udnytte kunstig intelligens",
			"Betydningen af kunstig intelligens i fremtiden"
		  ]
		}
	  ]
	}
	`
	var topicRequest models.TopicRequest
	json.NewDecoder(c.Request.Body).Decode(&topicRequest)

	prompt := fmt.Sprintf(agenda, topicRequest.Topic)
	//fmt.Println(prompt)
	generatedAgendaJSON := utils.Generator(prompt)
	var presentation models.Presentation

	err := json.Unmarshal([]byte(generatedAgendaJSON), &presentation)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return presentation
}

func AgendaPresentationCreator(c *gin.Context) models.Presentation {
	//var presentation Presentation
	var presentation models.Presentation
	json.NewDecoder(c.Request.Body).Decode(&presentation)

	var FinalPresentation models.Presentation

  var wg sync.WaitGroup
	var mu sync.Mutex

	for _, slide := range presentation.Slides {
		var slideTitle = slide.Title
		var subTitles []string
		for _, subtitle := range slide.Subtitle {
			wg.Add(1)
			go func(subtitle string) {
				defer wg.Done()
				resultString := bulletPointCreator(slide.Title, subtitle)
				mu.Lock()
				subTitles = append(subTitles, resultString)
				mu.Unlock()
			}(subtitle)
		}
		wg.Wait()

		//fmt.Println("Results:")
 
		FinalPresentation.Slides = append(FinalPresentation.Slides, models.Slide{Title: slideTitle, Subtitle: subTitles})
	}

	//fmt.Println(FinalPresentation)
	return FinalPresentation

}

func PresentationCreator(c *gin.Context) models.Presentation {

	var presentation models.Presentation

	presentation = AgendaCreator(c)
//	fmt.Println(presentation)
	var FinalPresentation models.Presentation

  var wg sync.WaitGroup
	var mu sync.Mutex

	for _, slide := range presentation.Slides {
		var slideTitle = slide.Title
		var subTitles []string
		for _, subtitle := range slide.Subtitle {
			wg.Add(1)
			go func(subtitle string) {
				defer wg.Done()
				resultString := bulletPointCreator(slide.Title, subtitle)
				mu.Lock()
				subTitles = append(subTitles, resultString)
				mu.Unlock()
			}(subtitle)
		}
		wg.Wait()

		//fmt.Println("Results:")

		FinalPresentation.Slides = append(FinalPresentation.Slides, models.Slide{Title: slideTitle, Subtitle: subTitles})
	}

 //	fmt.Println(FinalPresentation)
	return FinalPresentation

}

func bulletPointCreator(mainTopic string, subTopic string) string {
	const prompt = `Skriv en sætning om emnet %s vedrørende %s på dansk`
	var concattedPrompt = fmt.Sprintf(prompt, subTopic, mainTopic)

	generatedbulletpoint := utils.Generator(concattedPrompt)
	//fmt.Println(generatedbulletpoint)
	return generatedbulletpoint

}
