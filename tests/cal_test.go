package tests

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/marseliandrius/shell-emulator/src/commands"
)

// TestFormatMonth проверяет расположение дней високосного февраля.
func TestFormatMonth(t *testing.T) {
	got, err := commands.FormatMonth(2024, time.February)
	if err != nil {
		t.Fatal(err)
	}

	want := "   February 2024\n" +
		"Su Mo Tu We Th Fr Sa\n" +
		"             1  2  3\n" +
		" 4  5  6  7  8  9 10\n" +
		"11 12 13 14 15 16 17\n" +
		"18 19 20 21 22 23 24\n" +
		"25 26 27 28 29\n"

	if got != want {
		t.Errorf("календарь:\n%s\nожидалось:\n%s", got, want)
	}
}

// TestFormatYear проверяет наличие всех месяцев в календаре года.
func TestFormatYear(t *testing.T) {
	got, err := commands.FormatYear(2024)
	if err != nil {
		t.Fatal(err)
	}

	for month := time.January; month <= time.December; month++ {
		title := fmt.Sprintf("%s 2024", month.String())
		if !strings.Contains(got, title) {
			t.Errorf("отсутствует месяц: %s", title)
		}
	}
}

// TestCalErrors проверяет отклонение недопустимого месяца.
func TestCalErrors(t *testing.T) {
	shouldExit, err := commands.Execute(nil, "cal", []string{"13", "2024"})
	if err == nil {
		t.Error("ожидалась ошибка команды cal")
	}
	if shouldExit {
		t.Error("cal не должна завершать эмулятор")
	}
}
