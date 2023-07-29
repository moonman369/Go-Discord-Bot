package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/moonman369/Go-Discord-Bot/config"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is up and running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	allowedGreetings := []string{"hello", "Hello", "hi", "Hi", "hey", "Hey", "Yo", "yo", "Wassup", "wassup", "ssup", "Ssup"}

	for _, greeting := range allowedGreetings {
		if m.Content == greeting {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v!!  I am PingBot-v0. How can I help you!!", greeting))
			return
		}
	}

	s.ChannelMessageSend(m.ChannelID, "Sorry! Didn't quite get that...")

}
