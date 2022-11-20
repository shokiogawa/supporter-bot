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
	fixedCostRepository := repository_imp.NewFixedCostRepository(init.database)

	costQueryService := query_service_imp.NewCostQueryService(init.database)
	restaurantQueryService := query_service_imp.NewFetchRestaurantQueryService()
	commonQueryService := query_service_imp.NewCommonQueryService(init.database)
	fixedCostQueryService := query_service_imp.NewFixedCostQueryService(init.database)

	saveCostUsecase := command.NewSaveCostUseCase(costRepository)
	saveUserUseCase := command.NewSaveUserUseCase(userRepository)
	saveFixedCostUseCase := command.NewSaveFixedCostUseCase(fixedCostRepository)
	deleteCostUseCase := command.NewDeleteCostUseCase(costRepository)

	//line controller
	costController := line_controller.NewCostController(saveCostUsecase, deleteCostUseCase, costQueryService, lineBot)
	userController := line_controller.NewUserController(*saveUserUseCase)
	restaurantController := line_controller.NewRestaurantController(restaurantQueryService)
	fixedCostController := line_controller.NewFixedCostUseCase(saveFixedCostUseCase, fixedCostQueryService)

	//front controller
	frontCostController := front_controller.NewCostController(*saveCostUsecase, costQueryService)

	frontHandler := front.NewHandler(frontCostController)
	lineHandler, err := line.NewLineHandler(
		lineBot,
		costController,
		weatherController,
		userController,
		restaurantController,
		fixedCostController,
		commonQueryService)
	batch := line.NewLineBatch(lineBot, *weatherController, *costController)
	if err != nil {
		return
	}

	init.lineBatch = batch
	init.lineHandler = lineHandler
	init.frontHandler = frontHandler
	return
}
