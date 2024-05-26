package helpers

import "time"

func CheckDateIntersection(startDate1, endDate1, startDate2, endDate2 time.Time) bool {
	return (startDate1.Before(endDate2) && endDate1.After(startDate2)) ||
		startDate1.Equal(startDate2) || endDate1.Equal(endDate2)
}
