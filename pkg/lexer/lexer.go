package lexer

import (
	"bufio"
	"os"

	"github.com/horvatic/vaticlang/pkg/token"
)

func BuildTokens(input string) []*token.Token {
	var tokens []*token.Token

	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for scanner.Scan() {
		for _, c := range scanner.Text() {
			if c == '+' {
				tokens = append(tokens, token.NewToken(token.Plus, c))
			} else if c == '-' {
				tokens = append(tokens, token.NewToken(token.Subtract, c))
			} else if isNum(c) {
				tokens = append(tokens, token.NewToken(token.Int, c))
			} else if c == ';' {
				tokens = append(tokens, token.NewEndCodeBlockToken())
			} else {
				panic("undefined symbol")
			}
		}
	}

	tokens = append(tokens, token.NewEOFToken())

	return tokens
}

func isNum(c rune) bool {
	return c == '0' || c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9'
}
