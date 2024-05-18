package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	bot "example.com/home_finder_bot/Bot"
	imovirtual "example.com/home_finder_bot/Imovirtual"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	main()
	fmt.Println("Api triggered")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Hello World from Go! 👋"
	resp["language"] = "go"
	resp["cloud"] = "Hosted on Vercel! ▲"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)
}

func main() {
	bot_token := os.Getenv("BOT_TOKEN")

	if bot_token == "" {
		log.Fatal("BOT_TOKEN is not set")
	}

	bot.BotToken = bot_token
	bot.Run()

	// lisboa_under_800
	messages := imovirtual.Search("800")
	for _, msg := range messages {
		bot.SendMessage("1241143437533777931", msg)
	}

	bot.Close()
}
