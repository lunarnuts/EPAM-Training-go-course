package lib

import (
	"context"
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
		{name: "empty", args: args{cityName: ""}, want: Weather{
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
