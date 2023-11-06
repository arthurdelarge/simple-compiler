package main

import (
	"fmt"

	"github.com/arthurdelarge/simple-compiler/analyzer"
	"github.com/arthurdelarge/simple-compiler/analyzer/lexical"
)

func main() {
	file := "input.txt"
	scanner, err := analyzer.NewLexicalAnalyzer(file)
	if err != nil {
		panic(err)
	}
	defer scanner.Close()

	var token lexical.Token
	for err == nil {
		token, err = scanner.NextToken()
		if token.GetClass() == lexical.TokenClassIgnore {
			continue
		}

		if token.GetClass() == lexical.TokenClassError {
			if len(token.GetLexeme()) > 1 {
				fmt.Printf("Classe: ERRO léxico - palavra %s inválida. Linha %d, coluna %d.\n", token.GetLexeme(), scanner.GetRow()+1, scanner.GetColumn()+1)
			} else {
				fmt.Printf("Classe: ERRO léxico - caractere %s inválido. Linha %d, coluna %d.\n", token.GetLexeme(), scanner.GetRow()+1, scanner.GetColumn()+1)
			}
			continue
		}
		fmt.Printf("Classe: %s, Lexema: %s, Tipo: %s\n", token.GetClass().String(), token.GetLexeme(), token.GetType().String())
	}
}
