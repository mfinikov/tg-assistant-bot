package telegram

import (
	"fmt"
	"time"

	"github.com/mfinikov/tg-assistant-bot/internal/config"
	tele "gopkg.in/telebot.v4"
)

// InitBot is a function to initialize tg bot
func InitBot() {

	defer fmt.Println("Bot stopped")

	pref := tele.Settings{
		Token:  config.BotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		panic(err)
	}

	injectCommandHandlers(b)

	fmt.Println("Bot started")
	b.Start()

}

func injectCommandHandlers(b *tele.Bot) {
	b.Handle("/start", func(c tele.Context) error {
		return c.Send(`👋 Hey! I am bot for Telegram Business.

I will notify you of any changes or deletions of messages in private chats.

Instructions:
1. Open the settings.
2. Telegram for business -> Chatbots
3. Choose me as your chatbot.

Bot created by @mfinikov`)
	})
}
