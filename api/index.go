package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"

	bot "example.com/home_finder_bot/Bot"
	imovirtual "example.com/home_finder_bot/Imovirtual"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Api triggered")

	bot_token := os.Getenv("BOT_TOKEN")

	if bot_token == "" {
		log.Fatal("BOT_TOKEN is not set")
	}

	bot.BotToken = bot_token
	bot.Run()

	// lisboa_under_800
	messages := imovirtual.Search("800", "[lisboa/lisboa/santa-clara,lisboa/lisboa/benfica,lisboa/lisboa/avenidas-novas,lisboa/lisboa/arroios,lisboa/lisboa/alvalade,lisboa/lisboa/penha-de-franca,lisboa/lisboa/olivais,lisboa/lisboa/lumiar,lisboa/lisboa/parque-das-nacoes,lisboa/lisboa/areeiro]")
	for _, msg := range messages {
		bot.SendMessage("1241143437533777931", msg)
	}

	bot.Close()
}
