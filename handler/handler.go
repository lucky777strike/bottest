package handler

import (
	"context"

	"github.com/lucky777strike/bottest/domain"
	"github.com/lucky777strike/tgmux"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	ctx     context.Context
	token   string
	usecase domain.Service
	logger  *logrus.Logger
}

func NewHandler(ctx context.Context, logger *logrus.Logger, usecase *domain.Service, token string) *Handler {
	return &Handler{
		ctx:     ctx,
		token:   token,
		usecase: *usecase,
		logger:  logger,
	}
}

func (h *Handler) Start() error {
	bot, err := tgmux.NewHandlerWithContext(h.ctx, h.token)
	bot.SetLogger(h.logger)
	if err != nil {
		return err
	}
	bot.HandleCmd("/start", h.startCommand)

	bot.HandleCmd("/stat", h.stat)
	bot.HandleCmd("/reset", h.reset)

	bot.HandleCmd("/weather", h.GetWeather)
	bot.HandleState("weather", h.GetWeather)
	bot.HandleCmd("/currency", h.GetCurrency)
	bot.HandleState("currency", h.GetCurrency)

	h.logger.Info("Starting bot...")
	go func() {
		bot.Start()
	}()
	return nil
}
