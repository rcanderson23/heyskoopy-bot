package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	db2 "github.com/rcanderson23/heyskoopy-bot/db"
	"github.com/rcanderson23/heyskoopy-bot/heyskoopy"
	log "github.com/sirupsen/logrus"
	"net/http"
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

	mongo, err := db2.NewMongo(connString, dbName, collections)
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
