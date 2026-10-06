package tests

import (
	"bytes"
	"testing"

	"github.com/marseliandrius/shell-emulator/src/commands"
	"github.com/marseliandrius/shell-emulator/src/vfs"
)

// TestReverseLines проверяет вывод строк в обратном порядке.
func TestReverseLines(t *testing.T) {
	got := commands.ReverseLines([]byte("one\ntwo\nthree\n"))
	want := "three\ntwo\none\n"

	if string(got) != want {
		t.Errorf("результат: %q, ожидалось: %q", got, want)
	}
}

// TestTacErrors проверяет ошибку при отсутствии аргументов.
func TestTacErrors(t *testing.T) {
	fs, err := vfs.Load("../data/nested.csv")
	if err != nil {
		t.Fatal(err)
	}

	shouldExit, err := commands.Execute(fs, "tac", nil)
	if err == nil {
		t.Error("ожидалась ошибка команды tac")
	}
	if shouldExit {
		t.Error("tac не должна завершать эмулятор")
	}
}

// TestReadFileCopy проверяет содержимое и независимость возвращаемой копии.
func TestReadFileCopy(t *testing.T) {
	fs, err := vfs.Load("../data/nested.csv")
	if err != nil {
		t.Fatal(err)
	}

	content, err := fs.ReadFile("/tmp/note.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := []byte("Hello\n")
	if !bytes.Equal(content, want) {
		t.Fatalf("содержимое: %q, ожидалось: %q", content, want)
	}
	content[0] = 'X'

	again, err := fs.ReadFile("/tmp/note.txt")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(again, want) {
		t.Fatal("изменение копии изменило файл в VFS")
	}
}
