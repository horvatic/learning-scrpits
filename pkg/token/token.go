package token

type Token struct {
	tokenType TokenType
	rawVal    rune
}

func NewToken(tokenType TokenType, rawVal rune) *Token {
	return &Token{
		tokenType: tokenType,
		rawVal:    rawVal,
	}
}

func (t *Token) GetTokenType() TokenType {
	return t.tokenType
}

func (t *Token) GetRawVal() rune {
	return t.rawVal
}

func NewEOFToken() *Token {
	return &Token{
		tokenType: EOF,
	}
}
func NewEndCodeBlockToken() *Token {
	return &Token{
		tokenType: EndCodeBlock,
	}
}
