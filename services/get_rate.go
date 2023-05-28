package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	baseUrl      = "https://api.binance.com"
	pathEndpoint = "/api/v3/avgPrice"
)

func GetBasePrice() string {
	_url, _ := url.ParseRequestURI(baseUrl)
	_url.Path = pathEndpoint
	parameters := url.Values{}
	parameters.Add("symbol", "BTCUAH")
	_url.RawQuery = parameters.Encode()

	response, _ := http.Get(fmt.Sprintf("%v", _url))
	responseBody, _ := ioutil.ReadAll(response.Body)

	var price struct {
		Minutes int    `json:"mins"`
		Price   string `json:"price"`
	}

	err := json.Unmarshal([]byte(responseBody), &price)
	if err != nil {
		fmt.Println(err.Error())
	}

	return price.Price
}

func GetPrice(w http.ResponseWriter, r *http.Request) {
	price := GetBasePrice()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{rate:%s}", price)))
}
