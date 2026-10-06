package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/marseliandrius/shell-emulator/src/commands"
	"github.com/marseliandrius/shell-emulator/src/parser"
)

// runScript выполняет команды файла, останавливаясь при первой ошибке или exit.
func runScript(path, prompt string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("не удалось открыть стартовый скрипт: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		fmt.Printf("%s%s\n", prompt, line)
		parts, err := parser.Parse(line)
		if err != nil {
			return fmt.Errorf("строка %d: %w", lineNumber, err)
		}
		if len(parts) == emptyCommandLength {
			continue
		}
		shouldExit, err := commands.Execute(parts[0], parts[1:])
		if err != nil {
			return fmt.Errorf("строка %d: %w", lineNumber, err)
		}
		if shouldExit {
			return nil
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка чтения стартового скрипта: %w", err)
	}
	return nil
}
