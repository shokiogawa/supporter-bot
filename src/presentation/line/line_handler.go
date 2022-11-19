package line

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"household.api/src/presentation/line/line_controller"
	"household.api/src/usecase/query/query_service_interface"
	"log"
	"strconv"
	"strings"
)

//LineHandler ルーターのような役割(どのコントローラーを使用するもの)
type LineHandler struct {
	bot                  *linebot.Client
	costController       *line_controller.CostController
	userController       *line_controller.UserController
	weatherController    *line_controller.WeatherController
	restaurantController *line_controller.RestaurantController
	fixedCostController  *line_controller.FixedCostController
	commonQueryService   query_service_interface.CommonQueryService
}

func NewLineHandler(
	bot *linebot.Client,
	costController *line_controller.CostController,
	weatherController *line_controller.WeatherController,
	userController *line_controller.UserController,
	restaurantController *line_controller.RestaurantController,
	fixedCostController *line_controller.FixedCostController,
	commonQueryService query_service_interface.CommonQueryService) (lineHandler *LineHandler, err error) {
	lineHandler = new(LineHandler)
	lineHandler.bot = bot
	lineHandler.costController = costController
	lineHandler.weatherController = weatherController
	lineHandler.userController = userController
	lineHandler.restaurantController = restaurantController
	lineHandler.fixedCostController = fixedCostController
	lineHandler.commonQueryService = commonQueryService
	if err != nil {
		fmt.Println(err)
	}
	return
}

// EventHandler linebotから渡ってきた値をデータベースに保存する。
func (handler *LineHandler) EventHandler(e echo.Context) (err error) {
	var replyMessage string = ""
	events, err := handler.bot.ParseRequest(e.Request())
	if err != nil {
		return
	}
	//lineUserIdより、publicUserIdを取得
	publicUserId, err := handler.commonQueryService.GetPublicUserId(events[0].Source.UserID)
	if err != nil {
		log.Fatal(err)
	}
	for _, event := range events {
		switch event.Type {
		//友達追加時
		case linebot.EventTypeFollow:
		//テキスト送信時
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				receiveText := message.Text
				//天気
				if strings.Contains(receiveText, "天気") {
					replyMessage, err = handler.weatherController.GetWeather()
					//その他(金額)
				} else if strings.Contains(receiveText, "ユーザー登録") {
					replyMessage, err = handler.userController.SaveUser(event.Source.UserID)
				} else if strings.Contains(receiveText, "今日の支出") {
					replyMessage, err = handler.costController.CostPerDay(publicUserId)
				} else if strings.Contains(receiveText, "今月の支出") {
					replyMessage, err = handler.costController.CostPerMonth(publicUserId)
				} else if strings.Contains(receiveText, "固定費:") {
					//TOOD: 固定費に変更
					replyMessage, err = handler.fixedCostController.SaveFixedCost(message.Text, publicUserId)
				} else if strings.Contains(receiveText, ":") {
					replyMessage, err = handler.costController.SaveCost(message.Text, publicUserId)
				} else {
					replyMessage = receiveText
				}
				//ここもどうにかわかりやすいように変更予定
				_, err = handler.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				if err != nil {
					_, err = handler.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(err.Error())).Do()
					log.Fatal(err)
					if err != nil {
						log.Fatal(err)
					}
				}
			case *linebot.LocationMessage:
				msg := event.Message.(*linebot.LocationMessage)
				lat := strconv.FormatFloat(msg.Latitude, 'f', 2, 64)
				lng := strconv.FormatFloat(msg.Longitude, 'f', 2, 64)
				restaurants, err := handler.restaurantController.GetRestaurant(lat, lng)
				if err != nil {
					fmt.Println(err)
				}
				var ccs []*linebot.CarouselColumn
				for _, rest := range restaurants {
					fmt.Println(rest)
					cc := linebot.NewCarouselColumn(rest.Photo, rest.Name, rest.Address, linebot.NewURIAction("ホットペッパーで開く", rest.URL)).WithImageOptions("#FFFFFF")
					ccs = append(ccs, cc)
				}
				res := linebot.NewTemplateMessage(
					"レストラン一覧",
					linebot.NewCarouselTemplate(ccs...).WithImageOptions("rectangle", "cover"),
				)
				if _, err := handler.bot.ReplyMessage(event.ReplyToken, res).Do(); err != nil {
					fmt.Println("ここ")
					fmt.Println(err)
				}
			}
		}

	}
	return
}
