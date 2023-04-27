package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart = "start"
	replyStart = "Привет! Для того чтобы сохранить ссылку в Pocket аккаунте тебе необходимо дать мне доступ, бро! Для этого переходиим по ссылке: \n%s"
	replyToAuth = "Ты уже авторезирован, тебе лишь надо скинуть мне ссылку :)"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {

	switch message.Command() {
	case commandStart:
		 return b.handleStartCommand(message)
	default:
		return b.unknownCommands(message)
	}

}

func (b *Bot) handleMessage(message *tgbotapi.Message)  {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	b.bot.Send(msg)
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	_, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return b.initAuthProcess(message)
	}
	
	msg := tgbotapi.NewMessage(message.Chat.ID,replyToAuth)
	b.bot.Send(msg)
	return err
}

func (b *Bot) unknownCommands(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды :(")
	_, err :=b.bot.Send(msg)
		return err
}



