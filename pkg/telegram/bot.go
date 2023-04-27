package telegram

import (
	"log"

	"github.com/ChingizAdamov/pocket_bot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
)
type Bot struct {
	bot *tgbotapi.BotAPI
	pocketClient *pocket.Client
	TokenRepositiry repository.TokenRepositore
	redirectURL string
}

func NewBot(bot *tgbotapi.BotAPI, pocketClient *pocket.Client, tr repository.TokenRepositore ,redirectURL string) *Bot {
	return &Bot{bot: bot, pocketClient: pocketClient, redirectURL: redirectURL, TokenRepositiry: tr}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}

	b.handleUpdates(updates)
	return nil	
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil { // If we got a message
			continue
	}
	if update.Message.IsCommand() {
		b.handleCommand(update.Message)
		continue
	}

	b.handleMessage(update.Message)
}
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u), nil

} 



