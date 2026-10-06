package vfs

import (
	"encoding/base64"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
)

const (
	pathColumn = iota
	typeColumn
	contentColumn
	columnCount
)

const (
	rootPath        = "/"
	directoryType   = "directory"
	fileType        = "file"
	emptyCount      = 0
	firstDataRecord = 2
)

// Entry описывает файл или папку виртуальной файловой системы.
type Entry struct {
	Path    string
	IsDir   bool
	Content []byte
}

// FileSystem хранит элементы виртуальной файловой системы в памяти.
type FileSystem struct {
	Entries map[string]Entry
}

// Load загружает виртуальную файловую систему из CSV-файла.
func Load(filename string) (*FileSystem, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть VFS: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = columnCount
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения CSV: %w", err)
	}

	return parseRecords(records)
}

// parseRecords проверяет заголовок и создаёт файловую систему из записей CSV.
func parseRecords(records [][]string) (*FileSystem, error) {
	if len(records) == emptyCount {
		return nil, errors.New("CSV-файл пуст")
	}
	if strings.Join(records[0], ",") != "path,type,content" {
		return nil, errors.New("ожидался заголовок path,type,content")
	}

	fs := &FileSystem{Entries: make(map[string]Entry)}
	for index, record := range records[1:] {
		entry, err := parseRecord(record)
		if err != nil {
			return nil, fmt.Errorf("запись %d: %w", index+firstDataRecord, err)
		}
		if _, exists := fs.Entries[entry.Path]; exists {
			return nil, fmt.Errorf("повторяющийся путь: %q", entry.Path)
		}
		fs.Entries[entry.Path] = entry
	}

	if err := validateTree(fs); err != nil {
		return nil, err
	}
	return fs, nil
}

// parseRecord проверяет запись и декодирует содержимое файла.
func parseRecord(record []string) (Entry, error) {
	entry := Entry{Path: record[pathColumn]}
	if err := validatePath(entry.Path); err != nil {
		return Entry{}, err
	}

	switch record[typeColumn] {
	case directoryType:
		if record[contentColumn] != "" {
			return Entry{}, errors.New("содержимое папки должно быть пустым")
		}
		entry.IsDir = true
	case fileType:
		content, err := base64.StdEncoding.DecodeString(record[contentColumn])
		if err != nil {
			return Entry{}, fmt.Errorf("неверный Base64: %w", err)
		}
		entry.Content = content
	default:
		return Entry{}, fmt.Errorf("неизвестный тип: %q", record[typeColumn])
	}

	return entry, nil
}

// validatePath проверяет абсолютный путь в формате UNIX.
func validatePath(name string) error {
	if !path.IsAbs(name) || path.Clean(name) != name {
		return fmt.Errorf("некорректный виртуальный путь: %q", name)
	}
	if strings.ContainsAny(name, "\x00\r\n") {
		return fmt.Errorf("недопустимые символы в пути: %q", name)
	}
	return nil
}

// validateTree проверяет корень и наличие родительских папок.
func validateTree(fs *FileSystem) error {
	root, exists := fs.Entries[rootPath]
	if !exists || !root.IsDir {
		return errors.New("VFS должна содержать корневую папку /")
	}

	for name := range fs.Entries {
		if name == rootPath {
			continue
		}
		parent, exists := fs.Entries[path.Dir(name)]
		if !exists {
			return fmt.Errorf("для %q отсутствует родительская папка", name)
		}
		if !parent.IsDir {
			return fmt.Errorf("родитель элемента %q не является папкой", name)
		}
	}
	return nil
}
