package query_service_imp

import (
	"encoding/json"
	"fmt"
	"household.api/src/domain/entity"
	"io"
	"net/http"
)

type WeatherMapper struct {
	Area string `json:"targetArea"`
	Headline string `json:"headlineText"`
	Body string `json:"text"`
}

type FetchWeatherQueryService struct {

}

func NewFetchWeatherQueryService()*FetchWeatherQueryService{
	queryService := new(FetchWeatherQueryService)
	return queryService
}

func (queryService *FetchWeatherQueryService) Invoke() (weather *entity.Weather, err error){
	response ,err := http.Get("https://www.jma.go.jp/bosai/forecast/data/overview_forecast/120000.json")
	if err != nil{
		return
	}
	body, err := io.ReadAll(response.Body)
	if err != nil{
		return
	}
	defer response.Body.Close()

	weatherMapper := new(WeatherMapper)
	err = json.Unmarshal(body, weatherMapper)
	if err != nil{
		fmt.Println(err)
		return
	}
	weather = &entity.Weather{
		Area: weatherMapper.Area,
		Headline: weatherMapper.Headline,
		Body: weatherMapper.Body,
	}
	return
}