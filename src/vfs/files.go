package vfs

import (
	"bytes"
	"fmt"
)

// ReadFile возвращает копию содержимого виртуального файла.
// Относительный путь отсчитывается от текущего каталога.
func (fs *FileSystem) ReadFile(name string) ([]byte, error) {
	target, err := fs.Resolve(name)
	if err != nil {
		return nil, err
	}

	entry, exists := fs.Entries[target]
	if !exists {
		return nil, fmt.Errorf("путь не существует: %s", target)
	}
	if entry.IsDir {
		return nil, fmt.Errorf("является каталогом: %s", target)
	}

	return bytes.Clone(entry.Content), nil
}
