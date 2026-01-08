package main

import (
	"github.com/mfinikov/tg-assistant-bot/internal/config"
	"github.com/mfinikov/tg-assistant-bot/internal/telegram"
)

func main() {
	config.EnvLoad()
	telegram.InitBot()
}
