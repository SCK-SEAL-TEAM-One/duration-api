package models

type RequestDate struct {
	Start YearMonthDay
	End   YearMonthDay
}

type YearMonthDay struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type MonthDay struct {
	Months int `json:"months"`
	Days   int `json:"days"`
}

type WeekDay struct {
	Weeks int `json:"weeks"`
	Days  int `json:"days"`
}

type Duration struct {
	Days          int      `json:"days"`
	StartFullDate string   `json:"startFullDate"`
	EndFullDate   string   `json:"endFullDate"`
	DaysToMonths  MonthDay `json:"daysToMonths"`
	DaysToWeeks   WeekDay  `json:"daysToWeeks"`
	Seconds       int      `json:"seconds"`
	Minutes       int      `json:"minutes"`
	Hours         int      `json:"hours"`
}
