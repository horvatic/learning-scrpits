package interpreter

import (
	"fmt"

	"github.com/horvatic/vaticlang/pkg/parser"
	"github.com/horvatic/vaticlang/pkg/token"
)

type Interpreter struct {
	dataStore *DataStore
}

func NewInterpreter(dataStore *DataStore) *Interpreter {
	return &Interpreter{
		dataStore: dataStore,
	}
}

func (interpreter *Interpreter) Interpret(input string) {
	tree := parser.Parse(input)

	for _, n := range tree.RootNode.GetLeafs() {
		interpreter.processNode(n)
	}
}

func (interpreter *Interpreter) processNode(node *parser.Node) {
	if node.GetTokenType() == token.Type {
		if node.GetLeafs()[0].GetTokenType() == token.Equal {
			context := node.GetLeafs()[0].GetLeafs()[0].GetVal().(string)
			if isMath(node.GetLeafs()[0].GetLeafs()[1].GetTokenType()) {
				interpreter.dataStore.AddData(context, mathMode(node.GetLeafs()[0].GetLeafs()[1], interpreter.dataStore))
			} else if node.GetLeafs()[0].GetLeafs()[1].GetTokenType() == token.Number {
				interpreter.dataStore.AddData(context, node.GetLeafs()[0].GetLeafs()[1].GetVal())
			}
		} else if node.GetLeafs()[0].GetTokenType() == token.Label {
			interpreter.dataStore.AddData(node.GetLeafs()[0].GetVal().(string), nil)
		} else {
			panic("unknown symbol")
		}
	} else if node.GetTokenType() == token.Label {
		if node.GetLeafs()[0].GetTokenType() == token.Equal {
			if node.GetLeafs()[0].GetLeafs()[0].GetTokenType() == token.Number {
				interpreter.dataStore.AddData(node.GetVal().(string), node.GetLeafs()[0].GetLeafs()[0].GetVal())
			} else if node.GetLeafs()[0].GetLeafs()[0].GetTokenType() == token.Label {
				interpreter.dataStore.AddData(node.GetVal().(string), interpreter.dataStore.GetData(node.GetLeafs()[0].GetLeafs()[0].GetVal().(string)))
			} else if isMath(node.GetLeafs()[0].GetLeafs()[0].GetTokenType()) {
				interpreter.dataStore.AddData(node.GetVal().(string), mathMode(node.GetLeafs()[0].GetLeafs()[0], interpreter.dataStore))
			} else {
				panic("unknown symbol")
			}
		} else {
			panic("unknown symbol")
		}
	} else if node.GetTokenType() == token.Out {
		context := node.GetLeafs()[0].GetVal().(string)
		fmt.Println(interpreter.dataStore.GetData(context))
	}

	if node.GetTokenType() == token.EOF {
		return
	}
}
