package main

import (
	"os"

	"github.com/horvatic/vaticlang/pkg/interpreter"
)

func main() {
	path := os.Args[1]

	interpreter.Interpret(path)
}
