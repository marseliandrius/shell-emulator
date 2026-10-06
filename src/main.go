package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"

	"github.com/marseliandrius/shell-emulator/src/commands"
	"github.com/marseliandrius/shell-emulator/src/parser"
)

const emptyCommandLength = 0

// main разбирает параметры и запускает эмулятор.
func main() {
	cfg := parseConfig()
	fmt.Printf("VFS: %q\n", cfg.vfsPath)
	fmt.Printf("Script: %q\n", cfg.scriptPath)

	currentUser, err := user.Current()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при определении текущего пользователя:", err)
		os.Exit(1)
	}
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при определении имени компьютера:", err)
		os.Exit(1)
	}
	prompt := fmt.Sprintf("%s@%s:~$ ", currentUser.Username, hostname)
	var runErr error
	if cfg.scriptPath != "" {
		runErr = runScript(cfg.scriptPath, prompt)
	} else {
		runErr = runREPL(prompt)
	}
	if runErr != nil {
		fmt.Fprintln(os.Stderr, "Ошибка выполнения:", runErr)
		os.Exit(1)
	}
}

// runREPL выводит приглашение и читает строки до команды exit или конца ввода.
// При ошибке чтения возвращает её вызывающей функции.
func runREPL(prompt string) error {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		check := scanner.Scan()
		if !check {
			return scanner.Err()
		}
		line := scanner.Text()
		parts, err := parser.Parse(line)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка разбора:", err)
			continue
		}
		if len(parts) == emptyCommandLength {
			continue
		}
		shouldExit, err := commands.Execute(parts[0], parts[1:])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка команды:", err)
			continue
		}
		if shouldExit {
			return nil
		}
	}
}
