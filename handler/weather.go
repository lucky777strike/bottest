package handler

import (
	"errors"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lucky777strike/bottest/domain"
	"github.com/lucky777strike/tgmux"
)

func (h *Handler) GetWeather(c *tgmux.Ctx) {
	h.usecase.IncrementUserStatistics(h.ctx, c.Msg.From.ID)
	currentFunction := c.State.GetCurrentFunction()
	if currentFunction == "" {
		c.State.SetCurrentFunction("weather")

		cities := h.usecase.AviableCities()

		var sb strings.Builder
		for _, element := range cities {
			sb.WriteString(element)
			sb.WriteString("\n")
		}

		message := fmt.Sprintf("Привет %s.  Выбери город,\n%s ", c.Msg.From.FirstName, sb.String())
		reply := tgbotapi.NewMessage(c.Msg.Chat.ID, message)
		reply.ReplyToMessageID = c.Msg.MessageID
		_, err := c.Bot.Send(reply)
		if err != nil {
			h.logger.Errorf("Error sending message: %v\n", err)
		}
	}
	if currentFunction == "weather" {
		w, err := h.usecase.GetWeather(h.ctx, c.Msg.Text)
		if err != nil {
			if errors.Is(err, domain.ErrCityNotFound) {
				message := fmt.Sprintf("Прости %s, но я не знаю города %s \n Попробуй еще раз", c.Msg.From.FirstName, c.Msg.Text)
				reply := tgbotapi.NewMessage(c.Msg.Chat.ID, message)
				reply.ReplyToMessageID = c.Msg.MessageID
				_, err := c.Bot.Send(reply)
				if err != nil {
					h.logger.Errorf("Error sending message: %v\n", err)
				}
				return
			}
			c.SendErrorMessage(err)
			return
		}

		message := fmt.Sprintf(" \n Погода %+d %s, \nДата обновления %s",
			w.Temp, w.Condition, w.LastUpd)

		reply := tgbotapi.NewMessage(c.Msg.Chat.ID, message)
		reply.ReplyToMessageID = c.Msg.MessageID

		_, err = c.Bot.Send(reply)
		c.State.SetCurrentFunction("")
		if err != nil {
			h.logger.Errorf("Error sending message: %v\n", err)
		}
	}

}
