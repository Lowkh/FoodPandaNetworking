package jsonserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// WeatherJSON is the data struct for JSON URL
type WeatherJSON struct {
	Items []struct {
		UpdateTimestamp time.Time `json:"update_timestamp"`
		Timestamp       time.Time `json:"timestamp"`
		ValidPeriod     struct {
			Start time.Time `json:"start"`
			End   time.Time `json:"end"`
		} `json:"valid_period"`
		General struct {
			Forecast         string `json:"forecast"`
			RelativeHumidity struct {
				Low  int `json:"low"`
				High int `json:"high"`
			} `json:"relative_humidity"`
			Temperature struct {
				Low  int `json:"low"`
				High int `json:"high"`
			} `json:"temperature"`
			Wind struct {
				Speed struct {
					Low  int `json:"low"`
					High int `json:"high"`
				} `json:"speed"`
				Direction string `json:"direction"`
			} `json:"wind"`
		} `json:"general"`
		Periods []struct {
			Time struct {
				Start time.Time `json:"start"`
				End   time.Time `json:"end"`
			} `json:"time"`
			Regions struct {
				West    string `json:"west"`
				East    string `json:"east"`
				Central string `json:"central"`
				South   string `json:"south"`
				North   string `json:"north"`
			} `json:"regions"`
		} `json:"periods"`
	} `json:"items"`
	APIInfo struct {
		Status string `json:"status"`
	} `json:"api_info"`
}

func weatherpage(ress http.ResponseWriter, req *http.Request) {
	result, err := http.Get("https://api.data.gov.sg/v1/environment/24-hour-weather-forecast")

	if err != nil {
		log.Println(err.Error())
	}

	JSONData, _ := ioutil.ReadAll(result.Body)

	var weather WeatherJSON

	err = json.Unmarshal(JSONData, &weather)
	if err != nil {
		fmt.Println(err)
	}

	forecast := "<br>Forecast: " + weather.Items[0].General.Forecast

	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<br><a href="/temperature">Temperature</a>
	</body></html>`

	ress.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(ress, "<strong>Weather Forecast</strong><br>")
	fmt.Fprintln(ress, forecast)
	fmt.Fprintln(ress, body)
}

func temperaturepage(ress http.ResponseWriter, req *http.Request) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<br><a href="/">Weather</a>
	</body></html>`
	ress.Header().Set("Content-Type", "text/html; charset=utf-8")

	result, err := http.Get("https://api.data.gov.sg/v1/environment/24-hour-weather-forecast")

	if err != nil {
		log.Println(err.Error())
	}

	JSONData, _ := ioutil.ReadAll(result.Body)

	var weather WeatherJSON

	err = json.Unmarshal(JSONData, &weather)
	if err != nil {
		fmt.Println(err)
	}

	highTemperatureString := "<br>Highest Temperature: " + strconv.FormatInt(int64(weather.Items[0].General.Temperature.High), 10)
	lowTemperatureString := "<br>Lowest Temperature: " + strconv.FormatInt(int64(weather.Items[0].General.Temperature.Low), 10) + "<br>"

	fmt.Fprintln(ress, "<strong>Weather Forecast</strong><br>")
	fmt.Fprintln(ress, highTemperatureString)
	fmt.Fprintln(ress, lowTemperatureString)
	fmt.Fprintln(ress, body)
}

// StartJSONQuery query JSON URL
func StartJSONQuery() {

	http.Handle("/", http.HandlerFunc(weatherpage))
	http.Handle("/temperature", http.HandlerFunc(temperaturepage))

	http.ListenAndServe(":5331", nil)

}
