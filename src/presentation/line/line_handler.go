package line

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"household.api/src/presentation/line/controller"
	"strings"
)
//LineHandler ルーターのような役割(どのコントローラーを使用するもの)
type LineHandler struct {
	bot            *linebot.Client
	costController controller.CostController
	weatherController controller.WeatherController
}

func NewLineHandler(bot *linebot.Client, costController controller.CostController,weatherController controller.WeatherController) (lineHandler *LineHandler, err error) {
	lineHandler = new(LineHandler)
	lineHandler.bot = bot
	lineHandler.costController = costController
	lineHandler.weatherController = weatherController
	if err != nil{
		log.Fatal(err)
	}
	return
}
// EventHandler linebotから渡ってきた値をデータベースに保存する。
func (handler *LineHandler) EventHandler(e echo.Context)(err error){
	events, err := handler.bot.ParseRequest(e.Request())
	if err != nil{
		return
	}
	for _,event := range events{
			switch event.Type {
			//友達追加時
			case linebot.EventTypeFollow:

			//テキスト送信時
			case linebot.EventTypeMessage:
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					receiveText := message.Text
					var replyMessage string
					//天気
					if strings.Contains(receiveText, "天気"){
						replyMessage , err = handler.weatherController.GetWeather()
						if err != nil{
							return
						}
						//その他(金額)
					}else {
						replyMessage , err = handler.costController.SaveCost(message.Text,event.Source.UserID)
						if err != nil{
							_,err = handler.bot.ReplyMessage(event.ReplyToken,linebot.NewTextMessage("フォーマットが正しくありません。")).Do()
							if err != nil{
								fmt.Println(err)
							}
						}
					}
					_,err = handler.bot.ReplyMessage(event.ReplyToken,linebot.NewTextMessage(replyMessage)).Do()
					if err != nil{
						return
					}
				}
			}
	}
	return
}