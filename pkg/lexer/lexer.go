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
		line := []rune(scanner.Text())
		for i := 0; i < len(line); i++ {
			if line[i] == '+' {
				tokens = append(tokens, token.NewToken(token.Plus, string(line[i])))
			} else if line[i] == '-' {
				tokens = append(tokens, token.NewToken(token.Subtract, string(line[i])))
			} else if isNum(line[i]) {
				tokens = append(tokens, token.NewToken(token.Number, string(line[i])))
			} else if line[i] == 'i' && i+1 < len(line) && i+2 < len(line) && i+3 < len(line) && line[i+1] == 'n' && line[i+2] == 't' && line[i+3] == ' ' {
				tokens = append(tokens, token.NewToken(token.Type, "int"))
				i += 4
				label := ""
				for i < len(line) {
					if line[i] == ';' || line[i] == ' ' {
						i--
						break
					}
					label += string(line[i])
					i++
				}
				tokens = append(tokens, token.NewToken(token.Label, label))
			} else if line[i] == 'o' && i+1 < len(line) && i+2 < len(line) && i+3 < len(line) && line[i+1] == 'u' && line[i+2] == 't' && line[i+3] == ' ' {
				tokens = append(tokens, token.NewOutToken())
				i += 4
				label := ""
				for i < len(line) {
					if line[i] == ';' || line[i] == ' ' {
						i--
						break
					}
					label += string(line[i])
					i++
				}
				tokens = append(tokens, token.NewToken(token.Label, label))
			} else if line[i] == ';' {
				tokens = append(tokens, token.NewEndCodeBlockToken())
			} else if line[i] == '=' {
				tokens = append(tokens, token.NewEqualToken())
			} else if line[i] == '/' {
				if i+1 < len(line) && line[i+1] == '/' {
					break
				}
			} else if line[i] == ' ' {
				continue
			} else {
				panic("undefined symbol: " + string(line[i]))
			}
		}
	}

	tokens = append(tokens, token.NewEOFToken())

	return tokens
}

func isNum(c rune) bool {
	return c == '0' || c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9'
}
