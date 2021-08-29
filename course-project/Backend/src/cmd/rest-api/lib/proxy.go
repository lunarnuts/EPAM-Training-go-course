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

func GetResponseFromWeatherApp(req *http.Request) ([]byte, error) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func ParseJSONFromApi(body []byte) (Weather, error) {
	var ff Forecast
	err := json.Unmarshal(body, &ff)
	if err != nil {
		log.Print(err)
		return Weather{}, fmt.Errorf("cant parse JSON: %v", err)
	}
	if len(ff.List) < 1 {
		log.Printf("%v", ff)
		return Weather{}, fmt.Errorf("weatherApi - Empty Response: len < 1")
	}
	l, ok := ff.List[0]["main"].(map[string]interface{})
	if !ok {
		return Weather{}, fmt.Errorf("weatherApi - Empty Response")
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
	res, err := GetResponseFromWeatherApp(req)
	if err != nil {
		log.Printf("Error occured: %v", err)
		return Weather{}, err
	}
	forecast, err := ParseJSONFromApi(res)
	if err != nil {
		log.Printf("Error occured: %v", err)
		return Weather{}, err
	}
	return forecast, nil
}
