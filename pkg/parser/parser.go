package parser

import (
	"github.com/horvatic/vaticlang/pkg/lexer"
	"github.com/horvatic/vaticlang/pkg/token"
)

func Parse(input string) *SyntaxTree {
	tokens := lexer.BuildTokens(input)
	tree := BuildTree()

	var codeBlock []*token.Token
	for _, t := range tokens {
		if t.GetTokenType() == token.EndCodeBlock {
			tree.RootNode.AddLeaf(buildBlock(codeBlock))
			codeBlock = nil
		} else {
			codeBlock = append(codeBlock, t)
		}
	}

	return tree
}

func buildBlock(codeBlock []*token.Token) *Node {
	var nodes []*Node
	for i := 0; i < len(codeBlock); {
		if codeBlock[i].GetTokenType() == token.Number {
			buildingNum := true
			num := ""
			for buildingNum {
				num += string(codeBlock[i].GetRawVal())
				i++
				if i >= len(codeBlock) || codeBlock[i].GetTokenType() != token.Number {
					buildingNum = false
				}
			}
			nodes = append(nodes, BuildNode(token.Number, num))
		} else if codeBlock[i].GetTokenType() == token.Label {
			nodes = append(nodes, BuildNode(token.Label, codeBlock[i].GetRawVal()))
			i++
		} else if codeBlock[i].GetTokenType() == token.Type {
			nodes = append(nodes, BuildNode(token.Type, codeBlock[i].GetRawVal()))
			i++
		} else {
			nodes = append(nodes, BuildNode(codeBlock[i].GetTokenType(), nil))
			i++
		}
	}
	return linkNodes(nodes)
}

func linkNodes(nodes []*Node) *Node {
	root := nodes[0]
	if len(nodes) == 1 {
		return root
	}
	if root.tokeType == token.Type {
		if len(nodes) == 2 {
			if nodes[1].GetTokenType() == token.Label {
				root.AddLeaf(nodes[1])
			} else {
				panic("unknown symbol")
			}
		} else if nodes[2].GetTokenType() == token.Equal {
			root.AddLeaf(nodes[2])
			nodes[2].AddLeaf(nodes[1])
			if len(nodes) == 4 && nodes[3].GetTokenType() == token.Number {
				nodes[2].AddLeaf(nodes[3])
			} else {
				nodes[2].AddLeaf(linkMathNodes(nodes[3:]))
			}
		} else {
			panic("expected equals symbol")
		}
	} else if root.tokeType == token.Label {
		if len(nodes) == 3 {
			if (nodes[2].GetTokenType() == token.Label || nodes[2].GetTokenType() == token.Number) && nodes[1].GetTokenType() == token.Equal {
				root.AddLeaf(nodes[1])
				nodes[1].AddLeaf(nodes[2])
			} else {
				panic("unknown symbol")
			}
		} else if len(nodes) == 5 {
			if (nodes[2].GetTokenType() == token.Label || nodes[2].GetTokenType() == token.Number) && nodes[1].GetTokenType() == token.Equal {
				root.AddLeaf(nodes[1])
				nodes[1].AddLeaf(linkMathNodes(nodes[2:]))
			} else {
				panic("unknown symbol")
			}
		} else {
			panic("expected equals symbol")
		}
	} else if root.tokeType == token.Out {
		root.AddLeaf(nodes[1])
	}

	return root
}

func linkMathNodes(nodes []*Node) *Node {
	var root *Node
	var nodeStore *Node
	for i := 0; i < len(nodes); {
		if nodes[i].tokeType == token.Number || nodes[i].tokeType == token.Label {
			nodeStore = nodes[i]
		} else if nodes[i].tokeType == token.Plus || nodes[i].tokeType == token.Subtract {
			if root == nil {
				root = nodes[i]
			} else {
				root.AddLeaf(nodes[i])
			}

			if nodeStore != nil {
				nodes[i].AddLeaf(nodeStore)
				nodeStore = nil
			}
			i++
			if i < len(nodes) {
				if nodes[i].tokeType == token.Number || nodes[i].tokeType == token.Label {
					nodes[i-1].AddLeaf(nodes[i])
				}
			}
		} else {
			root.AddLeaf(nodes[i])
		}
		i++
	}
	return root
}
