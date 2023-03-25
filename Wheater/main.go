package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const DateTime = "2006-01-02 15:04:05"

type Main struct {
	Temp float32 `json:"temp"`
}

type Forecast struct {
	DateTime string `json:"dt_txt"`
	Main     Main   `json:"main"`
}

type Payload struct {
	City interface{} `json:"city"`
	List []Forecast  `json:"list"`
}

func Day(day time.Weekday) string {
	var result string
	switch day.String() {
	case "Monday":
		result = "Mon"
	case "Tuesday":
		result = "Tue"
	case "Wednesday":
		result = "Wed"
	case "Thursday":
		result = "Thu"
	case "Friday":
		result = "Fri"
	case "Saturday":
		result = "Sat"
	case "Sunday":
		result = "Sun"
	}
	return result
}

func Month(month time.Month) string {
	var result string
	switch month.String() {
	case "January":
		result = "Jan"
	case "February":
		result = "Feb"
	case "March":
		result = "Mar"
	case "April":
		result = "Apr"
	case "May":
		result = "May"
	case "June":
		result = "Jun"
	case "July":
		result = "Jul"
	case "August":
		result = "Aug"
	case "September":
		result = "Sep"
	case "October":
		result = "Oct"
	case "November":
		result = "Nov"
	case "December":
		result = "Dec"
	}
	return result
}

func GetWheater(lat string, lon string, key string) (*Payload, error) {
	var payload Payload

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast?lat=%s&lon=%s&units=metric&appid=%s", lat, lon, key)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&payload); err != nil {
		return nil, err
	}

	return &payload, nil
}

func main() {
	var current int
	res, err := GetWheater("-6.200000", "106.816666", "5cdbbfb7471b430e02c433b6c5fa2385")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Wheater Forcast:")
	for _, data := range res.List {
		dt, _ := time.Parse(DateTime, data.DateTime)
		wd := Day(dt.Weekday())
		y, m, d := dt.Date()

		if current != d {
			fmt.Printf("%v, %d %v %d: %.2f°C\n", wd, d, Month(m), y, data.Main.Temp)
			current = d
		}
	}
}
