package line

import (
	"github.com/labstack/gommon/log"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"os"
)

// LineInit LineBotの作成のみを行う関数
func LineInit() (bot *linebot.Client, err error) {
	bot, err = linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"))
	if err != nil{
		log.Fatal(err)
	}
	return
}
