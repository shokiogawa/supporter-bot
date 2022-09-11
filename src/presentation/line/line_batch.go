package line

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"household.api/src/presentation/line/controller"
)

type LineBatch struct {
	bot               *linebot.Client
	weatherController controller.WeatherController
	costController    controller.CostController
}

func NewLineBatch(bot *linebot.Client, weatherController controller.WeatherController, costController controller.CostController) *LineBatch {
	batch := new(LineBatch)
	batch.bot = bot
	batch.weatherController = weatherController
	batch.costController = costController
	return batch
}

func (batch *LineBatch) GetWeather(e echo.Context) (err error) {
	replayMessage, err := batch.weatherController.GetWeather()
	if err != nil {
		return
	}
	_, err = batch.bot.BroadcastMessage(linebot.NewTextMessage(replayMessage)).Do()
	if err != nil {
		fmt.Print(err)
	}
	return
}

func (batch *LineBatch) GetOutComePerMonth(e echo.Context) (err error) {
	replyMessages, err := batch.costController.CostPerMonthList()
	if err != nil {
		return
	}
	for key, value := range replyMessages {
		fmt.Println(value)
		if _, err := batch.bot.PushMessage(key, linebot.NewTextMessage(value)).Do(); err != nil {
		}
	}
	return
}

//複数のユーザーにメッセージを送信
//userIDs := []string{ ... }
//bot, err := linebot.New(<channel secret>, <channel token>)
//if err != nil {
//...
//}
//if _, err := bot.Multicast(userIDs, linebot.NewTextMessage("hello")).Do(); err != nil {
//...
//}

//特定のユーザーにメッセージを送信
//bot, err := linebot.New(<channel secret>, <channel token>)
//if err != nil {
//...
//}
//if _, err := bot.PushMessage(<to>, linebot.NewTextMessage("hello")).Do(); err != nil {
//...
//}
