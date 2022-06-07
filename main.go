package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {

	s := InitSession()
	SetRouter(s)
	Connect(s)
}


func SetRouter(s *discordgo.Session) {
	s.AddHandler(messageCreate)
//	s.AddHandler(messageDelete)
}
func InitSession() (*discordgo.Session) {
	token := ""

	// Create a new Discord session using the provided bot token.
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
	}
	
	return s
}
func Connect(s *discordgo.Session) {
	// In this example, we only care about receiving message events.
	s.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err := s.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	s.Close()
}
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

/*
	} else if m.Content == "pong" {  //tam eşleşme
*/

	if m.Author.Bot {
		return
	}

	if strings.Contains(m.Content, "ping") {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	} else if strings.Contains(m.Content, " koçum ") {
		s.ChannelMessageSend(m.ChannelID, "Buyur koçum!")
	} else if strings.Contains(m.Content, "amk") {
		s.ChannelMessageSend(m.ChannelID, "yine ne oluyor amk")
	} else if strings.Contains(m.Content, "günaydın") {
		s.ChannelMessageSend(m.ChannelID, "Günaydın")
	} else if strings.Contains(m.Content, "iyi geceler") {
		s.ChannelMessageSend(m.ChannelID, "iyi geceler")
	} else if strings.Contains(m.Content, "selam") {
		s.ChannelMessageSend(m.ChannelID, "Aleyküm Selam Hoşgeldin")


	}
}

/*func messageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {

	s.ChannelMessageSend(m.ChannelID, "https://stayhipp.com/wp-content/uploads/2019/02/you-better-watch.jpg")
}*/