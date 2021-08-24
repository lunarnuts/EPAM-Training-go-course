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

func GetCurrentWeatherFromAPI(ctx context.Context, cityName string) (Weather, error) {
	r := fmt.Sprintf(uri, cityName) //"https://" + rapidApi + "/find?q=" + cityName + "&cnt=1&mode=null&lon=0&lat=0&units=metric"
	req, err := http.NewRequest("GET", r, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("x-rapidapi-host", rapidApi)
	req.Header.Add("x-rapidapi-key", rapidKey)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
		return Weather{}, fmt.Errorf("WeatherApi service unreachable: %v", err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var ff Forecast
	err = json.Unmarshal(body, &ff)
	if err != nil {
		log.Print(err)
		return Weather{}, fmt.Errorf("Cant parse JSON: %v", err)
	}
	l, ok := ff.List[0]["main"].(map[string]interface{})
	if !ok {
		return Weather{}, fmt.Errorf("WeatherApi - Empty Response")
	}
	forecast := Weather{
		CityName:    ff.List[0]["name"].(string),
		Temperature: l["temp"].(float64),
	}
	return forecast, nil
}
