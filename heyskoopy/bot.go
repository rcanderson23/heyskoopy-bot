package heyskoopy

import (
	"github.com/bwmarrin/discordgo"
	db2 "github.com/rcanderson23/heyskoopy-bot/db"
	log "github.com/sirupsen/logrus"
	"regexp"
)

// Bot is the base struct to start the HeySkoopy bot
type Bot struct {
	DB               db2.DB
	BotCommandString *regexp.Regexp
	session          *discordgo.Session
}

// Run starts the bot with the provided authKey
func (b *Bot) Run(authKey string) {
	session, err := discordgo.New("Bot " + authKey)
	if err != nil {
		log.Fatalf("Error creating discord session: %v", err)
	}

	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	session.AddHandler(b.messageCreate)
	session.AddHandler(b.messageUpdate)
	session.AddHandler(b.messageReactionAdd)

	err = session.Open()
	if err != nil {
		log.Fatalf("Failed to open session: %v", err)
	}

	log.Info("HeySkoopy Bot Started")
}

// Exit is used to gracefully terminate the bot
func (b *Bot) Exit() {
	_ = b.session.Close()
}
