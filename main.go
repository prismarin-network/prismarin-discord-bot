package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"primsarin-discord-bot/roles"
)

var token string

func initFlags() {
	flag.StringVar(&token, "token", "default", "Discord bot auth token")
	flag.Parse()
}

func main() {
	initFlags()

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}
	discord.Open()
	defer discord.Close()

	printBanner()
	ModifyLogger()

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
