package heyskoopy

import (
	"context"
	"fmt"
	db2 "github.com/rcanderson23/heyskoopy-bot/db"
	"strings"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func (b *Bot) listCommand(s *discordgo.Session, input []string, m *discordgo.Message) {
	var (
		resp string
		err  error
	)

	name := strings.Join(input[3:], " ")

	switch input[2] {
	case "add":
		resp, err = b.listAdd(name, m.Author.ID)
		if err != nil {
			log.Errorf("Failed to add to list: %v", err)
		}
	case "delete":
		resp, err = b.listDelete(name)
		if err != nil {
			log.Errorf("Failed to delete from the list: %v", err)
		}
	case "print":
		resp, err = b.listPrint()
		if err != nil {
			log.Errorf("Failed to get the list: %v", err)
		}
	default:
		resp = listHelp()
	}

	_, err = s.ChannelMessageSend(m.ChannelID, resp)
	if err != nil {
		log.Errorf("failed to send message: %v", err)
	}
}

func (b *Bot) listAdd(li, author string) (string, error) {
	var resp string

	err := b.DB.AddListItem(context.TODO(), db2.ListItem{
		Name:    li,
		Creator: author,
	})
	if err != nil {
		return "Failed to persist to the list", err
	}

	resp = fmt.Sprintf("Successfully added %s to the list.", li)

	return resp, nil
}

func (b *Bot) listDelete(li string) (string, error) {
	var resp string

	count, err := b.DB.DeleteListItem(context.TODO(), db2.ListItem{
		Name: li,
	})
	if err != nil {
		return "Failed to delete list item", err
	}

	if count == 0 {
		resp := fmt.Sprintf("%s is not on the list.", li)
		return resp, nil
	}

	resp = fmt.Sprintf("Successfully removed %s from the list.", li)

	return resp, nil
}

func (b *Bot) listPrint() (string, error) {
	var resp string

	list, err := b.DB.GetList(context.TODO())
	if err != nil {
		return "Failed to get the list", err
	}

	if len(list) == 0 {
		return "List is empty", nil
	}

	resp += "```\nThe List\n"

	for i, item := range list {
		resp += fmt.Sprintf("%d: %s\n", i, item.Name)
	}

	resp += "```"

	return resp, nil
}

func listHelp() string {
	return ">>> __**List Commands:**__\n" +
		"**Add:** `!hs list add [name]`\n" +
		"**Delete:** `!hs list delete [name]`\n" +
		"**Print:** `!hs list print`\n" +
		"**Help:** `!hs list`"
}
