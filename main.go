package main

import (
	"log"
	"os"

	bot "example.com/home_finder_bot/Bot"
	imovirtual "example.com/home_finder_bot/Imovirtual"
	"github.com/joho/godotenv"
)

func checkNilErr(e error) {
	if e != nil {
		log.Fatal("Error message: ", e)
	}
}

func main() {
	err := godotenv.Load(".env")
	checkNilErr(err)

	bot.BotToken = os.Getenv("BOT_TOKEN")
	bot.Run()

	messages := imovirtual.Search()
	for _, msg := range messages {
		bot.SendMessage("1241108323655356497", msg)
	}

	bot.Close()
}
