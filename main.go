package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Константы для токена Telegram Bot API и URL API вопросов викторины.
const (
	TOKEN   = "Your token from botFather"
	siteURL = "https://the-trivia-api.com/v2/questions"
)

// Функция для создания новой конфигурации опроса для вопроса викторины.
func MyNewPoll(chatID int64, question string, someArray []string) tgbotapi.SendPollConfig {
	return tgbotapi.SendPollConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID: chatID,
		},
		Question:    question,
		Options:     someArray,
		IsAnonymous: true,
	}
}

// Функция для создания массива вариантов ответов для вопроса викторины.
func arrayOfQuestions(data QuestionsForQuiz) []string {
	var result = []string{data[0].CorrectAnswer}
	for i := 0; i < len(data[0].IncorrectAnswers); i++ {
		result = append(result, data[0].IncorrectAnswers[i])
	}
	rand.Shuffle(len(result), func(i, j int) { result[i], result[j] = result[j], result[i] })
	return result
}

// Функция для нахождения индекса правильного ответа в массиве вариантов ответов.
func CorrectAnswer(data QuestionsForQuiz, array []string) int64 {
	var result int
	for i := 0; i < len(array); i++ {
		if data[0].CorrectAnswer == array[i] {
			result = i
		}
	}
	return int64(result)
}

// Inline клавиатуры для Telegram бота.
var KeyboardForStart = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Start Quiz", "Start Quiz"),
		tgbotapi.NewInlineKeyboardButtonData("About", "About"),
	),
)

var KeyboardForStartQuiz = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Next Quiz", "Start Quiz"),
		tgbotapi.NewInlineKeyboardButtonData("Stop Quiz", "Stop Quiz"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			switch update.Message.Text {
			case "/start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет. Я бот для викторин и готов начать!")
				msg.ReplyMarkup = KeyboardForStart
				if _, err = bot.Send(msg); err != nil {
					panic(err)
				}
			case "/help":
				helpText := "Если у вас возникли сложности, вы можете написать @someHelper \n Или вы можете перезапустить бота командой /start"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, helpText)
				msg.ReplyMarkup = KeyboardForStart
				if _, err = bot.Send(msg); err != nil {
					panic(err)
				}
			}

		} else if update.CallbackQuery != nil {
			switch update.CallbackQuery.Data {
			case "Start Quiz":
				req, err := http.Get(siteURL)
				if err != nil {
					log.Panic(err)
				}
				var data QuestionsForQuiz
				err = json.NewDecoder(req.Body).Decode(&data)
				if err != nil {
					panic(err)
				}
				Quest := arrayOfQuestions(data)
				pollConfig := MyNewPoll(update.FromChat().ChatConfig().ChatID, data[1].Question.Text, Quest)
				pollConfig.CorrectOptionID = CorrectAnswer(data, Quest)
				pollConfig.IsAnonymous = true
				pollConfig.Type = "quiz"
				pollConfig.OpenPeriod = 30
				log.Println()
				msg := pollConfig
				msg.ReplyMarkup = KeyboardForStartQuiz
				_, err = bot.Send(msg)
				if err != nil {
					log.Panic(err)
				}

			case "About":
				sometext := "Этот бот был разработан для проекта, в настоящее время считается завершенным. \n Здесь вы можете выбрать некоторые команды: \n 1. /start \n 2. /help"
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, sometext)
				msg.ReplyMarkup = KeyboardForStart
				_, err = bot.Send(msg)
				if err != nil {
					log.Panic(err)
				}
			case "Stop Quiz":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Надеюсь, вам понравилось. Чтобы начать заново, нажмите кнопку (^˵◕ω◕˵^)")
				msg.ReplyMarkup = KeyboardForStart
				if _, err = bot.Send(msg); err != nil {
					panic(err)
				}

			}
		}
	}
}
