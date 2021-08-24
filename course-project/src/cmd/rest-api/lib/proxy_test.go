package lib

import (
	"reflect"
	"testing"
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
			Temperature: "",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCurrentWeatherFromAPI(tt.args.cityName); !reflect.DeepEqual(got.CityName, tt.want.CityName) {
				t.Errorf("GetCurrentWeatherFromAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}
