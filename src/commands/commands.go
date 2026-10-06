package commands

import (
	"fmt"

	"github.com/marseliandrius/shell-emulator/src/vfs"
)

const noArguments = 0

// Execute выполняет команду с указанными аргументами.
// Возвращает признак завершения эмулятора и возможную ошибку.
func Execute(fs *vfs.FileSystem, name string, args []string) (bool, error) {
	switch name {
	case "ls", "cd":
		fmt.Printf("%s: аргументы %q\n", name, args)
		return false, nil
	case "exit":
		if len(args) != noArguments {
			return false, fmt.Errorf("exit не принимает аргументы")
		} else {
			return true, nil
		}
	default:
		return false, fmt.Errorf("неизвестная команда: %s", name)
	}
}
