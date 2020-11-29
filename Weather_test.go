package weather_client

import (
	"fmt"
	"strconv"
	"testing"
)

func TestGetWeather(t *testing.T) {
	// GIVEN
	want := "Niort"

	// WHEN
	got := GetCurrentWeather(want, "fr")

	// THEN
	if got.Location.Name != want {
		t.Errorf("getCurrentWeather() = %q, want %q", got.Location.Name, want)
	}
	var rain string = "Non"
	if got.Forecast.Forecastday[0].Day.Daily_will_it_rain == 1 {
		rain = "Oui"
	}
	fmt.Println(got.Location.Localtime + " - Météo à " + got.Location.Name + " : " + got.Forecast.Forecastday[0].Day.Condition.Text + " - " + "Va-t-il pleuvoir ? " + rain)
	temp := strconv.FormatFloat(got.Current.Temp_c, 'f', -1, 64)
	wind := strconv.FormatFloat(got.Current.Wind_kph, 'f', -1, 64)
	fmt.Println("Maintenant ---- " + temp + "°C - Vent " + got.Current.Wind_dir + " à " + wind + " km/h")
	tempMin := strconv.FormatFloat(got.Forecast.Forecastday[0].Day.Mintemp_c, 'f', -1, 64)
	tempMax := strconv.FormatFloat(got.Forecast.Forecastday[0].Day.Maxtemp_c, 'f', -1, 64)
	fmt.Println("Température ---- Min " + tempMin + "°C - Max " + tempMax + "°C")
	windMax := strconv.FormatFloat(got.Forecast.Forecastday[0].Day.Maxwind_kph, 'f', -1, 64)
	rafale := strconv.FormatFloat(got.Current.Gust_kph, 'f', -1, 64)
	fmt.Println("Vent ---- " + windMax + " km/h max - Rafale à " + rafale + " km/h")

}
