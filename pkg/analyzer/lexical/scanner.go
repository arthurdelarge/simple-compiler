package lexical

import (
	"bufio"
	"errors"
	"fmt"
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

func (s *Scanner) nextToken() (token.Token, error) {
	mustReadNextSymbol := true
	if s.hasValidSymbolInBuffer() {
		mustReadNextSymbol = false
	}

	var state dictionary.State
	machineStop := false
	for !machineStop {
		if mustReadNextSymbol {
			symbol, err := s.nextChar()

			if err != nil && symbol != 0 {
				panic(err)
			}

			s.currentSymbol = symbol
			mustReadNextSymbol = false
		}

		state, machineStop = s.stateMachine.UpdateState(s.currentSymbol)

		if s.previousSymbol != 0 {
			s.lexemeBuilder.WriteByte(s.previousSymbol)
			s.previousSymbol = 0
		}

		if machineStop && s.stateMachine.RejectInLastMove() && state == dictionary.InitialState {
			s.lexemeBuilder.WriteByte(s.currentSymbol)
		}

		if !machineStop {
			s.previousSymbol = s.currentSymbol
			mustReadNextSymbol = true
		}

	}

	if s.IsError(state) {
		errorMsg := s.getErrorMessage(state)
		errorTkn := *token.NewToken(token.ClassError, s.GetCurrentLexeme(), token.TypeNull)
		return errorTkn, errors.New(errorMsg)
	}

	tkn := s.identifySMState(state)

	return tkn, nil
}

func (s *Scanner) NextToken() (token.Token, error) {
	var tkn token.Token
	var err error
	for {
		tkn, err = s.nextToken()

		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			continue
		}

		if tkn.GetClass() == token.ClassIgnore {
			continue
		}

		if tkn.GetClass() == token.ClassComment {
			continue
		}

		break
	}

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

func (s *Scanner) getErrorMessage(state dictionary.State) string {
	msg := "PANIC"

	if !s.stateMachine.InAlphabet(s.currentSymbol) {
		msg = s.strangeCharacterMessage()
	} else if !s.stateMachine.IsFinalState(state) {
		msg = s.notFinalStageMessage(state)
	}

	return msg
}

func (s *Scanner) strangeCharacterMessage() string {
	return fmt.Sprintf("Erro léxico - caractere \u001B[0;31m %s \u001B[0m desconhecido. Linha %d, coluna %d.\n", s.GetCurrentLexeme(), s.GetRow()+1, s.GetColumn()+1)
}

func (s *Scanner) notFinalStageMessage(state dictionary.State) string {
	lex := s.GetCurrentLexeme()
	row := s.GetRow() + 1
	col := s.GetColumn() + 1
	switch s.stateMachine.GetError(state) {
	case dictionary.InvalidNumber:
		return fmt.Sprintf("Erro Léxico - número \u001B[0;31m %s \u001B[0m inválido. Linha %d, coluna %d.\n", lex, row, col)
	case dictionary.IncompleteLiteral:
		return fmt.Sprintf("Erro Léxico, Linha %d, coluna %d - Literal incompleto:\u001B[0;31m %s \u001B[0m\n", row, col, lex)
	case dictionary.IncompleteComment:
		return fmt.Sprintf("Erro Léxico, Linha %d, coluna %d - Comentário incompleto:\u001B[0;31m %s \u001B[0m\n", row, col, lex)
	}

	return fmt.Sprintf("Erro Léxico - palavra\u001B[0;31m %s\u001B[0m inválida. Linha %d, coluna %d.\n", lex, row, col)
}

func (s *Scanner) IsError(state dictionary.State) bool {
	return !s.stateMachine.IsFinalState(state) || !s.stateMachine.InAlphabet(s.currentSymbol)
}

func (s *Scanner) hasValidSymbolInBuffer() bool {
	acceptedLastMove := s.stateMachine.StoppedInLastMove() && !s.stateMachine.RejectInLastMove()
	rejectedButKnownChar := s.stateMachine.RejectInLastMove() && s.stateMachine.InAlphabet(s.currentSymbol)
	rejectedInInitialSate := s.stateMachine.GetLastStopState() == dictionary.InitialState
	return acceptedLastMove || (rejectedButKnownChar && !rejectedInInitialSate)
}
