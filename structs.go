package main

// Структура данных для вопросов викторины, полученных из API.
type QuestionsForQuiz []struct {
	Category         string   `json:"category"`
	ID               string   `json:"id"`
	CorrectAnswer    string   `json:"correctAnswer"`
	IncorrectAnswers []string `json:"incorrectAnswers"`
	Question         struct {
		Text string `json:"text"`
	} `json:"question"`
	Tags       []string      `json:"tags"`
	Type       string        `json:"type"`
	Difficulty string        `json:"difficulty"`
	Regions    []interface{} `json:"regions"`
	IsNiche    bool          `json:"isNiche"`
}

type BaseChat struct {
	ChatID                   int64
	ChannelUsername          string
	ReplyToMessageID         int
	ReplyMarkup              interface{}
	DisableNotification      bool
	AllowSendingWithoutReply bool
}
