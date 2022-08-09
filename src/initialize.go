package main

import (
	"household.api/src/infrastructure"
	"household.api/src/infrastructure/query_service_imp"
	"household.api/src/infrastructure/repository_imp"
	"household.api/src/presentation/line"
	"household.api/src/presentation/line/controller"
	"household.api/src/usecase/command"
)

type Initialize struct {
	lineHandler *line.LineHandler
	lineBatch   *line.LineBatch
}

func NewInitialize() (init *Initialize, err error) {
	init = new(Initialize)
	//Line関連
	lineBot, err := line.LineInit()
	//Batch関連
	weatherQueryService := query_service_imp.NewFetchWeatherQueryService()
	weatherController := controller.NewWeatherController(weatherQueryService, lineBot)
	batch := line.NewLineBatch(lineBot, *weatherController)
	//Handler関連
	database, err := infrastructure.NewDatabase()
	costRepository := repository_imp.NewCostRepository(database)
	userRepository := repository_imp.NewUserRepository(database)

	costQueryService := query_service_imp.NewCostQueryService(database)

	saveCostUsecase := command.NewSaveCostUseCase(costRepository)
	saveUserUseCase := command.NewSaveUserUseCase(userRepository)

	costController := controller.NewCostController(*saveCostUsecase, costQueryService)
	userController := controller.NewUserController(*saveUserUseCase)

	handler, err := line.NewLineHandler(lineBot, costController, weatherController, userController)

	if err != nil {
		return
	}

	init.lineBatch = batch
	init.lineHandler = handler
	return
}
