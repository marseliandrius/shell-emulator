package commands

import (
	"fmt"
	"strings"
	"time"
)

const (
	minCalendarYear    = 1
	maxCalendarYear    = 9999
	firstCalendarDay   = 1
	calendarDaysInWeek = 7
	calendarDayWidth   = 3
	calendarWeekHeader = "Su Mo Tu We Th Fr Sa"
)

// FormatMonth создаёт календарь месяца по григорианскому календарю.
func FormatMonth(year int, month time.Month) (string, error) {
	if err := validateCalendarDate(year, month); err != nil {
		return "", err
	}

	first := time.Date(year, month, firstCalendarDay, 0, 0, 0, 0, time.UTC)
	lastDay := first.AddDate(0, 1, -1).Day()
	title := fmt.Sprintf("%s %d", month.String(), year)
	padding := (len(calendarWeekHeader) - len(title)) / 2

	var output strings.Builder
	fmt.Fprintf(&output, "%s%s\n", strings.Repeat(" ", padding), title)
	fmt.Fprintln(&output, calendarWeekHeader)

	week := strings.Repeat(" ", int(first.Weekday())*calendarDayWidth)
	for day := firstCalendarDay; day <= lastDay; day++ {
		week += fmt.Sprintf("%2d ", day)
		weekday := (int(first.Weekday()) + day - firstCalendarDay) % calendarDaysInWeek
		if weekday == int(time.Saturday) || day == lastDay {
			fmt.Fprintln(&output, strings.TrimRight(week, " "))
			week = ""
		}
	}
	return output.String(), nil
}

// FormatYear создаёт календари всех месяцев указанного года.
func FormatYear(year int) (string, error) {
	var output strings.Builder
	for month := time.January; month <= time.December; month++ {
		calendar, err := FormatMonth(year, month)
		if err != nil {
			return "", err
		}
		if month > time.January {
			output.WriteByte('\n')
		}
		output.WriteString(calendar)
	}
	return output.String(), nil
}

// validateCalendarDate проверяет допустимые значения года и месяца.
func validateCalendarDate(year int, month time.Month) error {
	if year < minCalendarYear || year > maxCalendarYear {
		return fmt.Errorf("год должен быть от %d до %d", minCalendarYear, maxCalendarYear)
	}
	if month < time.January || month > time.December {
		return fmt.Errorf("месяц должен быть от 1 до 12")
	}
	return nil
}
