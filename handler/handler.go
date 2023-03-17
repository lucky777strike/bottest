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
	cancel  context.CancelFunc
	token   string
	usecase domain.Usecase
}

func NewHandler(ctx context.Context, cancel context.CancelFunc, usecase domain.Usecase, token string) *Handler {
	return &Handler{ctx: ctx,
		usecase: usecase,
		token:   token}
}

func (h *Handler) Start() error {
	bot, err := tgmux.NewHandlerWithContext(h.ctx, h.cancel, h.token)
	if err != nil {
		return err
	}
	bot.HandleCmd("/sum", startCommand)
	bot.HandleState("sum", startCommand)
	go func() {
		bot.Start()
	}()
	return nil
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
