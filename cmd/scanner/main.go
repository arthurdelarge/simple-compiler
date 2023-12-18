package main

import (
	"fmt"
	"github.com/arthurdelarge/simple-compiler/pkg/analyzer"
	"github.com/arthurdelarge/simple-compiler/pkg/token"
)

func main() {
	file := "input.txt"
	scanner, err := analyzer.NewLexicalAnalyzer(file)
	if err != nil {
		panic(err)
	}
	defer scanner.Close()

	var tkn token.Token
	for err == nil {
		tkn, err = scanner.NextToken()
		fmt.Printf("Classe: %s, Lexema: %s, Tipo: %s\n", tkn.GetClass().String(), tkn.GetLexeme(), tkn.GetType().String())
	}
}
