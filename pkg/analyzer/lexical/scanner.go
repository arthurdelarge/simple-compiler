package lexical

import (
	"bufio"
	"errors"
	"github.com/arthurdelarge/simple-compiler/pkg/dictionary"
	"github.com/arthurdelarge/simple-compiler/pkg/symboltable"
	"github.com/arthurdelarge/simple-compiler/pkg/token"
	"os"
	"strings"
)

type Scanner struct {
	file           *os.File
	reader         *bufio.Reader
	lexemeBuilder  strings.Builder
	currentSymbol  byte
	previousSymbol byte
	stateMachine   *dictionary.MgolStateMachine
	symbolTable    *symboltable.SymbolTable
	row, column    int
}

func NewScanner(filename string) (*Scanner, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)

	return &Scanner{
		file:           file,
		reader:         reader,
		currentSymbol:  0,
		previousSymbol: 0,
		stateMachine:   dictionary.CreateMgolStateMachine(),
		symbolTable:    symboltable.CreateSymbolTable(),
		row:            0,
		column:         0,
	}, nil
}

func (s *Scanner) GetRow() int {
	return s.row
}

func (s *Scanner) GetColumn() int {
	return s.column - 1
}

func (s *Scanner) nextChar() (byte, error) {
	char, err := s.reader.ReadByte()
	if err != nil {
		return 0, err
	}

	s.column++
	if char == '\n' {
		s.column = 0
		s.row++
	}

	return char, nil
}

func (s *Scanner) NextToken() (token.Token, error) {
	stateMachineStop := false
	var state dictionary.State

	firstRun := true
	for !stateMachineStop {
		if s.currentSymbol == 0 {
			symbol, err := s.nextChar()

			if err != nil && symbol != 0 {
				panic(err)
			}

			s.currentSymbol = symbol
		}

		state, stateMachineStop = s.stateMachine.UpdateState(s.currentSymbol)

		if s.previousSymbol != 0 {
			s.lexemeBuilder.WriteByte(s.previousSymbol)
			s.previousSymbol = 0
		}

		if stateMachineStop && firstRun {
			s.lexemeBuilder.WriteByte(s.currentSymbol)
			s.currentSymbol = 0
		}

		if !stateMachineStop {
			s.previousSymbol = s.currentSymbol
			s.currentSymbol = 0
		}

		firstRun = false
	}

	tkn := s.identifySMState(state)

	if tkn.GetClass() == token.ClassEOF {
		return tkn, errors.New("EOF")
	}
	return tkn, nil
}

func (s *Scanner) GetCurrentLexeme() string {
	lexeme := s.lexemeBuilder.String()
	s.lexemeBuilder.Reset()
	return lexeme
}

func (s *Scanner) Close() error {
	return s.file.Close()
}

func (s *Scanner) searchSymbolTable(lexeme string) token.Token {
	return s.symbolTable.GetTokenFor(lexeme)
}

func (s *Scanner) identifySMState(state dictionary.State) token.Token {
	switch state {
	case dictionary.FinalStatesNum1, dictionary.FinalStatesNum2:
		return *token.NewToken(token.ClassNum, s.GetCurrentLexeme(), token.TypeInteger)
	case dictionary.FinalStatesNum3, dictionary.FinalStatesNum4, dictionary.FinalStatesNum5:
		return *token.NewToken(token.ClassNum, s.GetCurrentLexeme(), token.TypeReal)
	case dictionary.FinalStateLit:
		return *token.NewToken(token.ClassLit, s.GetCurrentLexeme(), token.TypeLiteral)
	case dictionary.FinalStateId:
		return s.searchSymbolTable(s.GetCurrentLexeme())
	case dictionary.FinalStateComment:
		return *token.NewToken(token.ClassComment, s.GetCurrentLexeme(), token.TypeNull)
	case dictionary.FinalStateEOF:
		return *token.NewToken(token.ClassEOF, s.GetCurrentLexeme(), token.TypeNull)
	case dictionary.FinalStateROperator1, dictionary.FinalStateROperator2, dictionary.FinalStateROperator3:
		return *token.NewToken(token.ClassROperator, s.GetCurrentLexeme(), token.TypeNull)
	case dictionary.FinalStateAOperator:
		return *token.NewToken(token.ClassAOperator, s.GetCurrentLexeme(), token.TypeNull)
	case dictionary.FinalStateReceive:
		return *token.NewToken(token.ClassReceive, s.GetCurrentLexeme(), token.TypeNull)
	case dictionary.FinalStateOpenP:
		return *token.NewToken(token.ClassOpenP, s.GetCurrentLexeme(), token.TypeNull)
	case dictionary.FinalStateCloseP:
		return *token.NewToken(token.ClassCloseP, s.GetCurrentLexeme(), token.TypeNull)
	case dictionary.FinalStateSemicolon:
		return *token.NewToken(token.ClassSemicolon, s.GetCurrentLexeme(), token.TypeNull)
	case dictionary.FinalStateComma:
		return *token.NewToken(token.ClassComma, s.GetCurrentLexeme(), token.TypeNull)
	case dictionary.FinalStateIgnore:
		return *token.NewToken(token.ClassIgnore, s.GetCurrentLexeme(), token.TypeNull)
	}

	return *token.NewToken(token.ClassError, s.GetCurrentLexeme(), token.TypeNull)
}
