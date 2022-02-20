package interpreter

import (
	"strconv"

	"github.com/horvatic/vaticlang/pkg/parser"
	"github.com/horvatic/vaticlang/pkg/token"
)

func isMath(tokenType token.TokenType) bool {
	return tokenType == token.Plus || tokenType == token.Subtract
}

func mathMode(node *parser.Node) int {
	if node.GetTokenType() == token.Plus {
		return addMode(node)
	} else if node.GetTokenType() == token.Subtract {
		return subtractMode(node)
	}
	panic("no math modes")
}

func addMode(node *parser.Node) int {
	total := 0
	for i, n := range node.GetLeafs() {
		if i == 0 {
			total = convertToInt(n.GetVal())
		} else if n.GetTokenType() == token.Number {
			total += convertToInt(n.GetVal())
		} else if n.GetTokenType() == token.Plus {
			total += addMode(n)
		} else if n.GetTokenType() == token.Subtract {
			total -= subtractMode(n)
		}
	}
	return total
}

func subtractMode(node *parser.Node) int {
	total := 0
	for i, n := range node.GetLeafs() {
		if n.GetTokenType() == token.Number {
			if i == 0 {
				total = convertToInt(n.GetVal())
			} else {
				total -= convertToInt(n.GetVal())
			}
		} else if n.GetTokenType() == token.Plus {
			total += addMode(n)
		} else if n.GetTokenType() == token.Subtract {
			total -= subtractMode(n)
		}
	}
	return total
}

func convertToInt(val interface{}) int {
	n, err := strconv.Atoi(val.(string))
	if err != nil {
		panic(err)
	}
	return n
}
