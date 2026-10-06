package commands

import (
	"fmt"

	"github.com/marseliandrius/shell-emulator/src/vfs"
)

const maxCDArguments = 1

// changeDirectory обрабатывает аргументы cd и меняет виртуальный каталог.
// Без аргументов команда переходит в корневой каталог.
func changeDirectory(fs *vfs.FileSystem, args []string) error {
	if len(args) > maxCDArguments {
		return fmt.Errorf("cd принимает не более одного аргумента")
	}

	target := "/"
	if len(args) == maxCDArguments {
		target = args[0]
	}

	if err := fs.ChangeDir(target); err != nil {
		return fmt.Errorf("cd: %w", err)
	}
	return nil
}
