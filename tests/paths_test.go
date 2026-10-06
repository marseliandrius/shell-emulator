package tests

import (
	"testing"

	"github.com/marseliandrius/shell-emulator/src/vfs"
)

// TestVFSInitialDirectory проверяет текущий каталог после загрузки VFS.
func TestVFSInitialDirectory(t *testing.T) {
	fs, err := vfs.Load("../data/nested.csv")
	if err != nil {
		t.Fatal(err)
	}
	if fs.CurrentDir != "/" {
		t.Fatalf("текущий каталог: %q, ожидалось: /", fs.CurrentDir)
	}
}

// TestResolve проверяет преобразование относительного виртуального пути.
func TestResolve(t *testing.T) {
	fs := &vfs.FileSystem{CurrentDir: "/home/user"}

	got, err := fs.Resolve("../user/./docs")
	if err != nil {
		t.Fatal(err)
	}
	if got != "/home/user/docs" {
		t.Errorf("путь: %q, ожидалось: /home/user/docs", got)
	}
}

// TestResolveInvalid проверяет отклонение пустого пути.
func TestResolveInvalid(t *testing.T) {
	fs := &vfs.FileSystem{CurrentDir: "/"}

	if _, err := fs.Resolve(""); err == nil {
		t.Fatal("ожидалась ошибка пути")
	}
}
