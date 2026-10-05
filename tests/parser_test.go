package tests

import (
	"reflect"
	"testing"

	"github.com/marseliandrius/shell-emulator/src/parser"
)

// TestParse проверяет разбор аргументов и обнаружение незакрытых кавычек.
func TestParse(t *testing.T) {
	cases := []struct {
		name    string
		input   string
		want    []string
		wantErr bool
	}{
		{"обычные аргументы", "ls folder", []string{"ls", "folder"}, false},
		{"двойные кавычки", `cd "my folder"`, []string{"cd", "my folder"}, false},
		{"одинарные кавычки", "cd 'моя папка'", []string{"cd", "моя папка"}, false},
		{"пустой аргумент", `ls ""`, []string{"ls", ""}, false},
		{"незакрытые кавычки", `cd "folder`, nil, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parser.Parse(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Fatal("ожидалась ошибка, но получен nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("неожиданная ошибка: %v", err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("получено %q, ожидалось %q", got, tc.want)
			}
		})
	}
}
