package syntactic

import (
	"fmt"
	"os"

	"github.com/arthurdelarge/simple-compiler/pkg/analyzer"
	"github.com/arthurdelarge/simple-compiler/pkg/analyzer/automaton"
	"github.com/arthurdelarge/simple-compiler/pkg/analyzer/semantic"
)

type Parser struct {
	automaton *automaton.PushdownAutomaton
	semantic  *semantic.Semantic
	scanner   analyzer.LexicalAnalyzer
	hasError  bool
}

func NewParser(file string) *Parser {
	scanner, err := analyzer.NewLexicalAnalyzer(file)
	if err != nil {
		panic(err)
	}

	sem := semantic.NewSemantic()

	return &Parser{
		automaton: automaton.NewPushdownAutomaton(sem),
		scanner:   scanner,
		semantic:  sem,
		hasError:  false,
	}
}

func (p *Parser) Parse() error {
	t, _ := p.scanner.NextToken()

	for {
		movePointer, err := p.automaton.Move(t)

		if err != nil {
			if err.Error() == "reject" || err.Error() == "accept" {
				break
			}

			p.hasError = true
			p.printError(err)
		}

		if movePointer {
			t, _ = p.scanner.NextToken()
		}
	}

	if !p.hasError {
		file, err := os.OpenFile("PROGRAMA.c", os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		file.WriteString(p.semantic.Generate())
	}

	return nil
}

func (p *Parser) printError(err error) {
	row := p.scanner.GetRow() + 1
	col := p.scanner.GetColumn() + 1
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	fmt.Print(colorRed)
	fmt.Printf("Erro! linha %d, coluna %d - %s\n", row, col, err.Error())
	fmt.Print(colorReset)
}

func (p *Parser) Close() error {
	return p.scanner.Close()
}
