package heyskoopy

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func (b *Bot) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	b.heyReaction(s, m.Message)
	b.kubernetesReaction(s, m.Message)
	b.botMention(s, m.Message)
	b.commandRouter(s, m.Message)
}

func (b *Bot) MessageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {
	b.heyReaction(s, m.Message)
	b.kubernetesReaction(s, m.Message)
}

func (b *Bot) MessageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.UserID == s.State.User.ID {
		return
	}

	err := s.MessageReactionAdd(r.ChannelID, r.MessageID, fmt.Sprintf("%s:%s", r.Emoji.Name, r.Emoji.ID))
	if err != nil {
		log.Errorf("failed to add emoji: %v", err)
	}

	log.Infof("emoji %s added to %s", r.Emoji.Name, r.MessageID)
}

func (b *Bot) botMention(s *discordgo.Session, m *discordgo.Message) {
	if strings.Contains(m.Content, "<@!"+s.State.User.ID+">") {
		c, err := s.Channel(m.ChannelID)
		if err != nil {
			log.Errorf("Failed to get channel: %v", err)
			return
		}

		log.Infof("SkoopyBot mentioned in %s by %s", c.Name, m.Author.String())

		resp := fmt.Sprintf("Fuck you %s", m.Author.Mention())

		_, err = s.ChannelMessageSend(m.ChannelID, resp)
		if err != nil {
			log.Errorf("Failed to send message: %v", err)
		}
	}
}
