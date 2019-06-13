package duration

import (
	"encoding/json"
)

func GetYearMonthDayFromRequestDate(datejson string) (year int,month int,day int) {
	var jsonBlob = []byte(datejson)
	type Date struct{
		Year int `json:"year"`
		Month int `json:"month"`
		Day int `json:"day"`
	}
	var date Date
	_ = json.Unmarshal(jsonBlob,&date)
	return date.Year,date.Month,date.Day
}
