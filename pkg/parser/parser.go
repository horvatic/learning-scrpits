package parser

import (
	"strconv"

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
		if codeBlock[i].GetTokenType() == token.Int {
			buildingNum := true
			num := ""
			for buildingNum {
				num += string(codeBlock[i].GetRawVal())
				i++
				if i >= len(codeBlock) || codeBlock[i].GetTokenType() != token.Int {
					buildingNum = false
				}
			}
			if builtInt, err := strconv.Atoi(num); err != nil {
				panic(err)
			} else {
				nodes = append(nodes, BuildNode(token.Int, builtInt))
			}
		} else if codeBlock[i].GetTokenType() == token.Plus {
			nodes = append(nodes, BuildNode(token.Plus, nil))
			i++
		} else if codeBlock[i].GetTokenType() == token.EOF {
			nodes = append(nodes, BuildNode(token.EOF, nil))
			i++
		}
	}
	return linkNodes(nodes)
}

func linkNodes(nodes []*Node) *Node {
	var root *Node
	var nodeStore *Node
	for i := 0; i < len(nodes); {
		if nodes[i].tokeType == token.Int {
			nodeStore = nodes[i]
		} else if nodes[i].tokeType == token.Plus {
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
				if nodes[i].tokeType == token.Int {
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
