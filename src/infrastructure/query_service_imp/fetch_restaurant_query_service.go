package query_service_imp

import (
	"encoding/json"
	"fmt"
	"household.api/src/domain/entity"
	"io"
	"net/http"
	"os"
)

type FetchRestaurantQueryService struct {
}

type response struct {
	Results results `json:"results"`
}

type results struct {
	Shop []shop `json:"shop"`
}

type shop struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Photo   photo  `json:"photo"`
	Urls    urls   `json:"urls"`
}
type photo struct {
	Mobile mobile `json:"mobile"`
}
type mobile struct {
	L string `json:"l"`
}
type urls struct {
	PC string `json:"pc"`
}

func NewFetchRestaurantQueryService() *FetchRestaurantQueryService {
	qs := new(FetchRestaurantQueryService)
	return qs
}

func (qs *FetchRestaurantQueryService) Invoke(lat string, lang string) (restaurants []*entity.Restaurant, err error) {
	fmt.Println(lat)
	fmt.Println(lang)
	apiKey := os.Getenv("HOTPEPPER_KEY")
	url := fmt.Sprintf("http://webservice.recruit.co.jp/hotpepper/gourmet/v1/?key=%s&lat=%s&lng=%s&format=json", apiKey, lat, lang)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var datas response
	if err := json.Unmarshal(body, &datas); err != nil {
		fmt.Println(err)
	}

	for _, data := range datas.Results.Shop {
		fmt.Println(data.Urls.PC)
		restaurants = append(restaurants, &entity.Restaurant{
			Name:    data.Name,
			Address: data.Address,
			Photo:   data.Photo.Mobile.L,
			URL:     data.Urls.PC})
	}
	return
}
