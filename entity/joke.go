package entity

type Joke struct {
	JokeId       string `json:"jokeId"`
	Content      string `json:"content"`
	Category     string `json:"category"`
	JokeSetup    string `json:"jokeSetup"`
	JokeDelivery string `json:"jokeDelivery"`
	Language     string `json:"language"`
}

type JokeContent struct {
	JokeId  string `json:"jokeId"`
	Content string `json:"content"`
}
