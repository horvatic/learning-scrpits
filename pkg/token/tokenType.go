package token

type TokenType int64

const (
	Plus TokenType = iota
	Subtract
	EndCodeBlock
	Start
	Comments
	EOF
	Number
	Equal
	Label
	Type

	Out
)
