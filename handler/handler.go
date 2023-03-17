package handler

import (
	"context"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lucky777strike/bottest/domain"
	"github.com/lucky777strike/tgmux"
)

type Handler struct {
	ctx     context.Context
	token   string
	usecase domain.Usecase
}

func NewHandler(ctx context.Context, usecase domain.Usecase, token string) *Handler {
	return &Handler{ctx: ctx,
		usecase: usecase,
		token:   token}
}

func startCommand(c *tgmux.Ctx) {
	welcomeMessage := fmt.Sprintf("Hello, %s! Welcome to the example bot.", c.Msg.From.FirstName)

	reply := tgbotapi.NewMessage(c.Msg.Chat.ID, welcomeMessage)
	reply.ReplyToMessageID = c.Msg.MessageID

	_, err := c.Bot.Send(reply)
	if err != nil {
		log.Printf("Error sending message: %v\n", err)
	}
}
