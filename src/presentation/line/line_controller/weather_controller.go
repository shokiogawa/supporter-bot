package line_controller

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"household.api/src/usecase/query/query_service_interface"
)

type WeatherController struct {
	fetchWeatherQueryService query_service_interface.FetchWeatherQueryService
	linebot                  *linebot.Client
}

func NewWeatherController(fetchWeatherQueryService query_service_interface.FetchWeatherQueryService, linebot *linebot.Client) *WeatherController {
	controller := new(WeatherController)
	controller.fetchWeatherQueryService = fetchWeatherQueryService
	controller.linebot = linebot
	return controller
}

func (con *WeatherController) GetWeather() (replyMessage string, err error) {
	weather, err := con.fetchWeatherQueryService.Invoke()
	if err != nil {
		return
	}

	area := fmt.Sprintf("%sの天気です。\n", weather.Area)
	head := fmt.Sprintf("%s\n", weather.Headline)
	body := fmt.Sprintf("%s\n", weather.Body)
	replyMessage = area + head + body
	return
}
