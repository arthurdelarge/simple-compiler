package syntactic

import (
	"fmt"
	"github.com/arthurdelarge/simple-compiler/pkg/analyzer"
	"github.com/arthurdelarge/simple-compiler/pkg/analyzer/syntactic/automaton"
)

type Parser struct {
	automaton *automaton.PushdownAutomaton
	scanner   analyzer.LexicalAnalyzer
}

func NewParser(file string) *Parser {
	scanner, err := analyzer.NewLexicalAnalyzer(file)
	if err != nil {
		panic(err)
	}

	return &Parser{
		automaton: automaton.NewPushdownAutomaton(),
		scanner:   scanner,
	}
}

func (p *Parser) Parse() error {
	t, err := p.scanner.NextToken()
	movePointer := true

	for {
		movePointer, err = p.automaton.Move(t.GetClass())

		if err != nil {
			if err.Error() == "reject" || err.Error() == "accept" {
				fmt.Println(err.Error())
				break
			}

			row := p.scanner.GetRow() + 1
			col := p.scanner.GetColumn() + 1
			fmt.Printf("Erro Sintático, linha %d, coluna %d - %s\n", row, col, err.Error())
		}

		if movePointer {
			t, _ = p.scanner.NextToken()
		}
	}

	return nil
}

func (p *Parser) Close() error {
	return p.scanner.Close()
}
