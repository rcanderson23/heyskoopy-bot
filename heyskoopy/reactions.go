package heyskoopy

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

func (b *Bot) heyReaction(s *discordgo.Session, m *discordgo.Message) {
	hey, err := regexp.MatchString(`^(hey)\b`, strings.ToLower(m.Content))
	if err != nil {
		log.Errorf("Regex error: %v", err)
	}

	if hey {
		err := s.MessageReactionAdd(m.ChannelID, m.ID, "pepecorner:730969584177381426")
		if err != nil {
			log.Errorf("%v", err)
		}
	}
}

func (b *Bot) kubernetesReaction(s *discordgo.Session, m *discordgo.Message) {
	k8s, err := regexp.MatchString(`(k8s|K8s|kubernetes|Kubernetes)\b`, m.Content)
	if err != nil {
		log.Errorf("Regex error: %v", err)
	}

	if k8s {
		err := s.MessageReactionAdd(m.ChannelID, m.ID, "kubernetes:712304909177061457")
		if err != nil {
			log.Errorf("%v", err)
		}
	}
}
