package commands

import (
	"fmt"
	"strconv"
	"time"
)

const (
	singleCalendarArgument = 1
	monthYearArguments     = 2
)

// calendarRequest хранит параметры запрошенного календаря.
type calendarRequest struct {
	year      int
	month     time.Month
	wholeYear bool
}

// showCalendar разбирает аргументы cal и выводит календарь.
func showCalendar(args []string) error {
	request, err := parseCalendarArgs(args, time.Now())
	if err != nil {
		return fmt.Errorf("cal: %w", err)
	}

	var calendar string
	if request.wholeYear {
		calendar, err = FormatYear(request.year)
	} else {
		calendar, err = FormatMonth(request.year, request.month)
	}
	if err != nil {
		return fmt.Errorf("cal: %w", err)
	}

	fmt.Print(calendar)
	return nil
}

// parseCalendarArgs получает год, месяц и режим из аргументов команды.
func parseCalendarArgs(args []string, now time.Time) (calendarRequest, error) {
	request := calendarRequest{year: now.Year(), month: now.Month()}
	yearText := ""

	switch len(args) {
	case noArguments:
		return request, nil
	case singleCalendarArgument:
		request.wholeYear = true
		yearText = args[0]
	case monthYearArguments:
		monthNumber, err := strconv.Atoi(args[0])
		if err != nil {
			return request, fmt.Errorf("неверный месяц: %q", args[0])
		}
		request.month = time.Month(monthNumber)
		yearText = args[1]
	default:
		return request, fmt.Errorf("ожидается cal, cal год или cal месяц год")
	}

	year, err := strconv.Atoi(yearText)
	if err != nil {
		return request, fmt.Errorf("неверный год: %q", yearText)
	}
	request.year = year
	return request, nil
}
