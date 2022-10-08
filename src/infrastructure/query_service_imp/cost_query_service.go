package query_service_imp

import (
	"fmt"
	"household.api/src/domain/entity"
	"household.api/src/infrastructure"
	"household.api/src/usecase/query/read_model"
	"time"
)

type CostQueryService struct {
	database *infrastructure.Database
}

func NewCostQueryService(database *infrastructure.Database) *CostQueryService {
	qs := new(CostQueryService)
	qs.database = database
	return qs
}

type ReceiveUserId struct {
	Id int `db:"id"`
}

type ReceiveCost struct {
	Title      string `db:"title"`
	OutCome    int    `db:"outcome"`
	LineUserId string `dc:"line_user_id"`
}

type ReceiveSumCost struct {
	LineUserId string `db:"line_user_id"`
	OutCome    int    `db:"outcome"`
	Date       string `db:"date"`
}

func (qs *CostQueryService) FetchPerMonth(lineUserId string) (readModel []read_model.CostSumReamModel, err error) {
	db, err := qs.database.Connect()
	if err != nil {
		return
	}
	query := `SELECT id FROM users WHERE line_user_id = ?`
	var receiceUserIdVar ReceiveUserId
	err = db.Get(&receiceUserIdVar, query, lineUserId)
	if err != nil {
		fmt.Println(err)
		return
	}
	today := time.Now()
	firstDate := time.Date(today.Year(), today.Month(), 1, 0, 0, 0, 0, time.Local)
	lastDate := time.Date(today.Year(), today.Month(), 1, 0, 0, 0, 0, time.Local).AddDate(0, 1, -1)
	query = `SELECT SUM(outcome) as outcome,DATE_FORMAT(created_at, '%Y年%m月%d日') as date FROM costs WHERE user_id = ? AND created_at BETWEEN ? AND ? GROUP BY DATE_FORMAT(created_at, '%Y年%m月%d日');`
	var receiveSumCosts []ReceiveSumCost
	err = db.Select(&receiveSumCosts, query, receiceUserIdVar.Id, firstDate, lastDate)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, cost := range receiveSumCosts {
		readModel = append(readModel, read_model.CostSumReamModel{
			OutCome: cost.OutCome,
			Date:    cost.Date,
		})
	}
	return
}

func (qs *CostQueryService) FetchPerDay(lineUserId string) (listCost []entity.Cost, err error) {
	db, err := qs.database.Connect()
	if err != nil {
		return
	}
	query := `SELECT id FROM users WHERE line_user_id = ?`
	var receiceUserIdVar ReceiveUserId
	err = db.Get(&receiceUserIdVar, query, lineUserId)
	if err != nil {
		fmt.Println(err)
		return
	}

	today := time.Now()
	firstDateTime := time.Date(today.Year(), today.Month(), today.Day(), 00, 00, 00, 00, time.Local)
	lastDateTime := time.Date(today.Year(), today.Month(), today.Day(), 23, 59, 59, 00, time.Local)

	query = `SELECT title, outcome FROM costs WHERE user_id = ? AND created_at BETWEEN ? AND ?`
	var receiveVars []ReceiveCost
	err = db.Select(&receiveVars, query, receiceUserIdVar.Id, firstDateTime, lastDateTime)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, cost := range receiveVars {
		listCost = append(listCost, entity.Cost{
			Title:   cost.Title,
			OutCome: cost.OutCome,
		})
	}
	return
}

// FetchPerMonthList 今月の支出を前ユーザーに送信する
func (qs *CostQueryService) FetchPerMonthList() (readModel []read_model.CostSumListReadModel, err error) {
	db, err := qs.database.Connect()
	if err != nil {
		return
	}
	today := time.Now()
	firstDate := time.Date(today.Year(), today.Month(), 1, 0, 0, 0, 0, time.Local)
	lastDate := time.Date(today.Year(), today.Month(), 1, 0, 0, 0, 0, time.Local).AddDate(0, 1, -1)
	//usersテーブルを軸にデータを抽出(1度も支出を保存していない人は今月の支出は0円であることを伝えるため。)
	query := `SELECT users.line_user_id as line_user_id, SUM(costs.outcome) as outcome ,DATE_FORMAT(costs.created_at, '%Y年%m月%d日') as date FROM users LEFT OUTER JOIN costs ON users.id = costs.user_id WHERE costs.created_at BETWEEN ? AND ? GROUP BY users.line_user_id, DATE_FORMAT(costs.created_at, '%Y年%m月%d日');`
	var receiveSumCosts []ReceiveSumCost
	err = db.Select(&receiveSumCosts, query, firstDate, lastDate)
	if err != nil {
		fmt.Println(err)
		return
	}
	costsMap := make(map[string][]read_model.CostSumReamModel)
	for _, cost := range receiveSumCosts {
		costsMap[cost.LineUserId] = append(costsMap[cost.LineUserId], read_model.CostSumReamModel{OutCome: cost.OutCome, Date: cost.Date})
	}
	for key, value := range costsMap {
		readModel = append(readModel, read_model.CostSumListReadModel{
			LineUserId:  key,
			CostSumList: value,
		})
	}
	return
}
