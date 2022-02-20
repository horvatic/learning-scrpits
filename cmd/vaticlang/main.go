package main

import (
	"os"

	"github.com/horvatic/vaticlang/pkg/interpreter"
)

func main() {
	path := os.Args[1]

	interpret := interpreter.NewInterpreter(interpreter.NewDataStore())
	interpret.Interpret(path)
}
