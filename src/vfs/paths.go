package vfs

import (
	"errors"
	"fmt"
	"path"
	"strings"
)

// Resolve преобразует путь в абсолютный путь внутри VFS.
// Относительный путь отсчитывается от CurrentDir.
// Существование элемента не проверяется.
func (fs *FileSystem) Resolve(name string) (string, error) {
	if name == "" {
		return "", errors.New("пустой путь")
	}
	if strings.ContainsAny(name, "\x00\r\n") {
		return "", errors.New("недопустимые символы в пути")
	}
	if path.IsAbs(name) {
		return path.Clean(name), nil
	}
	return path.Join(fs.CurrentDir, name), nil
}

// ChangeDir меняет текущий каталог после проверки виртуального пути.
// При ошибке текущий каталог остаётся прежним.
func (fs *FileSystem) ChangeDir(name string) error {
	target, err := fs.Resolve(name)
	if err != nil {
		return err
	}

	entry, exists := fs.Entries[target]
	if !exists {
		return fmt.Errorf("путь не существует: %s", target)
	}
	if !entry.IsDir {
		return fmt.Errorf("не является каталогом: %s", target)
	}

	fs.CurrentDir = target
	return nil
}
