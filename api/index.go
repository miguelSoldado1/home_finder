package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	bot "example.com/home_finder_bot/Discord"
	imovirtual "example.com/home_finder_bot/Imovirtual"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	numOfMessages := main()
	fmt.Println(numOfMessages, "new ads found!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Success!"
	resp["new_ads"] = fmt.Sprintf("%d", numOfMessages)
	resp["language"] = "go"
	resp["cloud"] = "Hosted on Vercel! ▲"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)
}

func main() int {
	// lisboa_under_800
	ads := imovirtual.Search("800")
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
