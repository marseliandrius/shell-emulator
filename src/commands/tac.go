package commands

import (
	"bytes"
	"fmt"
	"slices"

	"github.com/marseliandrius/shell-emulator/src/vfs"
)

// reverseFileLines выводит строки каждого виртуального файла в обратном порядке.
func reverseFileLines(fs *vfs.FileSystem, args []string) error {
	if len(args) == noArguments {
		return fmt.Errorf("tac требует хотя бы один путь к файлу")
	}

	for _, name := range args {
		content, err := fs.ReadFile(name)
		if err != nil {
			return fmt.Errorf("tac: %w", err)
		}
		fmt.Print(string(ReverseLines(content)))
	}
	return nil
}

// ReverseLines меняет порядок строк, сохраняя их исходные разделители.
func ReverseLines(content []byte) []byte {
	lines := bytes.SplitAfter(content, []byte("\n"))
	slices.Reverse(lines)
	return bytes.Join(lines, nil)
}
