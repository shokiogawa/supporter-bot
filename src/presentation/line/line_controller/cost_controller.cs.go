package line_controller

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"household.api/src/usecase/command"
	"household.api/src/usecase/query/query_service_interface"
	"household.api/src/usecase/query/read_model"
	"log"
	"strconv"
	"strings"
	"time"
)

type CostController struct {
	saveCostUseCase   *command.SaveCostUseCase
	deleteCostUseCase *command.DeleteCostUseCase
	costQueryService  query_service_interface.CostQueryService
	bot               *linebot.Client
}

func NewCostController(
	saveCostUseCase *command.SaveCostUseCase,
	deleteCostUseCase *command.DeleteCostUseCase,
	costQueryService query_service_interface.CostQueryService,
	bot *linebot.Client) *CostController {
	con := new(CostController)
	con.saveCostUseCase = saveCostUseCase
	con.deleteCostUseCase = deleteCostUseCase
	con.costQueryService = costQueryService
	con.bot = bot
	return con
}

func (con *CostController) SaveCost(message string, publicUserId string, replyToken string, lineUserId string) (replyMessage string, err error) {
	content := strings.Split(message, ":")
	title := content[0]
	outcome, err := strconv.Atoi(content[1])
	costId, err := con.saveCostUseCase.Invoke(title, outcome, publicUserId)
	if err != nil {
		fmt.Println(err)
		return
	}

	//メッセージ
	replyMessage = title + "を" + content[1] + "で登録しました。"
	//ポストバックイベントデータ
	data := fmt.Sprintf("type=%s&costId=%s", "delete", strconv.Itoa(int(costId)))
	// 現在のデータ取得。
	currentTime := time.Now().Format("2006年01月02日")
	buttonTemp := linebot.NewButtonsTemplate(
		"",
		currentTime,
		replyMessage,
		linebot.NewPostbackAction("削除する。", data, "", "", "", ""))
	resMessage := linebot.NewTemplateMessage("支出登録　", buttonTemp)
	//LINEに送信
	_, err = con.bot.ReplyMessage(replyToken, resMessage).Do()
	if err != nil {
		log.Fatal(err)
		//上記のエラーの場合、lineUserIdよりpushメッセージで送信。
		_, err = con.bot.PushMessage(lineUserId, resMessage).Do()
		if err != nil {
			log.Fatal(err)
		}
	}
	return
}

func (con *CostController) DeleteCost(costId int, lineUserId string) (err error) {
	err = con.deleteCostUseCase.Invoke(costId)
	if err != nil {
		return
	}
	_, err = con.bot.PushMessage(lineUserId, linebot.NewTextMessage("削除しました。")).Do()
	return
}

func (con *CostController) CostPerDay(publicUserId string) (replyMessage string, err error) {
	costs, err := con.costQueryService.FetchPerDay(publicUserId)
	replyMessage = "今日のこれまでの支出"
	totalOutCome := 0
	for _, cost := range costs {
		costString := strconv.Itoa(cost.OutCome)
		replyMessage = replyMessage + "\n" + cost.Title + ":" + costString
		totalOutCome = totalOutCome + cost.OutCome
	}
	replyMessage = replyMessage + "\n" + "----------------------" + "\n" + "合計:" + strconv.Itoa(totalOutCome) + "円"
	return
}

func (con *CostController) CostPerMonth(publicUserId string) (replyMessage string, err error) {
	costLists, err := con.costQueryService.FetchPerMonth(publicUserId)
	replyMessage = makeMessage(costLists)
	return
}

func (con *CostController) CostPerMonthList() (replyMessages map[string]string, err error) {
	costLists, err := con.costQueryService.FetchPerMonthList()
	if err != nil {
		return
	}
	messageMap := make(map[string]string)
	for _, value := range costLists {
		replyMessage := makeMessage(value.CostSumList)
		replyMessage += "昨日までの今月の支出です。\nこれをもとに今日の支出を考えてください！！"
		messageMap[value.LineUserId] = replyMessage
	}
	replyMessages = messageMap
	return
}

func makeMessage(costs []read_model.CostSumReamModel) (replyMessage string) {
	today := time.Now()
	year := strconv.Itoa(today.Year())
	month := strconv.Itoa(int(today.Month()))
	replyMessage = year + "年" + month + "月" + "の合計支出は。\n"
	totalCost := 0
	for _, cost := range costs {
		replyMessage = replyMessage + cost.Date + ":" + strconv.Itoa(cost.OutCome) + "円" + "\n"
		totalCost += cost.OutCome
	}
	replyMessage += "合計支出:" + strconv.Itoa(totalCost) + "円\n"
	return
}
