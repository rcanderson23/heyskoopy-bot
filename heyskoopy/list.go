package heyskoopy

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	db2 "github.com/rcanderson23/heyskoopy-bot/db"
	log "github.com/sirupsen/logrus"
	"strings"
)

// listCommand adds, deletes, or prints the list. The default action should be to print the list.
func (b *Bot) listCommand(input []string, m *discordgo.Message) (string, error) {
	if len(input) == 2 {
		log.Infof("List command print invoked by %s", m.Author.Username)
		return b.listPrint()
	}

	if len(input) == 3 {
		if input[2] == "print" {
			log.Infof("List command print invoked by %s", m.Author.Username)
			return b.listPrint()
		}

		log.Infof("List command help invoked by %s", m.Author.Username)

		return listHelp(), nil
	}

	command := input[2]
	name := strings.Join(input[3:], "")

	log.Infof("List command with command %s by %s", command, m.Author.Username)

	switch command {
	case "add":
		return b.listAdd(name, m.Author.ID)
	case "delete":
		return b.listDelete(name)
	default:
		return listHelp(), nil
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
		"**Print:** `!hs list`\n" +
		"**Help:** `!hs list help`"
}
