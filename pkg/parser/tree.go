package parser

import (
	"github.com/horvatic/vaticlang/pkg/token"
)

type SyntaxTree struct {
	RootNode *Node
}

func BuildTree() *SyntaxTree {
	return &SyntaxTree{
		RootNode: BuildNode(token.Start, nil),
	}
}
