package telegram

import (
	"context"
	"fmt"

	"github.com/ChingizAdamov/pocket_bot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
func (b *Bot) initAuthProcess(message *tgbotapi.Message) error {
	authLink, err := b.generateAuthLink(message.Chat.ID)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(message.Chat.ID,
	fmt.Sprintf(replyStart, authLink))

		 _, err = b.bot.Send(msg)
			return err
}

func(b *Bot) getAccessToken(chatID int64) (string, error) {
	return b.TokenRepositiry.Get(chatID, repository.AccessTokens)
}

func (b *Bot) generateAuthLink(chatID int64) (string, error) {
	redirectURL := b.generateRedirectURL(chatID)

	requestToken, err := b.pocketClient.GetRequestToken(context.Background(), redirectURL)
	if err != nil {
		return "", err
	}

	if err := b.TokenRepositiry.Save(chatID, requestToken, repository.RequestTokens); err != nil {
		return "", err
	}
	
	return b.pocketClient.GetAuthorizationURL(requestToken, b.redirectURL)
}

func (b *Bot) generateRedirectURL(chatID int64) string {
	return fmt.Sprintf("%s?chat_id=%d", b.redirectURL,chatID)
}