package models

type Step struct {
	Text  string `json:"text"`
	Image string `json:"image"`
}

type Recipe struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Image       string   `json:"image"`
	Cuisine     string   `json:"cuisine"`
	Ingredients []string `json:"ingredients"`
	Views       int      `json:"views"`
	Steps       []Step   `json:"steps"`
}