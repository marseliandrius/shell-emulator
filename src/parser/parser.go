package parser

import (
	"errors"
	"strings"
	"unicode"
)

const noQuote rune = 0

// Parse разбирает строку на слова с учётом одинарных и двойных кавычек.
// Возвращает ошибку, если кавычки не закрыты.
func Parse(line string) ([]string, error) {
	var args []string
	var token strings.Builder
	quote := noQuote
	started := false
	for _, ch := range line {
		switch {
		case quote != noQuote:
			if ch == quote {
				quote = noQuote
			} else {
				token.WriteRune(ch)
			}
		case ch == '\'' || ch == '"':
			quote = ch
			started = true
		case unicode.IsSpace(ch):
			if started {
				args = append(args, token.String())
				token.Reset()
				started = false
			}
		default:
			token.WriteRune(ch)
			started = true
		}
	}
	if quote != noQuote {
		return nil, errors.New("незакрытые кавычки")
	}
	if started {
		args = append(args, token.String())
	}
	return args, nil
}
