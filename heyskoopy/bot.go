package heyskoopy

import (
	"github.com/bwmarrin/discordgo"
	db2 "github.com/rcanderson23/heyskoopy-bot/db"
	log "github.com/sirupsen/logrus"
	"regexp"
)

type Bot struct {
	DB               db2.DB
	BotCommandString *regexp.Regexp
	session          *discordgo.Session
}

func (b *Bot) Run(authKey string) {
	session, err := discordgo.New("Bot " + authKey)
	if err != nil {
		log.Fatalf("Error creating discord session: %v", err)
	}

	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	session.AddHandler(b.MessageCreate)
	session.AddHandler(b.MessageUpdate)
	session.AddHandler(b.MessageReactionAdd)

	err = session.Open()
	if err != nil {
		log.Fatalf("Failed to open session: %v", err)
	}

	log.Info("HeySkoopy Bot Started")
}

func (b *Bot) Exit() {
	_ = b.session.Close()
}
