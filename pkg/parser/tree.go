package parser

import (
	"github.com/horvatic/vaticlang/pkg/token"
)

type SyntaxTree struct {
	RootNode *Node
}

type Node struct {
	tokeType token.TokenType
	val      interface{}
	nodes    []*Node
}

func BuildTree() *SyntaxTree {
	return &SyntaxTree{
		RootNode: BuildNode(token.Start, nil),
	}
}

func BuildNode(tokeType token.TokenType, val interface{}) *Node {
	return &Node{
		tokeType: tokeType,
		val:      val,
		nodes:    nil,
	}
}

func (n *Node) AddLeaf(node *Node) {
	n.nodes = append(n.nodes, node)
}

func (n *Node) GetTokenType() token.TokenType {
	return n.tokeType
}

func (n *Node) GetVal() interface{} {
	return n.val
}

func (n *Node) GetLeafs() []*Node {
	return n.nodes
}
