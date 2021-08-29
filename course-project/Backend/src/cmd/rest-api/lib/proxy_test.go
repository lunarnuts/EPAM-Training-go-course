package lib

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestGetCurrentWeatherFromAPI(t *testing.T) {
	type args struct {
		cityName string
	}
	tests := []struct {
		name string
		args args
		want Weather
	}{
		{name: "london", args: args{cityName: "london"}, want: Weather{
			CityName:    "London",
			Temperature: 0.0,
		}},
		{name: "astana", args: args{cityName: "astana"}, want: Weather{
			CityName:    "Nur-Sultan",
			Temperature: 0.0,
		}},
		{name: "almaty", args: args{cityName: "almaty"}, want: Weather{
			CityName:    "Almaty",
			Temperature: 0.0,
		}},
		{name: "asdasda", args: args{cityName: "asdasda"}, want: Weather{
			CityName:    "",
			Temperature: 0.0,
		}},
		{name: "empty", args: args{cityName: "NULL"}, want: Weather{
			CityName:    "",
			Temperature: 0.0,
		}},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := GetCurrentWeatherFromAPI(ctx, tt.args.cityName); !reflect.DeepEqual(got.CityName, tt.want.CityName) {
				t.Errorf("GetCurrentWeatherFromAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseJSONFromApi(t *testing.T) {
	temp := make([]map[string]interface{}, 0)
	b, _ := json.Marshal(temp)
	t1 := make(map[string]interface{})
	l1 := make(map[string]interface{})
	l1["temp"] = 17.0
	t1["main"] = l1
	t1["name"] = "test"
	temp = append(temp, t1)
	var ff Forecast
	ff.List = temp
	b1, _ := json.Marshal(ff)
	type args struct {
		body []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Weather
		wantErr bool
	}{
		{name: "Empty Response", args: args{body: b}, want: Weather{}, wantErr: true},
		{name: "Correct Response", args: args{body: b1}, want: Weather{CityName: "test", Temperature: 17.0}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJSONFromApi(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseJSONFromApi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseJSONFromApi() = %v, want %v", got, tt.want)
			}
		})
	}
}
