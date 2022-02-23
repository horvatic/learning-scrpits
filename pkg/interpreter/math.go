package interpreter

import (
	"strconv"

	"github.com/horvatic/vaticlang/pkg/parser"
	"github.com/horvatic/vaticlang/pkg/token"
)

func isMath(tokenType token.TokenType) bool {
	return tokenType == token.Plus || tokenType == token.Subtract
}

func mathMode(node *parser.Node, dataStore *DataStore) int {
	if node.GetTokenType() == token.Plus {
		return addMode(node, dataStore)
	} else if node.GetTokenType() == token.Subtract {
		return subtractMode(node, dataStore)
	}
	panic("no math modes")
}

func addMode(node *parser.Node, dataStore *DataStore) int {
	total := 0
	for i, n := range node.GetLeafs() {
		if n.GetTokenType() == token.Number || n.GetTokenType() == token.Label {
			if i == 0 {
				if n.GetTokenType() == token.Number {
					total = convertToInt(n.GetVal())
				} else if n.GetTokenType() == token.Label {
					total = convertToInt(dataStore.GetData(n.GetVal().(string)))
				} else {
					panic("unknown symbol")
				}
			} else {
				if n.GetTokenType() == token.Number {
					total += convertToInt(n.GetVal())
				} else if n.GetTokenType() == token.Label {
					total += convertToInt(dataStore.GetData(n.GetVal().(string)))
				} else {
					panic("unknown symbol")
				}
			}
		} else if n.GetTokenType() == token.Plus {
			total += addMode(n, dataStore)
		} else if n.GetTokenType() == token.Subtract {
			total -= subtractMode(n, dataStore)
		}
	}
	return total
}

func subtractMode(node *parser.Node, dataStore *DataStore) int {
	total := 0
	for i, n := range node.GetLeafs() {
		if n.GetTokenType() == token.Number || n.GetTokenType() == token.Label {
			if i == 0 {
				if n.GetTokenType() == token.Number {
					total = convertToInt(n.GetVal())
				} else if n.GetTokenType() == token.Label {
					total = convertToInt(dataStore.GetData(n.GetVal().(string)))
				} else {
					panic("unknown symbol")
				}
			} else {
				if n.GetTokenType() == token.Number {
					total -= convertToInt(n.GetVal())
				} else if n.GetTokenType() == token.Label {
					total -= convertToInt(dataStore.GetData(n.GetVal().(string)))
				} else {
					panic("unknown symbol")
				}
			}
		} else if n.GetTokenType() == token.Plus {
			total += addMode(n, dataStore)
		} else if n.GetTokenType() == token.Subtract {
			total -= subtractMode(n, dataStore)
		}
	}
	return total
}

func convertToInt(val interface{}) int {
	if i, ok := val.(int); ok {
		return i
	}
	n, err := strconv.Atoi(val.(string))
	if err != nil {
		panic(err)
	}
	return n
}
