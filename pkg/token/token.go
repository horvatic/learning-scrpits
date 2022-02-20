package token

type Token struct {
	tokenType TokenType
	rawVal    string
}

func NewToken(tokenType TokenType, rawVal string) *Token {
	return &Token{
		tokenType: tokenType,
		rawVal:    rawVal,
	}
}

func (t *Token) GetTokenType() TokenType {
	return t.tokenType
}

func (t *Token) GetRawVal() string {
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

func NewEqualToken() *Token {
	return &Token{
		tokenType: Equal,
	}
}

func NewOutToken() *Token {
	return &Token{
		tokenType: Out,
	}
}
