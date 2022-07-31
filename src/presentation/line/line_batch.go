package line

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"household.api/src/presentation/line/controller"
)

type LineBatch struct {
	bot *linebot.Client
	weatherController controller.WeatherController
}

func NewLineBatch(bot *linebot.Client, weatherController controller.WeatherController)*LineBatch{
	batch := new(LineBatch)
	batch.bot = bot
	batch.weatherController = weatherController
	return batch
}

func (batch *LineBatch) GetWeather(e echo.Context)(err error){
	replayMessage, err := batch.weatherController.GetWeather()
	if err != nil{
		return
	}
	_, err = batch.bot.BroadcastMessage(linebot.NewTextMessage(replayMessage)).Do()
	if err != nil{
		fmt.Print(err)
	}
	return
}
