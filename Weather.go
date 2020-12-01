package weather_client

import (
	"bytes"
	"encoding/json"
	"github.com/LoicC04/weatherapi-go-client/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const apiForecastUrl string = "http://api.weatherapi.com/v1/forecast.json"

const apiConditionsUrl string = "https://www.weatherapi.com/docs/conditions.json"

// TODO : Conditions

func GetCurrentWeather(loc string, lang_iso string) model.WeatherResponse {
	apiKey := getApiKey()
	fullUrl := apiForecastUrl + "?key=" + apiKey + "&days=2&q=" + loc
	//log.Println("Calling .... " + fullUrl)
	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	/*body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(string(body))*/

	var result model.WeatherResponse
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	result.Current.Condition.Text = GetLocalConditions(result.Current.Condition.Code, lang_iso, false)
	for i, f := range result.Forecast.Forecastday {
		result.Forecast.Forecastday[i].Day.Condition.Text = GetLocalConditions(f.Day.Condition.Code, "fr", false)
	}
	return result
}

func GetLocalConditions(code int, lang_iso string, night bool) string {
	fullUrl := apiConditionsUrl

	// HTTP Call
	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read to String
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Escape BOM Characters
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}

	// JSON --> Conditions
	var result model.Conditions
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	for _, r := range result {
		if r.Code == code {
			for _, l := range r.Languages {
				if l.Lang_iso == lang_iso {
					if night {
						return l.Night_text
					} else {
						return l.Day_text
					}
				}
			}
		}
	}

	return ""
}

func getApiKey() string {
	apiKey := os.Getenv("WEATHERAPI_APIKEY")
	if apiKey == "" {
		log.Fatal("WEATHERAPI_APIKEY is not defined")
		os.Exit(1)
	}
	return apiKey
}
