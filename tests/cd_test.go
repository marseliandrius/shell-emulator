package tests

import (
	"testing"

	"github.com/marseliandrius/shell-emulator/src/commands"
	"github.com/marseliandrius/shell-emulator/src/vfs"
)

// cdCase описывает аргументы команды и ожидаемое состояние VFS.
type cdCase struct {
	name    string
	args    []string
	wantDir string
	wantErr bool
}

// TestCD проверяет переходы по каталогам и ошибки команды cd.
func TestCD(t *testing.T) {
	cases := []cdCase{
		{"без аргументов", nil, "/", false},
		{"относительный путь", []string{"docs"}, "/home/user/docs", false},
		{"путь к файлу", []string{"docs/lines.txt"}, "/home/user", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			checkCDCase(t, tc)
		})
	}
}

// checkCDCase проверяет результат команды и сохранность каталога при ошибке.
func checkCDCase(t *testing.T, tc cdCase) {
	t.Helper()
	fs, err := vfs.Load("../data/nested.csv")
	if err != nil {
		t.Fatal(err)
	}
	fs.CurrentDir = "/home/user"

	shouldExit, err := commands.Execute(fs, "cd", tc.args)
	if shouldExit {
		t.Error("cd не должна завершать эмулятор")
	}
	if (err != nil) != tc.wantErr {
		t.Fatalf("ошибка: %v, ожидалась ошибка: %t", err, tc.wantErr)
	}
	if fs.CurrentDir != tc.wantDir {
		t.Errorf("каталог: %q, ожидалось: %q", fs.CurrentDir, tc.wantDir)
	}
}
