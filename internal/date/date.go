package date

import (
	"time"
)

// GetShortMonth returns month in short word format (Jan, Feb, etc.)
func GetShortMonth(t time.Time) string {
	return t.Format("Jan")
}

// GetQuarter returns the quarter (1-4) for a given time
func GetQuarter(t time.Time) int {
	month := t.Month()
	switch {
	case month <= 3:
		return 1
	case month <= 6:
		return 2
	case month <= 9:
		return 3
	default:
		return 4
	}
}

// GetPrevQuarter returns the previous quarter and its year
func GetPrevQuarter(quarter, year int) (int, int) {
	if quarter == 1 {
		return 4, year - 1
	}
	return quarter - 1, year
}

// GetNextQuarter returns the next quarter and its year
func GetNextQuarter(quarter, year int) (int, int) {
	if quarter == 4 {
		return 1, year + 1
	}
	return quarter + 1, year
}

// GetQuarterMonths returns the three months in a quarter as short names
func GetQuarterMonths(quarter int, year int) []string {
	firstMonthNum := (quarter-1)*3 + 1
	months := make([]string, 3)

	for i := range [3]int{0, 1, 2} {
		t := time.Date(year, time.Month(firstMonthNum+i), 1, 0, 0, 0, 0, time.Local)
		months[i] = GetShortMonth(t)
	}

	return months
}

// GetPrevMonth returns the previous month and its year
func GetPrevMonth(month time.Month, year int) (string, int) {
	var prevMonth time.Month
	var prevYear int

	if month == time.January {
		prevMonth = time.December
		prevYear = year - 1
	} else {
		prevMonth = month - 1
		prevYear = year
	}

	t := time.Date(prevYear, prevMonth, 1, 0, 0, 0, 0, time.Local)
	return GetShortMonth(t), prevYear
}

// GetNextMonth returns the next month and its year
func GetNextMonth(month time.Month, year int) (string, int) {
	var nextMonth time.Month
	var nextYear int

	if month == time.December {
		nextMonth = time.January
		nextYear = year + 1
	} else {
		nextMonth = month + 1
		nextYear = year
	}

	t := time.Date(nextYear, nextMonth, 1, 0, 0, 0, 0, time.Local)
	return GetShortMonth(t), nextYear
}

// GetMonthWeeks returns all days in a month organized by ISO week
func GetMonthWeeks(month time.Month, year int) []Week {
	// Create a date at the first day of the month
	t := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)

	var weeks []Week
	var currentWeek Week
	var lastISOWeek int

	// Process each day of the month
	for {
		if t.Month() != month {
			// If we have days collected for the last week, add it
			if len(currentWeek.Days) > 0 {
				weeks = append(weeks, currentWeek)
			}
			break
		}

		// Get ISO week
		_, isoWeek := t.ISOWeek()

		// If this is a new week, finalize the previous one (if any) and start a new one
		if isoWeek != lastISOWeek && len(currentWeek.Days) > 0 {
			weeks = append(weeks, currentWeek)
			currentWeek = Week{WeekNum: isoWeek, Days: []Day{}}
		} else if len(currentWeek.Days) == 0 {
			currentWeek = Week{WeekNum: isoWeek, Days: []Day{}}
		}

		// Add this day to the current week
		day := Day{
			Day:    t.Day(),
			IsLast: false,
		}
		currentWeek.Days = append(currentWeek.Days, day)

		lastISOWeek = isoWeek
		t = t.AddDate(0, 0, 1)
	}

	// Mark the last day in each week
	for i := range weeks {
		if len(weeks[i].Days) > 0 {
			weeks[i].Days[len(weeks[i].Days)-1].IsLast = true
		}
	}

	return weeks
}

// GetAdjacentDays returns the previous and next day
func GetAdjacentDays(day int, month time.Month, year int) (int, string, int, int, string, int) {
	date := time.Date(year, month, day, 0, 0, 0, 0, time.Local)

	// Previous day
	prevDate := date.AddDate(0, 0, -1)
	prevDay := prevDate.Day()
	prevMonth := GetShortMonth(prevDate)
	prevYear := prevDate.Year()

	// Next day
	nextDate := date.AddDate(0, 0, 1)
	nextDay := nextDate.Day()
	nextMonth := GetShortMonth(nextDate)
	nextYear := nextDate.Year()

	return prevDay, prevMonth, prevYear, nextDay, nextMonth, nextYear
}

// Week represents a week with its number and days
type Week struct {
	WeekNum int
	Days    []Day
}

// Day represents a day in a week
type Day struct {
	Day    int
	IsLast bool
}

