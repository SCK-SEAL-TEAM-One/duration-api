package api

import (
	. "../models"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"net/http"
)

func DurationBetweenTimes(writer http.ResponseWriter, request *http.Request)  {
		duration := Duration{
			StartFullDate : "Thursday, 16 October 1997",
			EndFullDate : "Monday, 10 June 2019",
			DaysToMonths : MonthDay{
				Months:259,
				Days:25,
			},
			DaysToWeeks: WeekDay{
				Weeks:1129,
				Days:4,
			},
			Days : 7907,
			Seconds:683164800,
			Minutes:11386080,
			Hours:189768,
		}
		durationjson,_ := json.Marshal(duration)
		fmt.Fprint(writer, string(durationjson))
}

func GetYearMonthDayFromRequestDate(request *http.Request) RequestDate{
	body, _ := ioutil.ReadAll(request.Body)
	var requestDate RequestDate
	_ = json.Unmarshal(body,&requestDate)
		return requestDate
}