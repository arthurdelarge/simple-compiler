package token

type Token struct {
	class     Class
	lexeme    string
	tokenType Type
}

func NewToken(class Class, lexeme string, tType Type) *Token {
	return &Token{
		class:     class,
		lexeme:    lexeme,
		tokenType: tType,
	}
}

func (t *Token) GetClass() Class {
	return t.class
}

func (t *Token) GetLexeme() string {
	if t.class == ClassEOF {
		return "EOF"
	}
	return t.lexeme
}

func (t *Token) GetType() Type {
	return t.tokenType
}
