package analyzer

import "github.com/arthurdelarge/simple-compiler/analyzer/lexical"

type LexicalAnalyzer interface {
	NextToken() (lexical.Token, error)
	Close() error
	GetRow() int
	GetColumn() int
}

func NewLexicalAnalyzer(filename string) (LexicalAnalyzer, error) {
	return lexical.NewScanner(filename)
}
