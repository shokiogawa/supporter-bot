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
	Title   string `db:"title"`
	OutCome int    `db:"outcome"`
}

type ReceiveSumCost struct {
	OutCome int    `db:"SUM(outcome)"`
	Date    string `db:"date"`
}

func (qs *CostQueryService) FetchPerMonth(lineUserId string) (readModel []read_model.CostSumMonthReadModel, err error) {
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
	query = `SELECT SUM(outcome),DATE_FORMAT(created_at, '%Y年%m月%d日') as date FROM costs WHERE user_id = ? AND created_at BETWEEN ? AND ? GROUP BY DATE_FORMAT(created_at, '%Y年%m月%d日');`
	var receiveSumCosts []ReceiveSumCost
	err = db.Select(&receiveSumCosts, query, receiceUserIdVar.Id, firstDate, lastDate)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, cost := range receiveSumCosts {
		readModel = append(readModel, read_model.CostSumMonthReadModel{
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
