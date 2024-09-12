package grammar

type TokenType int8

var (
	TokenTypeDigit         TokenType = 0 // 1-0
	TokenTypeChar          TokenType = 3 // a-z and A-Z
	TokenTypePlusOperator  TokenType = 4 // +
	TokenTypeMinusOperator TokenType = 5 // -
	TokenTypeSlashOperator TokenType = 6 // /
	TokenTypeEqualOperator TokenType = 7 // =
)

type Token struct {
	Type     TokenType
	Literal  string
	Position int
}

func NewToken(tokenType TokenType, literal string, position int) *Token {
	return &Token{
		Type:     tokenType,
		Literal:  literal,
		Position: position,
	}
}
