package heyskoopy

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

const (
	MinCommandLength = 3
)

func (b *Bot) commandRouter(s *discordgo.Session, m *discordgo.Message) {
	if !b.BotCommandString.MatchString(m.Content) {
		return
	}

	input := strings.Split(m.Content, " ")

	if len(input) < MinCommandLength {
		_, err := s.ChannelMessageSend(m.ChannelID, commandHelp())
		if err != nil {
			log.Errorf("failed to print command help: %v", err)
		}

		return
	}

	switch input[1] {
	case "list":
		b.listCommand(s, input, m)
	default:
		_, err := s.ChannelMessageSend(m.ChannelID, "Not a valid command")
		if err != nil {
			log.Errorf("Failed to send message to channel: %s", m.ChannelID)
		}
	}
}

func commandHelp() string {
	return ">>> __**Commands:**__\n" +
		"**List**: `!hs list <add|delete|print|help> [name]`"
}
