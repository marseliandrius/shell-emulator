package commands

import (
	"fmt"

	"github.com/marseliandrius/shell-emulator/src/vfs"
)

const singleLSTarget = 1

// listEntries выводит файлы и содержимое виртуальных каталогов.
// Без аргументов используется текущий каталог.
func listEntries(fs *vfs.FileSystem, args []string) error {
	if len(args) == noArguments {
		args = []string{"."}
	}

	for _, name := range args {
		names, err := fs.List(name)
		if err != nil {
			return fmt.Errorf("ls: %w", err)
		}

		if len(args) > singleLSTarget {
			fmt.Printf("%s:\n", name)
		}
		for _, entryName := range names {
			fmt.Println(entryName)
		}
	}
	return nil
}
