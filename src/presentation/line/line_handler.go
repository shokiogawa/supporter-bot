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
	if err != nil {
		log.Fatal(err)
	}
	for _, event := range events {
		switch event.Type {
		//ユーザーがアカウントフォロー時
		case linebot.EventTypeFollow:
			replyMessage, err = handler.userController.SaveUser(event.Source.UserID)
		case linebot.EventTypePostback:
			data := event.Postback.Data
			paramsMap := GetParamsMap(data)
			if paramsMap["type"] == "delete" {
				costId, err := strconv.Atoi(paramsMap["costId"])
				if err != nil {
					log.Fatal("costIdが文字列です。")
				}
				err = handler.costController.DeleteCost(costId, event.Source.UserID)
				if err != nil {
					log.Fatal("削除に失敗しました。")
				}
			}
		//テキスト送信時
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				receiveText := message.Text
				//userid取得
				publicUserId, err := handler.commonQueryService.GetPublicUserId(event.Source.UserID)
				//天気
				if strings.Contains(receiveText, "天気") {
					replyMessage, err = handler.weatherController.GetWeather()
					_, err = handler.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
					//その他(金額)
				} else if strings.Contains(receiveText, "今日の支出") {
					replyMessage, err = handler.costController.CostPerDay(publicUserId)
					_, err = handler.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				} else if strings.Contains(receiveText, "今月の支出") {
					replyMessage, err = handler.costController.CostPerMonth(publicUserId)
					_, err = handler.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				} else if strings.Contains(receiveText, "現在の固定費") {
					replyMessage, err = handler.fixedCostController.GetFixedCostList(publicUserId)
					_, err = handler.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				} else if strings.Contains(receiveText, "固定費:") {
					//TOOD: 固定費に変更
					replyMessage, err = handler.fixedCostController.SaveFixedCost(message.Text, publicUserId)
					_, err = handler.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				} else if strings.Contains(receiveText, ":") {
					//支出保存
					replyMessage, err = handler.costController.SaveCost(message.Text, publicUserId, event.ReplyToken, event.Source.UserID)
				} else {
					replyMessage = receiveText
				}
				//ここもどうにかわかりやすいように変更予定
				//_, err = handler.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				if err != nil {
					_, err = handler.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("エラーが発生しました。管理者にお問合せください。")).Do()
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

func GetParamsMap(queryString string) (queryMap map[string]string) {
	queryMap = make(map[string]string)
	querys := strings.Split(queryString, "&")
	for _, query := range querys {
		keyAndValue := strings.Split(query, "=")
		key := keyAndValue[0]
		value := keyAndValue[1]
		queryMap[key] = value
	}
	return
}
