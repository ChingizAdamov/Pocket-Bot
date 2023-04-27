package main

import (
	"log"

	"github.com/ChingizAdamov/pocket_bot/pkg/repository"
	"github.com/ChingizAdamov/pocket_bot/pkg/repository/tgboltdb"
	"github.com/ChingizAdamov/pocket_bot/pkg/server"
	"github.com/ChingizAdamov/pocket_bot/pkg/telegram"
	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6135967932:AAHH25mabWetO46hMEXWeL7YMdOwKyP0qkU")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient("107062-54d5cc8eeed3f9c0825add3")
	if err != nil {
		log.Fatal(err)
	}

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	TokenRepositiry := tgboltdb.NewTokenRepositiry(db)

	telegramBot := telegram.NewBot(bot, pocketClient,TokenRepositiry, "http://localhost/")

	authorizationServer := server.NewAuthServer(pocketClient, TokenRepositiry, "https://t.me/mvp_pocket_bot")

	go func ()  {
		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := authorizationServer.Start(); err != nil {
		log.Fatal(err)
	}
}

func initDB() (*bolt.DB, error) {
	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.AccessTokens))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(repository.RequestTokens))
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return db, nil
}