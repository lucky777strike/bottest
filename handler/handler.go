package handler

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

func (h *Handler) startCommand(c *tgmux.Ctx) {
	h.usecase.Stat.IncrementUserStatistics(h.ctx, c.Msg.From.ID)
	welcomeMessage := fmt.Sprintf(`Привет, %s! Доступные команды:
	/weather -- прогноз погоды
	/currency -- курсы валют
	/stat -- статистика запросов
	/reset -- сброс статистики`, c.Msg.From.FirstName)

	reply := tgbotapi.NewMessage(c.Msg.Chat.ID, welcomeMessage)
	reply.ReplyToMessageID = c.Msg.MessageID

	_, err := c.Bot.Send(reply)
	if err != nil {
		h.logger.Errorf("Error sending message: %v\n", err)
	}
}
