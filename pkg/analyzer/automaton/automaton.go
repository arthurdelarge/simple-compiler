package automaton

import (
	"errors"
	"fmt"
	"slices"

	"github.com/arthurdelarge/simple-compiler/pkg/analyzer/semantic"
	"github.com/arthurdelarge/simple-compiler/pkg/data"
	"github.com/arthurdelarge/simple-compiler/pkg/dictionary"
	"github.com/arthurdelarge/simple-compiler/pkg/token"
)

type PushdownAutomaton struct {
	stack           *data.Stack[state]
	actions         map[state]map[token.Class]action
	detours         map[state]map[dictionary.NonTerminal]state
	prods           map[byte]*dictionary.Production
	syncSymbols     map[state][]token.Class
	syncNonTerminal map[state]dictionary.NonTerminal
	semantic        *semantic.Semantic
	panicMode       bool
	status          string
}

func NewPushdownAutomaton(semantic *semantic.Semantic) *PushdownAutomaton {
	stack := data.CreateStack[state]()
	stack.Push(state(0))

	return &PushdownAutomaton{
		stack:           stack,
		actions:         GetMgolActionTable(),
		detours:         GetMgolDetourTable(),
		prods:           dictionary.GetMgolProductions(),
		syncNonTerminal: GetSyncNonTerminal(),
		syncSymbols:     GetSyncTerminals(),
		semantic:        semantic,
		panicMode:       false,
		status:          "accept",
	}
}

func (p *PushdownAutomaton) shift(act action) {
	state := act.value.(shiftAction).target
	p.stack.Push(state)
}

func (p *PushdownAutomaton) getProduction(act action) dictionary.Production {
	prodId := act.value.(reduceAction).target
	return *p.prods[prodId]
}

func (p *PushdownAutomaton) reduce(prod dictionary.Production) {
	for i := 0; i < prod.RightSideSize(); i++ {
		p.stack.Pop()
	}

	nterm := prod.LeftSide()
	currentState, _ := p.stack.Top()
	state := p.detours[currentState][nterm]
	p.stack.Push(state)
}

func (p *PushdownAutomaton) Move(tok *token.Token) (bool, error) {
	symbol := tok.GetClass()

	if p.panicMode {
		return p.panicMove(symbol)
	}

	currentState, _ := p.stack.Top()

	act, ok := p.actions[currentState][symbol]
	if !ok {
		return p.errorRecovery(symbol, currentState)
	}

	switch act.kind {
	case shiftActionKind:
		p.shift(act)
		p.semantic.PushToken(tok)
		return true, nil
	case reduceActionKind:
		prod := p.getProduction(act)
		p.reduce(prod)
		err := p.semantic.Evaluate(prod)
		return false, err
	case acceptActionKind:
		return false, errors.New(p.status)
	}

	return true, errors.New("no valid move")
}

func (p *PushdownAutomaton) errorRecovery(symbol token.Class, currentState state) (bool, error) {
	p.status = "reject"

	if p.canRecoverPhrase(symbol) {
		return p.recoverPhrase(symbol)
	}

	p.panicMode = true
	return false, errors.New(currentState.rejectMessage(symbol))
}

func (p *PushdownAutomaton) panicMove(symbol token.Class) (bool, error) {
	currentState, _ := p.stack.Top()

	isEndOfFile := symbol == token.ClassEOF
	if isEndOfFile {
		return false, errors.New(p.status)
	}

	isVarfim := symbol == token.ClassVarfim
	isVarscope := currentState < 19
	isResyncSymbol := slices.Contains(p.syncSymbols[currentState], symbol)
	if isVarfim && isVarscope && !isResyncSymbol {
		return p.resyncVarfim()
	}

	if !isResyncSymbol {
		fmt.Printf("Panico: descarte {%s}\n", symbol.String())
		return true, nil
	}

	fmt.Printf("RESYNC {%s}\n", symbol.String())
	act, _ := p.actions[currentState][symbol]

	switch act.kind {
	case shiftActionKind:
		p.shift(act)
		p.panicMode = false
		return true, nil
	case reduceActionKind:
		prod := p.getProduction(act)
		p.reduce(prod)
		p.panicMode = false
		return false, nil
	case acceptActionKind:
		p.panicMode = false
		return false, errors.New(p.status)
	}

	nterm := p.syncNonTerminal[currentState]
	currentState, _ = p.stack.Top()
	for _, ok := p.detours[currentState][nterm]; !ok; {
		p.stack.Pop()
		currentState, _ = p.stack.Top()
		_, ok = p.detours[currentState][nterm]
	}

	state := p.detours[currentState][nterm]
	p.stack.Push(state)

	p.panicMode = false
	return true, nil
}

func (p *PushdownAutomaton) resyncVarfim() (bool, error) {
	currentState, _ := p.stack.Top()

	for currentState > 3 {
		p.stack.Pop()
		currentState, _ = p.stack.Top()

	}

	if currentState == state(0) {
		p.stack.Push(state(2))
		currentState = state(2)
	}

	if currentState == state(2) {
		p.stack.Push(state(3))
		currentState = state(3)
	}

	act := p.actions[currentState][token.ClassVarfim]
	p.shift(act)

	p.panicMode = false
	fmt.Printf("RESYNC {%s}\n", token.ClassVarfim.String())
	return true, nil
}

func (p *PushdownAutomaton) canRecoverPhrase(symbol token.Class) bool {
	currentState, _ := p.stack.Top()
	act, canMoveWithSemicolon := p.actions[currentState][token.ClassSemicolon]
	if !canMoveWithSemicolon {
		return false
	}

	currentState = p.simulateAction(act)

	_, simulationSuccess := p.actions[currentState][symbol]
	if simulationSuccess {
		return true
	}

	return false
}

func (p *PushdownAutomaton) simulateAction(act action) state {
	var currentState state
	copyStack := p.stack.Copy()

	for act.kind == reduceActionKind {
		currentState, _ = copyStack.Top()
		act, _ = p.actions[currentState][token.ClassSemicolon]

		switch act.kind {
		case shiftActionKind:
			state := act.value.(shiftAction).target
			copyStack.Push(state)
		case reduceActionKind:
			prodId := act.value.(reduceAction).target
			prod := p.prods[prodId]
			for i := 0; i < prod.RightSideSize(); i++ {
				copyStack.Pop()
			}

			nterm := prod.LeftSide()
			currentState, _ := copyStack.Top()
			state := p.detours[currentState][nterm]
			copyStack.Push(state)
		}
	}

	currentState, _ = copyStack.Top()
	return currentState
}

func (p *PushdownAutomaton) recoverPhrase(symbol token.Class) (bool, error) {
	currentState, _ := p.stack.Top()
	act, _ := p.actions[currentState][token.ClassSemicolon]

	for act.kind == reduceActionKind {
		currentState, _ = p.stack.Top()
		act, _ = p.actions[currentState][token.ClassSemicolon]

		switch act.kind {
		case shiftActionKind:
			p.shift(act)
			return false, errors.New(fmt.Sprintf("Possível falta de {%s}", token.ClassSemicolon.String()))
		case reduceActionKind:
			prod := p.getProduction(act)
			p.reduce(prod)
		}
	}

	return false, errors.New("failed to recover")
}
