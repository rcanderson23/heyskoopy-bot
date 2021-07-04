package main

import (
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rcanderson23/heyskoopy-bot/db"
	"github.com/rcanderson23/heyskoopy-bot/heyskoopy"
	log "github.com/sirupsen/logrus"
)

const (
	// DiscordAuthKey is the API key used by the bot to authenticate with Discord
	DiscordAuthKey        = "DISCORD_AUTH_KEY"

	// MongoConnectionString is the connection string used to connect to MongoDB
	MongoConnectionString = "MONGO_CONNECTION_STRING"

	// MongoDBName is the name of the DB
	MongoDBName           = "MONGO_DB_NAME"
)

func main() {
	authKey := os.Getenv(DiscordAuthKey)
	connString := os.Getenv(MongoConnectionString)
	dbName := os.Getenv(MongoDBName)

	mongo, err := db.NewMongo(connString, dbName)
	if err != nil {
		log.Fatalf("failed to create mongo client: %v", err)
	}

	command := regexp.MustCompile(`^!hs?\w`)

	bot := heyskoopy.Bot{
		DB:               mongo,
		BotCommandString: command,
	}
	bot.Run(authKey)


	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(":9090", nil))
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	bot.Exit()
}
