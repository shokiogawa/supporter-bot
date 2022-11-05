package main

import (
	"household.api/src/infrastructure"
	"household.api/src/infrastructure/query_service_imp"
	"household.api/src/infrastructure/repository_imp"
	"household.api/src/presentation/front"
	"household.api/src/presentation/front/front_controller"
	"household.api/src/presentation/line"
	"household.api/src/presentation/line/line_controller"
	"household.api/src/usecase/command"
)

type Initialize struct {
	lineHandler  *line.LineHandler
	lineBatch    *line.LineBatch
	frontHandler *front.Handler
	database     *infrastructure.Database
}

func NewInitialize() (init *Initialize, err error) {
	init = new(Initialize)
	//Line関連
	lineBot, err := line.LineInit()
	//Batch関連
	weatherQueryService := query_service_imp.NewFetchWeatherQueryService()
	weatherController := line_controller.NewWeatherController(weatherQueryService, lineBot)

	//Handler関連
	init.database, err = infrastructure.NewDatabase()
	costRepository := repository_imp.NewCostRepository(init.database)
	userRepository := repository_imp.NewUserRepository(init.database)

	costQueryService := query_service_imp.NewCostQueryService(init.database)
	restaurantQueryService := query_service_imp.NewFetchRestaurantQueryService()

	saveCostUsecase := command.NewSaveCostUseCase(costRepository)
	saveUserUseCase := command.NewSaveUserUseCase(userRepository)

	//line controller
	costController := line_controller.NewCostController(*saveCostUsecase, costQueryService)
	userController := line_controller.NewUserController(*saveUserUseCase)
	restaurantController := line_controller.NewRestaurantController(restaurantQueryService)

	//front controller
	frontCostController := front_controller.NewCostController(*saveCostUsecase, costQueryService)

	frontHandler := front.NewHandler(frontCostController)
	lineHandler, err := line.NewLineHandler(lineBot, costController, weatherController, userController, restaurantController)
	batch := line.NewLineBatch(lineBot, *weatherController, *costController)
	if err != nil {
		return
	}

	init.lineBatch = batch
	init.lineHandler = lineHandler
	init.frontHandler = frontHandler
	return
}
