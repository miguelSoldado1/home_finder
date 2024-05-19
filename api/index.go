package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	bot "github.com/miguelSoldado1/home_finder_bot/pkg/discord"
	imovirtual "github.com/miguelSoldado1/home_finder_bot/pkg/imovirtual"
)

// the cron job runs every hour
const offset = time.Duration(-1) * time.Hour

func main() int {
	// lisboa_under_800
	ads := imovirtual.Search("800", offset)
	numOfAds := len(ads)

	// only run the bot if there are new ads
	if numOfAds > 0 {
		bot.Run()
		for _, msg := range ads {
			bot.SendMessage("1241143437533777931", msg)
		}
		bot.Close()
	}

	return numOfAds
}

func Handler(w http.ResponseWriter, r *http.Request) {
	numOfMessages := main()
	fmt.Println(numOfMessages, "new ads found!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = fmt.Sprintf("%d new ads found!", numOfMessages)
	resp["language"] = "go"
	resp["cloud"] = "Hosted on Vercel! ▲"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)
}
