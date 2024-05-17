package main

import (
	"log"
	"os"

	bot "example.com/home_finder_bot/Bot"
	imovirtual "example.com/home_finder_bot/Imovirtual"
	"github.com/joho/godotenv"
)

const priceMax = "800"
const locations = "[lisboa/lisboa/santa-clara,lisboa/lisboa/benfica,lisboa/lisboa/avenidas-novas,lisboa/lisboa/arroios,lisboa/lisboa/alvalade,lisboa/lisboa/penha-de-franca,lisboa/lisboa/olivais,lisboa/lisboa/lumiar,lisboa/lisboa/parque-das-nacoes,lisboa/lisboa/areeiro]"

func checkNilErr(e error) {
	if e != nil {
		log.Fatal("Error message: ", e)
	}
}

func main() {
	err := godotenv.Load()
	checkNilErr(err)

	bot.BotToken = os.Getenv("BOT_TOKEN")
	bot.Run()

	messages := imovirtual.Search(priceMax, locations)
	for _, msg := range messages {
		bot.SendMessage("1241108323655356497", msg)
	}

	bot.Close()
}
