package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	BotToken string
	discord  *discordgo.Session
)

func checkNilErr(e error) {
	if e != nil {
		log.Fatal("Error message: ", e)
	}
}

func Run() {
	var err error
	// create a session
	discord, err = discordgo.New("Bot " + BotToken)
	checkNilErr(err)

	// open session
	err = discord.Open()
	checkNilErr(err)

}

func SendMessage(channelID, message string) {
	_, err := discord.ChannelMessageSend(channelID, message)
	checkNilErr(err)
}

func Close() {
	defer discord.Close()
}
