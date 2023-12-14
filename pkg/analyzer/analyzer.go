package analyzer

import (
	"github.com/arthurdelarge/simple-compiler/pkg/analyzer/lexical"
	"github.com/arthurdelarge/simple-compiler/pkg/token"
)

type LexicalAnalyzer interface {
	NextToken() (token.Token, error)
	Close() error
	GetRow() int
	GetColumn() int
}

func NewLexicalAnalyzer(filename string) (LexicalAnalyzer, error) {
	return lexical.NewScanner(filename)
}
