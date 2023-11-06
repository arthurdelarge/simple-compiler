package lexical

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/arthurdelarge/simple-compiler/analyzer/lexical/dictionary"
)

type Scanner struct {
	file           *os.File
	reader         *bufio.Reader
	lexemeBuilder  strings.Builder
	currentSymbol  byte
	previousSymbol byte
	stateMachine   *dictionary.MgolStateMachine
	symbolTable    *SymbolTable
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
		symbolTable:    CreateSymbolTable(),
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

func (s *Scanner) NextChar() (byte, error) {
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

func (s *Scanner) NextToken() (Token, error) {
	stateMachineStop := false
	var state dictionary.SMState

	firstRun := true
	for !stateMachineStop {
		if s.currentSymbol == 0 {
			symbol, err := s.NextChar()

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

	token := s.identifySMState(state)

	if token.class == TokenClassEOF {
		return token, errors.New("EOF")
	}
	return token, nil
}

func (s *Scanner) GetCurrentLexeme() string {
	lexeme := s.lexemeBuilder.String()
	s.lexemeBuilder.Reset()
	return lexeme
}

func (s *Scanner) Close() error {
	return s.file.Close()
}

func (s *Scanner) searchSymbolTable(lexeme string) Token {
	return s.symbolTable.GetTokenFor(lexeme)
}

func (s *Scanner) identifySMState(state dictionary.SMState) Token {
	switch state {
	case dictionary.FinalStatesNum1, dictionary.FinalStatesNum2:
		return *NewToken(TokenClassNum, s.GetCurrentLexeme(), TokenTypeInteger)
	case dictionary.FinalStatesNum3, dictionary.FinalStatesNum4, dictionary.FinalStatesNum5:
		return *NewToken(TokenClassNum, s.GetCurrentLexeme(), TokenTypeReal)
	case dictionary.FinalStateLit:
		return *NewToken(TokenClassLit, s.GetCurrentLexeme(), TokenTypeLiteral)
	case dictionary.FinalStateId:
		return s.searchSymbolTable(s.GetCurrentLexeme())
	case dictionary.FinalStateComment:
		return *NewToken(TokenClassComment, s.GetCurrentLexeme(), TokenTypeNull)
	case dictionary.FinalStateEOF:
		return *NewToken(TokenClassEOF, s.GetCurrentLexeme(), TokenTypeNull)
	case dictionary.FinalStateROperator1, dictionary.FinalStateROperator2, dictionary.FinalStateROperator3:
		return *NewToken(TokenClassROperator, s.GetCurrentLexeme(), TokenTypeNull)
	case dictionary.FinalStateAOperator:
		return *NewToken(TokenClassAOperator, s.GetCurrentLexeme(), TokenTypeNull)
	case dictionary.FinalStateReceive:
		return *NewToken(TokenClassReceive, s.GetCurrentLexeme(), TokenTypeNull)
	case dictionary.FinalStateOpenP:
		return *NewToken(TokenClassOpenP, s.GetCurrentLexeme(), TokenTypeNull)
	case dictionary.FinalStateCloseP:
		return *NewToken(TokenClassCloseP, s.GetCurrentLexeme(), TokenTypeNull)
	case dictionary.FinalStateSemicolon:
		return *NewToken(TokenClassSemicolon, s.GetCurrentLexeme(), TokenTypeNull)
	case dictionary.FinalStateComma:
		return *NewToken(TokenClassComma, s.GetCurrentLexeme(), TokenTypeNull)
	case dictionary.FinalStateIgnore:
		return *NewToken(TokenClassIgnore, s.GetCurrentLexeme(), TokenTypeNull)
	}

	return *NewToken(TokenClassError, s.GetCurrentLexeme(), TokenTypeNull)
}
