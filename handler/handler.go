package handler

import (
	"context"

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
	bot.HandleCmd("/start", h.startCommand)
	bot.HandleState("/start", h.startCommand)

	bot.HandleCmd("/stat", h.stat)
	bot.HandleCmd("/reset", h.reset)
	go func() {
		bot.Start()
	}()
	return nil
}
