package main

import (
	"household.api/src/infrastructure/query_service_imp"
	"household.api/src/presentation/line"
	"household.api/src/presentation/line/controller"
	"household.api/src/usecase/command"
)

type Initialize struct {
	lineHandler *line.LineHandler
	lineBatch *line.LineBatch
}

func NewInitialize()(init *Initialize, err error){
	init = new(Initialize)
	//Line関連
	lineBot, err := line.LineInit()
	//Batch関連
	weatherQueryService := query_service_imp.NewFetchWeatherQueryService()
	weatherController := controller.NewWeatherController(weatherQueryService, lineBot)
	batch := line.NewLineBatch(lineBot, *weatherController)
	//Handler関連
	saveCostUsecase := command.NewSaveCostUseCase()
	costController := controller.NewCostController(*saveCostUsecase)
	handler, err := line.NewLineHandler(lineBot, *costController, *weatherController)

	if err != nil{
		return
	}

	init.lineBatch = batch
	init.lineHandler = handler
	return
}
