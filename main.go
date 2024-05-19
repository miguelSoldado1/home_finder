package main

import (
	"time"

	bot "github.com/miguelSoldado1/home_finder_bot/pkg/discord"
	imovirtual "github.com/miguelSoldado1/home_finder_bot/pkg/imovirtual"
)

// the cron job runs every hour
const offset = time.Duration(-1) * time.Hour

func main() {
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
}
