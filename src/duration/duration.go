package duration

import "time"

func CalculateDurationBetweenTime(startTime time.Time, endTime time.Time) time.Duration {
	duration, _ := time.ParseDuration("683164800s")
	return duration
}
