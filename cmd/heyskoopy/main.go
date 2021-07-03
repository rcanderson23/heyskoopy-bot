package main

import (
	"github.com/rcanderson23/heyskoopy-bot/bot/db"
	hs "github.com/rcanderson23/heyskoopy-bot/bot/heyskoopy"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
)

const (
	DiscordAuthKey        = "DISCORD_AUTH_KEY"
	MongoConnectionString = "MONGO_CONNECTION_STRING"
	MongoDBName           = "MONGO_DB_NAME"
	MongoCollections      = "MONGO_COLLECTIONS"
)

func main() {
	authKey := os.Getenv(DiscordAuthKey)
	connString := os.Getenv(MongoConnectionString)
	dbName := os.Getenv(MongoDBName)
	collections := strings.Split(os.Getenv(MongoCollections), ",")

	mongo, err := db.NewMongo(connString, dbName, collections)
	if err != nil {
		log.Fatalf("failed to create mongo client: %v", err)
	}

	command := regexp.MustCompile(`^!hs?\w`)

	bot := hs.Bot{
		DB:               mongo,
		BotCommandString: command,
	}
	bot.Run(authKey)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	bot.Exit()
}
