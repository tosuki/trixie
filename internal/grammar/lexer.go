package grammar

type Lexer struct {
	Tokens          []*Token
	CurrentPosition int
	Raw             string
}

func (lexer *Lexer) Tokenize() error {
	return nil
}

func NewLexer(raw string) *Lexer {
	return &Lexer{
		Raw:             raw,
		CurrentPosition: 0,
		Tokens:          []*Token{},
	}
}
