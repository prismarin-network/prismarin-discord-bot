package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"primsarin-discord-bot/roles"
	"time"
)

var token = "OTY3NjE2MDI1Mzc1OTM2NTgy.YmS4pQ.IvpXoMNsi1dlCadVlPsw_jeei0k"

type CustomWriter struct {
	File *os.File
}

func (writer *CustomWriter) Write(data []byte) (int, error) {
	message := string(data)
	fmt.Printf(message)
	writer.File.WriteString(message)
	return len(data), nil
}

func modifyLogger() {
	file, _ := os.OpenFile(fmt.Sprintf("%s.log", time.Now().Format("2006-01-02")), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	w := CustomWriter{
		File: file,
	}
	defer file.Close()
	log.SetOutput(&w)
}

func main() {
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}
	discord.Open()
	defer discord.Close()

	printBanner()
	modifyLogger()

	log.Println("Bot is now running")
	roles.InitReactionRoles(discord)

	discord.AddHandler(roles.MessageReactAdd)
	discord.AddHandler(roles.MessageReactRemove)

	fmt.Scanln()
}

func printBanner() {
	data, _ := os.ReadFile("banner.txt")
	text := string(data)

	fmt.Println(text)
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

}
