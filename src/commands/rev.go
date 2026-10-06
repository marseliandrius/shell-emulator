package commands

import (
	"fmt"
	"slices"
	"strings"
	"unicode/utf8"

	"github.com/marseliandrius/shell-emulator/src/vfs"
)

// reverseFileCharacters разворачивает символы строк виртуальных файлов.
func reverseFileCharacters(fs *vfs.FileSystem, args []string) error {
	if len(args) == noArguments {
		return fmt.Errorf("rev требует хотя бы один путь к файлу")
	}

	for _, name := range args {
		content, err := fs.ReadFile(name)
		if err != nil {
			return fmt.Errorf("rev: %w", err)
		}
		result, err := ReverseText(content)
		if err != nil {
			return fmt.Errorf("rev: %w", err)
		}
		fmt.Print(string(result))
	}
	return nil
}

// ReverseText разворачивает кодовые точки каждой строки корректного UTF-8.
// Порядок строк и разделители LF и CRLF сохраняются.
func ReverseText(content []byte) ([]byte, error) {
	if !utf8.Valid(content) {
		return nil, fmt.Errorf("содержимое не является корректным UTF-8")
	}

	var output strings.Builder
	for _, line := range strings.SplitAfter(string(content), "\n") {
		output.WriteString(reverseTextLine(line))
	}
	return []byte(output.String()), nil
}

// reverseTextLine разворачивает строку, сохраняя её перевод строки.
func reverseTextLine(line string) string {
	ending := ""
	if strings.HasSuffix(line, "\n") {
		ending = "\n"
		line = strings.TrimSuffix(line, "\n")
		if strings.HasSuffix(line, "\r") {
			ending = "\r\n"
			line = strings.TrimSuffix(line, "\r")
		}
	}

	characters := []rune(line)
	slices.Reverse(characters)
	return string(characters) + ending
}
