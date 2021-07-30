package heyskoopy

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

const (
	// MinCommandLength is the minimum number of strings that should be present to use commandRouter
	MinCommandLength = 2
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

	log.Infof("Bot action invoked with command %s by %s", input[1], m.Author.Username)

	var (
		resp string
		err  error
	)

	switch input[1] {
	case "list":
		resp, err = b.listCommand(input, m)
		if err != nil {
			log.Errorf("List command failed: %v", err)
		}
	default:
		resp = "Not a valid command"
	}

	_, err = s.ChannelMessageSend(m.ChannelID, resp)
	if err != nil {
		log.Errorf("Failed to send message to channel %s: %v", m.ChannelID, err)
	}
}

func commandHelp() string {
	return ">>> __**Commands:**__\n" +
		"**List**: `!hs list <[add]|[delete]|[help]> [name]`"
}
