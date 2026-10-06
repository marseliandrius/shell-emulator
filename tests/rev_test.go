package tests

import (
	"testing"

	"github.com/marseliandrius/shell-emulator/src/commands"
	"github.com/marseliandrius/shell-emulator/src/vfs"
)

// TestReverseText проверяет разворот Unicode и сохранение порядка строк.
func TestReverseText(t *testing.T) {
	got, err := commands.ReverseText([]byte("Привет\nHello\n"))
	if err != nil {
		t.Fatal(err)
	}
	want := "тевирП\nolleH\n"

	if string(got) != want {
		t.Errorf("результат: %q, ожидалось: %q", got, want)
	}
}

// TestReverseTextInvalid проверяет отклонение неверного UTF-8.
func TestReverseTextInvalid(t *testing.T) {
	if _, err := commands.ReverseText([]byte{0xff, 0xfe}); err == nil {
		t.Fatal("ожидалась ошибка UTF-8")
	}
}

// TestRevErrors проверяет ошибку при передаче каталога.
func TestRevErrors(t *testing.T) {
	fs, err := vfs.Load("../data/nested.csv")
	if err != nil {
		t.Fatal(err)
	}

	shouldExit, err := commands.Execute(fs, "rev", []string{"/home"})
	if err == nil {
		t.Error("ожидалась ошибка команды rev")
	}
	if shouldExit {
		t.Error("rev не должна завершать эмулятор")
	}
}
