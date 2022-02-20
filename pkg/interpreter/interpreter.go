package interpreter

import (
	"fmt"

	"github.com/horvatic/vaticlang/pkg/parser"
	"github.com/horvatic/vaticlang/pkg/token"
)

func Interpret(input string) {
	tree := parser.Parse(input)

	for _, n := range tree.RootNode.GetLeafs() {
		processNode(n)
	}
}

func processNode(node *parser.Node) {
	if node.GetTokenType() == token.Plus {
		fmt.Println(AddMode(node))
	} else if node.GetTokenType() == token.Subtract {
		fmt.Println(SubtractMode(node))
	} else if node.GetTokenType() == token.EOF {
		return
	}
}

func AddMode(node *parser.Node) int {
	total := 0
	for i, n := range node.GetLeafs() {
		if i == 0 {
			total = n.GetVal().(int)
		} else if n.GetTokenType() == token.Int {
			total += n.GetVal().(int)
		} else if n.GetTokenType() == token.Plus {
			total += AddMode(n)
		} else if n.GetTokenType() == token.Subtract {
			total -= SubtractMode(n)
		}
	}
	return total
}

func SubtractMode(node *parser.Node) int {
	total := 0
	for i, n := range node.GetLeafs() {
		if n.GetTokenType() == token.Int {
			if i == 0 {
				total = n.GetVal().(int)
			} else {
				total -= n.GetVal().(int)
			}
		} else if n.GetTokenType() == token.Plus {
			total += AddMode(n)
		} else if n.GetTokenType() == token.Subtract {
			total -= SubtractMode(n)
		}
	}
	return total
}
