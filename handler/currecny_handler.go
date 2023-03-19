package handler

import (
	"errors"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lucky777strike/bottest/domain"
	"github.com/lucky777strike/tgmux"
)

func (h *Handler) GetCurrency(c *tgmux.Ctx) {
	h.usecase.Stat.IncrementUserStatistics(h.ctx, c.Msg.From.ID)
	currentFunction := c.State.GetCurrentFunction()
	if currentFunction == "" {
		c.State.SetCurrentFunction("currency")

		aviable := h.usecase.Currency.AvailableCurrencies()

		var sb strings.Builder
		for _, element := range aviable {
			sb.WriteString(element)
			sb.WriteString("\n")
		}

		message := fmt.Sprintf("Привет %s.  Выбери Валюту,\n%s ", c.Msg.From.FirstName, sb.String())
		reply := tgbotapi.NewMessage(c.Msg.Chat.ID, message)
		reply.ReplyToMessageID = c.Msg.MessageID
		_, err := c.Bot.Send(reply)
		if err != nil {
			h.logger.Errorf("Error sending message: %v\n", err)
		}
	}
	if currentFunction == "currency" {
		w, err := h.usecase.Currency.GetCurrency(h.ctx, c.Msg.Text)
		if err != nil {
			if errors.Is(err, domain.ErrCurrencyUnknown) {
				message := fmt.Sprintf("Прости %s, но я не знаю валюты %s \n Попробуй еще раз", c.Msg.From.FirstName, c.Msg.Text)
				reply := tgbotapi.NewMessage(c.Msg.Chat.ID, message)
				reply.ReplyToMessageID = c.Msg.MessageID
				_, err := c.Bot.Send(reply)
				if err != nil {
					h.logger.Errorf("Error sending message: %v\n", err)
					c.SendErrorMessage(fmt.Errorf("error"))
				}
				return
			}
			h.logger.Errorf("Error in handler: %v\n", err)
			return
		}

		message := fmt.Sprintf(" \n Курс %s %f рублей, \nДата обновления %s",
			w.Name, w.Value, w.LastUpd)

		reply := tgbotapi.NewMessage(c.Msg.Chat.ID, message)
		reply.ReplyToMessageID = c.Msg.MessageID

		_, err = c.Bot.Send(reply)
		c.State.SetCurrentFunction("")
		if err != nil {
			h.logger.Errorf("Error in handler: %v\n", err)
		}
	}

}
