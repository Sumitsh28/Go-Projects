package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)


type ApiConfigData struct {

	WeatherApiKey string `json:"weatherapikey"`
}

type Weather struct {

	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	}`json:"main"`
}

func configFile(filename string) (ApiConfigData, error) {
	
	bytes,err := os.ReadFile(filename)

	if err != nil {
		return ApiConfigData{},err
	}

	var configData ApiConfigData

	err = json.Unmarshal(bytes, &configData)

	if err != nil {
		return ApiConfigData{},err
	}

	return configData, nil
}
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from go!\n"))
}

func query(city string) (Weather, error) {
	apiConfig, err := configFile(".apiConfig")
	if err != nil {
		return Weather{}, err
	}

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.WeatherApiKey + "&q=" + city)
	if err != nil {
		return Weather{}, err
	}

	defer resp.Body.Close()

	var d Weather
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return Weather{}, err
	}
	return d, nil
}

func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/",
		func(w http.ResponseWriter, r *http.Request) {
			city := strings.SplitN(r.URL.Path, "/", 3)[2]
			data, err := query(city)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(data)
		})

	http.ListenAndServe(":8080", nil)
}