package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"

	"github.com/marseliandrius/shell-emulator/src/commands"
	"github.com/marseliandrius/shell-emulator/src/parser"
	"github.com/marseliandrius/shell-emulator/src/vfs"
)

const emptyCommandLength = 0

// main разбирает параметры, загружает VFS и запускает эмулятор.
func main() {
	cfg := parseConfig()
	fmt.Printf("VFS: %q\n", cfg.vfsPath)
	fmt.Printf("Script: %q\n", cfg.scriptPath)

	fs, err := vfs.Load(cfg.vfsPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка загрузки VFS:", err)
		os.Exit(1)
	}
	fmt.Printf("Элементов VFS: %d\n", len(fs.Entries))

	prompt, err := makePrompt()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var runErr error
	if cfg.scriptPath != "" {
		runErr = runScript(cfg.scriptPath, prompt, fs)
	} else {
		runErr = runREPL(prompt, fs)
	}
	if runErr != nil {
		fmt.Fprintln(os.Stderr, "Ошибка выполнения:", runErr)
		os.Exit(1)
	}
}

// makePrompt создаёт приглашение на основе реальных данных ОС.
func makePrompt() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("не удалось определить пользователя: %w", err)
	}
	hostname, err := os.Hostname()
	if err != nil {
		return "", fmt.Errorf("не удалось определить имя компьютера: %w", err)
	}
	return fmt.Sprintf("%s@%s:~$ ", currentUser.Username, hostname), nil
}

// runREPL читает команды до exit или конца ввода.
// При ошибке чтения возвращает её вызывающей функции.
func runREPL(prompt string, fs *vfs.FileSystem) error {
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
		shouldExit, err := commands.Execute(fs, parts[0], parts[1:])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка команды:", err)
			continue
		}
		if shouldExit {
			return nil
		}
	}
}
