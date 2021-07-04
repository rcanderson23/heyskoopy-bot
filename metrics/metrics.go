package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	defaultLabels = []string{"guildID", "channelID"}

	MessagesReceived = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "discord_messages_received_total",
			Help: "Messages received from discord",
		}, append(defaultLabels, "userID"))

	ReactionsAdded = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "discord_reaction_added_total",
			Help: "Reactions added to messages",
		}, append(defaultLabels, "emojiID", "emojiName"))
)