package api

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestApiDurationBetweenTimesPut_StartDay_16_StartMonth_10_StartYear_1997_EndDay_10_EndMonth_6_EndYear_2019_Should_Get_Duration_Response(t *testing.T){
	url := "/duration"
	body := strings.NewReader(`{"Start": {"day": 16,"month": 10,"year": 1997},"End": {"day": 10,"month": 6,"year": 2019}}`)
	request := httptest.NewRequest("GET",url,body)
	recorder := httptest.NewRecorder()
	expectedResult := `{
		"days":7907,
		"startFullDate":"Thursday, 16 October 1997",
		"endFullDate":"Monday, 10 June 2019",
		"daysToMonths":{
			"months":259,
			"days":25
		},
		"daysToWeeks":{
			"weeks":1129,
			"days":4
		},
		"seconds":683164800,
		"minutes":11386080,
		"hours":189768
		}`

	DurationBetweenTimes(recorder,request)
	response := recorder.Result()
	responseBodyRaw, _ := ioutil.ReadAll(response.Body)
	responseBody := string(responseBodyRaw)



	if expectedResult != responseBody {
		 t.Errorf("need %v but got %v",expectedResult,responseBody)
	}
}
