package models


type Slide struct {
	Title    string   `json:"title"`
	Subtitle []string `json:"subtitle"`
}

type Presentation struct {
	Slides []Slide `json:"slides"`
}

type TopicRequest struct {
	Topic string `json:"topic"`
}