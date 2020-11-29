package model


type Weather struct {
	location    string
	temperature string
	wind        string
}

type WeatherResponse struct {
	Location WeatherLocation
	Current  WeatherStatus
	Forecast WeatherForecast
	//Alert WeatherAlert
}

type WeatherForecast struct {
	Forecastday []WeatherForecastDay
}

type WeatherForecastDay struct {
	Date  string
	Day   WeatherDay
	Astro WeatherAstro
}

type WeatherDay struct {
	Maxtemp_c            float64
	Mintemp_c            float64
	Avgtemp_c            float64
	Maxwind_kph          float64
	Totalprecip_mm       float64
	Avgvis_km            float64
	Avghumidity          float64
	Daily_will_it_rain   int
	Daily_chance_of_rain string
	Daily_will_it_snow   int
	Daily_chance_of_snow string
	Condition            WeatherCondition
	Uv                   float64
}

type WeatherCondition struct {
	Text string
	Icon string
	Code int
}

type WeatherAstro struct {
	Sunrise           string
	Sunset            string
	Moonrise          string
	Moonset           string
	Moon_phase        string
	Moon_illumination string
}

type WeatherLocation struct {
	Name      string
	Region    string
	Country   string
	Lat       float64
	Lon       float64
	Localtime string
}
type WeatherStatus struct {
	Last_updated string
	Temp_c       float64
	Feelslike_c  float64
	Wind_kph     float64
	Wind_dir     string
	Pressure_mb  float64
	Precip_mm    float64
	Humidity     int
	Uv           float64
	// Visibilité
	Vis_km float64
	// Rafale de vent
	Gust_kph  float64
	Condition WeatherCondition
}
