package tests

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/marseliandrius/shell-emulator/src/vfs"
)

const (
	vfsHeader = "path,type,content\n"
	vfsRoot   = "/,directory,\n"
)

// TestLoadValid проверяет загрузку трёх вариантов VFS и содержимое файлов.
func TestLoadValid(t *testing.T) {
	cases := []struct {
		filename    string
		wantCount   int
		filePath    string
		wantContent string
	}{
		{"minimal.csv", 1, "", ""},
		{"vfs.csv", 4, "/lines.txt", "one\ntwo\nthree\n"},
		{"nested.csv", 7, "/home/user/docs/lines.txt", "one\ntwo\nthree\n"},
	}

	for _, tc := range cases {
		t.Run(tc.filename, func(t *testing.T) {
			fs, err := vfs.Load(filepath.Join("..", "data", tc.filename))
			if err != nil {
				t.Fatal(err)
			}
			if len(fs.Entries) != tc.wantCount {
				t.Errorf("элементов: %d, ожидалось: %d", len(fs.Entries), tc.wantCount)
			}
			root, exists := fs.Entries["/"]
			if !exists || !root.IsDir {
				t.Fatal("корневая папка отсутствует")
			}
			if tc.filePath != "" {
				entry := requireVFSFile(t, fs, tc.filePath)
				if string(entry.Content) != tc.wantContent {
					t.Errorf("содержимое: %q, ожидалось: %q", entry.Content, tc.wantContent)
				}
			}
		})
	}
}

// TestLoadInvalid проверяет ошибки заголовка CSV и кодирования Base64.
func TestLoadInvalid(t *testing.T) {
	cases := []struct {
		name    string
		content string
	}{
		{"неверный заголовок", "name,type,content\n" + vfsRoot},
		{"неверный Base64", vfsHeader + vfsRoot + "/a,file,!\n"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			filename := writeVFSFixture(t, tc.content)
			if _, err := vfs.Load(filename); err == nil {
				t.Fatal("ожидалась ошибка формата VFS")
			}
		})
	}
}

// TestLoadMissing проверяет ошибку отсутствующего источника VFS.
func TestLoadMissing(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "missing.csv")
	_, err := vfs.Load(filename)
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("ожидалась ошибка отсутствующего файла, получено: %v", err)
	}
}

// TestVFSMemory проверяет двоичные данные и сохранность источника VFS.
func TestVFSMemory(t *testing.T) {
	source := vfsHeader + vfsRoot + "/data.bin,file,AP8BAg==\n"
	filename := writeVFSFixture(t, source)
	fs, err := vfs.Load(filename)
	if err != nil {
		t.Fatal(err)
	}

	entry := requireVFSFile(t, fs, "/data.bin")
	if !bytes.Equal(entry.Content, []byte{0, 255, 1, 2}) {
		t.Fatal("двоичное содержимое декодировано неверно")
	}
	entry.Content[0] = 99
	delete(fs.Entries, "/data.bin")

	after, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(after, []byte(source)) {
		t.Fatal("изменение VFS в памяти изменило исходный CSV")
	}
}

// writeVFSFixture создаёт CSV во временной папке теста.
func writeVFSFixture(t *testing.T, content string) string {
	t.Helper()
	filename := filepath.Join(t.TempDir(), "vfs.csv")
	if err := os.WriteFile(filename, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
	return filename
}

// requireVFSFile возвращает файл или завершает текущий тест с ошибкой.
func requireVFSFile(t *testing.T, fs *vfs.FileSystem, name string) vfs.Entry {
	t.Helper()
	entry, exists := fs.Entries[name]
	if !exists || entry.IsDir {
		t.Fatalf("файл %q отсутствует или является папкой", name)
	}
	return entry
}
