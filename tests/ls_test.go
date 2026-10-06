package tests

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/marseliandrius/shell-emulator/src/vfs"
)

// listCase описывает запрос к VFS и ожидаемый список имён.
type listCase struct {
	filename   string
	currentDir string
	name       string
	want       []string
	wantErr    bool
}

// TestList проверяет просмотр файлов, каталогов и ошибки путей.
func TestList(t *testing.T) {
	cases := []listCase{
		{"vfs.csv", "/", "/", []string{"empty/", "hello.txt", "lines.txt"}, false},
		{"nested.csv", "/", "/tmp/note.txt", []string{"note.txt"}, false},
		{"nested.csv", "/", "/missing", nil, true},
	}

	for _, tc := range cases {
		t.Run(tc.filename+" "+tc.name, func(t *testing.T) {
			checkListCase(t, tc)
		})
	}
}

// checkListCase проверяет список имён и неизменность текущего каталога.
func checkListCase(t *testing.T, tc listCase) {
	t.Helper()
	fs, err := vfs.Load(filepath.Join("..", "data", tc.filename))
	if err != nil {
		t.Fatal(err)
	}
	fs.CurrentDir = tc.currentDir

	got, err := fs.List(tc.name)
	if (err != nil) != tc.wantErr {
		t.Fatalf("ошибка: %v, ожидалась ошибка: %t", err, tc.wantErr)
	}
	if !reflect.DeepEqual(got, tc.want) {
		t.Errorf("список: %q, ожидалось: %q", got, tc.want)
	}
	if fs.CurrentDir != tc.currentDir {
		t.Error("просмотр содержимого изменил текущий каталог")
	}
}
