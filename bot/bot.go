package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/moonman369/Go-Discord-Bot/config"
	"github.com/moonman369/Go-Discord-Bot/gpt"
)

var BotID string

// var goBot *discordgo.Session

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

	go func() {
		goBot.AddHandler(messageHandler)
	}()

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

	if []byte(m.Content)[0] != []byte("!")[0] {
		return
	}

	if len(m.Content) <= 0 {
		return
	}

	// allowedGreetings := []string{"hello", "Hello", "hi", "Hi", "hey", "Hey", "Yo", "yo", "Wassup", "wassup", "ssup", "Ssup"}

	// for _, greeting := range allowedGreetings {
	// 	if m.Content == greeting {
	// 		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v!!  I am PingBot-v0. How can I help you!!", greeting))
	// 		return
	// 	}
	// }
	go func() {
		s.ChannelTyping(m.ChannelID)
	}()
	Resp := gpt.SendPrompt(fmt.Sprintf("Refer to your self as Ping-Bot-v0 whenever you are asked to identify yourself. %v", m.Content))

	if len(Resp.Choices) < 1 {
		s.ChannelMessageSend(m.ChannelID, "Could not create suitable response. Please try again.")
		return
	}

	go func() {
		s.ChannelTyping(m.ChannelID)
	}()

	fmt.Println(Resp.Choices[0].Message.Content)

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(`%v`, Resp.Choices[0].Message.Content))

}
