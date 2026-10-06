package main

import "flag"

// configuration хранит пути, заданные при запуске эмулятора.
type configuration struct {
	vfsPath    string
	scriptPath string
}

// parseConfig разбирает параметры запуска эмулятора.
func parseConfig() configuration {
	var cfg configuration

	flag.StringVar(&cfg.vfsPath, "vfs", "", "Путь к CSV-файлу VFS")
	flag.StringVar(&cfg.scriptPath, "script", "", "Путь к стартовому скрипту")
	flag.Parse()

	return cfg
}
