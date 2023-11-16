# TelegramQuizBot
Этот проект представляет собой Telegram бота для викторин. Бот получает вопросы из [Open Trivia Database](https://the-trivia-api.com/) и предоставляет пользователю возможность участвовать в викторине.
## Установка и запуск
1. Установите необходимые зависимости с помощью команды:

    ```go
   go get -u github.com/go-telegram-bot-api/telegram-bot-api/v5
   ```
    
2. Запустите бота, используя следующую команду:

   ```go
   go run main.go
   ```
   
## Использование
* /start - начать викторину
* /help - получить справку и поддержку

## Команды викторины
* Start Quiz - Начать новый раунд викторины
* Next Quiz - Перейти к следующему вопросу
* Stop Quiz - Завершить текущую викторину

## Примечание
 Перед использованием не забудьте заменить TOKEN на актуальный токен вашего бота Telegram.
