package handler

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lucky777strike/tgmux"
)

func (h *Handler) stat(c *tgmux.Ctx) {
	h.usecase.Stat.IncrementUserStatistics(h.ctx, c.Msg.From.ID)
	stat, err := h.usecase.Stat.GetUserStatistics(h.ctx, c.Msg.From.ID)
	if err != nil {
		c.SendErrorMessage(err)
		return
	}
	message := fmt.Sprintf("Привет, %s! \n Первый запрос %s \n Последний вопрос %s \n Всего запросов %d", c.Msg.From.FirstName,
		stat.FirstRequestTime, stat.LastRequestTime, stat.TotalRequests)

	//fmt.Println(h.usecase.GetUserStatistics(h.ctx, c.Msg.From.ID))
	reply := tgbotapi.NewMessage(c.Msg.Chat.ID, message)
	reply.ReplyToMessageID = c.Msg.MessageID

	_, err = c.Bot.Send(reply)
	if err != nil {
		h.logger.Errorf("Error sending message: %v\n", err)
	}
}

func (h *Handler) reset(c *tgmux.Ctx) {
	h.usecase.Stat.ResetUserStatistics(h.ctx, c.Msg.From.ID)
	stat, err := h.usecase.Stat.GetUserStatistics(h.ctx, c.Msg.From.ID)

	if err != nil {
		c.SendErrorMessage(err)
		return
	}
	message := fmt.Sprintf("Привет, %s, Статистика сброшена, ! \n Первый запрос %s \n Последний вопрос %s \n Всего запросов %d", c.Msg.From.FirstName,
		stat.FirstRequestTime, stat.LastRequestTime, stat.TotalRequests)

	//fmt.Println(h.usecase.GetUserStatistics(h.ctx, c.Msg.From.ID))
	reply := tgbotapi.NewMessage(c.Msg.Chat.ID, message)
	reply.ReplyToMessageID = c.Msg.MessageID

	_, err = c.Bot.Send(reply)
	if err != nil {
		h.logger.Errorf("Error sending message: %v\n", err)
	}
}
