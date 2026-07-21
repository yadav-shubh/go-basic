package main

import (
	"fmt"
	"time"
)

func main() {
	start, end := GetCurrentWeekDateRange(time.Local)
	fmt.Println(start, end)

}

func GetCurrentWeekDateRange(loc *time.Location) (string, string) {
	now := time.Now().In(loc)

	weekday := int(now.Weekday())
	if weekday == 0 { // Sunday
		weekday = 7
	}

	// Monday (DATE)
	start := now.AddDate(0, 0, -(weekday - 1))
	start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, loc)

	// Sunday (DATE)
	end := start.AddDate(0, 0, 6)
	end = time.Date(end.Year(), end.Month(), end.Day(), 0, 0, 0, 0, loc)

	return start.Format("2006-01-02"), end.Format("2006-01-02")
}
