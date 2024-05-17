package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"

	bot "example.com/home_finder_bot/Bot"
	imovirtual "example.com/home_finder_bot/Imovirtual"
)

const priceMax = "800"
const locations = "[lisboa/lisboa/santa-clara,lisboa/lisboa/benfica,lisboa/lisboa/avenidas-novas,lisboa/lisboa/arroios,lisboa/lisboa/alvalade,lisboa/lisboa/penha-de-franca,lisboa/lisboa/olivais,lisboa/lisboa/lumiar,lisboa/lisboa/parque-das-nacoes,lisboa/lisboa/areeiro]"

func Handler(w http.ResponseWriter, r *http.Request) {
	bot_token := os.Getenv("BOT_TOKEN")

	if bot_token == "" {
		log.Fatal("BOT_TOKEN is not set")
	}

	bot.BotToken = bot_token
	bot.Run()

	messages := imovirtual.Search(priceMax, locations)
	for _, msg := range messages {
		bot.SendMessage("1241108323655356497", msg)
	}

	bot.Close()

	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
