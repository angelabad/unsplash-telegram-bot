package main

import (
	"github.com/Sirupsen/logrus"

	"gopkg.in/telegram-bot-api.v4"
)

func getBot() (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(telegramID)
	if err != nil {
		return bot, err
	}

	return bot, nil
}

func sendPhoto(message *tgbotapi.Message, logger *logrus.Logger, photo Photo) error {
	bot, err := getBot()
	if err != nil {
		return err
	}

	logger.Debug("Sending photo from ", bot.Self.UserName)

	image, err := photo.getSmallImage()
	if err != nil {
		return err
	}
	b := tgbotapi.FileBytes{Name: "photo.jpg", Bytes: image}
	msgPhoto := tgbotapi.NewPhotoUpload(message.Chat.ID, b)
	originalBtn := tgbotapi.NewInlineKeyboardButtonURL("Download original", photo.Urls.Raw)
	userBtn := tgbotapi.NewInlineKeyboardButtonURL("by "+photo.User.Name, photo.User.Links.HTML)
	row1 := []tgbotapi.InlineKeyboardButton{
		originalBtn,
	}
	row2 := []tgbotapi.InlineKeyboardButton{
		userBtn,
	}
	inlineKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup(row1, row2)
	msgPhoto.ReplyMarkup = &inlineKeyboardMarkup
	_, err = bot.Send(msgPhoto)
	if err != nil {
		return err
	}

	return nil
}

func sendTxt(message *tgbotapi.Message, logger *logrus.Logger, txt string) error {
	bot, err := getBot()
	if err != nil {
		logger.Panic("ERROR CON BOT: ", err)
	}
	msgTxt := tgbotapi.NewMessage(message.Chat.ID, txt)
	msgTxt.DisableWebPagePreview = true
	_, err = bot.Send(msgTxt)
	if err != nil {
		return err
	}

	return nil
}

func sendHTML(message *tgbotapi.Message, logger *logrus.Logger, txt string) error {
	bot, err := getBot()
	if err != nil {
		logger.Panic("ERROR CON BOT: ", err)
	}
	msgTxt := tgbotapi.NewMessage(message.Chat.ID, txt)
	msgTxt.ParseMode = "HTML"
	msgTxt.DisableWebPagePreview = true
	_, err = bot.Send(msgTxt)
	if err != nil {
		return err
	}

	return nil
}
