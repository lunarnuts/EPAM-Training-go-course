package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m MockClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	return &http.Response{}, nil
}
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
		{name: "NULL", args: args{cityName: "NULL"}, want: Weather{
			CityName:    "",
			Temperature: 0.0,
		}},
		{name: "NULL", args: args{cityName: ""}, want: Weather{
			CityName:    "Globe",
			Temperature: 0.0,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := GetCurrentWeatherFromAPI(tt.args.cityName); !reflect.DeepEqual(got.CityName, tt.want.CityName) {
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
		wantErr error
	}{
		{name: "Empty Response", args: args{body: b}, want: Weather{}, wantErr: ErrorJSONResponse},
		{name: "Correct Response", args: args{body: b1}, want: Weather{CityName: "test", Temperature: 17.0}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJSONFromApi(tt.args.body)
			if (err != nil) && !errors.Is(err, tt.wantErr) {
				t.Errorf("ParseJSONFromApi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseJSONFromApi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetResponseFromWeatherApp(t *testing.T) {
	request, _ := NewRequest("Astana")
	type args struct {
		req    *http.Request
		client HTTPClient
	}
	tests := []struct {
		name    string
		args    args
		want    Weather
		wantErr error
	}{
		{name: "ok", args: args{req: request, client: MockClient{DoFunc: func(req *http.Request) (*http.Response, error) {
			list := make(map[string]interface{})
			list["temp"] = 17.0
			mymap := make(map[string]interface{})
			mymap["main"] = list
			mymap["name"] = "Astana"
			body := Forecast{
				List: []map[string]interface{}{mymap},
			}
			b, _ := json.Marshal(body)
			res := http.Response{
				StatusCode: http.StatusAccepted,
				Body:       ioutil.NopCloser(bytes.NewBuffer(b)),
			}
			return &res, nil
		}}}, want: Weather{CityName: "Astana", Temperature: 17.0}, wantErr: nil},
		{name: "Not Found", args: args{req: request, client: MockClient{DoFunc: func(req *http.Request) (*http.Response, error) {
			body := Forecast{
				List: []map[string]interface{}{},
			}
			b, _ := json.Marshal(body)
			res := http.Response{
				StatusCode: http.StatusAccepted,
				Body:       ioutil.NopCloser(bytes.NewBuffer(b)),
			}
			return &res, nil
		}}}, want: Weather{}, wantErr: ErrorNotFound},
		{name: "JSON error", args: args{req: request, client: MockClient{DoFunc: func(req *http.Request) (*http.Response, error) {
			res := http.Response{
				StatusCode: http.StatusAccepted,
				Body:       ioutil.NopCloser(bytes.NewBuffer(nil)),
			}
			return &res, nil
		}}}, want: Weather{}, wantErr: ErrorJSONResponse},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetResponseFromWeatherApp(tt.args.req, tt.args.client)
			if (err != nil) && !errors.Is(err, tt.wantErr) {
				t.Errorf("GetResponseFromWeatherApp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetResponseFromWeatherApp() = %v, want %v", got, tt.want)
			}
		})
	}
}
