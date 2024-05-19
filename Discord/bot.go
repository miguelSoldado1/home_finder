package bot

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
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

func init() {
	err := godotenv.Load()
	checkNilErr(err)

	// Get the bot token from the environment variables
	BotToken = os.Getenv("BOT_TOKEN")
	if BotToken == "" {
		log.Fatal("Bot token not found in environment variables")
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
