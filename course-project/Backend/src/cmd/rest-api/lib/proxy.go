package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	uri      = "https://community-open-weather-map.p.rapidapi.com/find?q=%s&cnt=1&mode=null&lon=0&lat=0&units=metric"
	rapidApi = "community-open-weather-map.p.rapidapi.com"
	rapidKey = "639fcce71amsh247779e1c92ce51p1583cdjsn3a618c7821a7"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Forecast struct {
	List []map[string]interface{} `json:"list"`
}

type Weather struct {
	CityName    string  `json:"cityName"`
	Temperature float64 `json:"temperature"`
}

func NewRequest(cityName string) (*http.Request, error) {
	r := fmt.Sprintf(uri, cityName)
	req, err := http.NewRequest("GET", r, nil)
	if err != nil {

		return nil, err
	}
	req.Header.Add("x-rapidapi-host", rapidApi)
	req.Header.Add("x-rapidapi-key", rapidKey)
	return req, nil
}

func GetResponseFromWeatherApp(req *http.Request, client HTTPClient) (Weather, error) {
	res, err := client.Do(req)
	if err != nil {
		return Weather{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Weather{}, err
	}
	return ParseJSONFromApi(body)
}

func ParseJSONFromApi(body interface{}) (Weather, error) {
	var ff Forecast
	err := json.Unmarshal(body.([]byte), &ff)
	if err != nil {
		log.Print("json:", err)
		return Weather{}, &ErrorJSONResponse{}
	}
	if len(ff.List) < 1 {
		log.Printf("%v", ff)
		return Weather{}, &ErrorNotFound{}
	}
	l, ok := ff.List[0]["main"].(map[string]interface{})
	if !ok {
		return Weather{}, &ErrorNotFound{}
	}
	forecast := Weather{
		CityName:    ff.List[0]["name"].(string),
		Temperature: l["temp"].(float64),
	}
	return forecast, nil
}

func GetCurrentWeatherFromAPI(ctx context.Context, cityName string) (Weather, error) {
	req, err := NewRequest(cityName)
	if err != nil {
		log.Printf("Error occured: %v", err)
		return Weather{}, err
	}
	client := http.DefaultClient
	res, err := GetResponseFromWeatherApp(req, client)
	if err != nil {
		log.Printf("Error occured: %v", err)
		return Weather{}, err
	}
	return res, nil
}
