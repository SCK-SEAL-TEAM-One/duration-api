package duration

import "time"

func CalculateDurationBetweenTime(startTime time.Time, endTime time.Time) time.Duration {
	return endTime.Sub(startTime)
}
