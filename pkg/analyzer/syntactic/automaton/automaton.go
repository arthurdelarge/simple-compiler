package automaton

import (
	"errors"
	"fmt"
	"github.com/arthurdelarge/simple-compiler/pkg/data"
	"github.com/arthurdelarge/simple-compiler/pkg/token"
)

type PushdownAutomaton struct {
	stack           *data.Stack[state]
	actions         map[state]map[token.Class]action
	detours         map[state]map[nonTerminal]state
	prods           map[byte]*production
	syncSymbol      map[state]token.Class
	syncNonTerminal map[state]nonTerminal
	panicMode       bool
	status          string
}

func NewPushdownAutomaton() *PushdownAutomaton {
	stack := data.CreateStack[state]()
	stack.Push(state(0))

	syncSymbol := map[state]token.Class{
		state(3):  token.ClassVarfim,
		state(4):  token.ClassVarfim,
		state(6):  token.ClassSemicolon,
		state(7):  token.ClassSemicolon,
		state(8):  token.ClassSemicolon,
		state(9):  token.ClassSemicolon,
		state(10): token.ClassSemicolon,
		state(12): token.ClassSemicolon,
		state(13): token.ClassSemicolon,
		state(14): token.ClassVarfim,
		state(15): token.ClassSemicolon,
		state(16): token.ClassSemicolon,
		state(48): token.ClassEntao,
		state(49): token.ClassEntao,
		state(50): token.ClassCloseP,
		state(51): token.ClassCloseP,
		state(52): token.ClassCloseP,
		state(53): token.ClassEntao,
		state(54): token.ClassEntao,
		state(55): token.ClassEntao,
	}

	syncNonTerminal := map[state]nonTerminal{
		state(3):  NTermV,
		state(4):  NTermV,
		state(6):  NTermLV,
		state(7):  NTermD,
		state(8):  NTermD,
		state(9):  NTermD,
		state(10): NTermD,
		state(12): NTermD,
		state(13): NTermD,
		state(14): NTermV,
		state(15): NTermD,
		state(16): NTermD,
		state(48): NTermCAB,
		state(49): NTermCAB,
		state(50): NTermEXPR,
		state(51): NTermEXPR,
		state(52): NTermEXPR,
		state(53): NTermCAB,
		state(54): NTermCAB,
		state(55): NTermCAB,
	}

	return &PushdownAutomaton{
		stack:           stack,
		actions:         GetMgolActionTable(),
		detours:         GetMgolDetourTable(),
		prods:           GetMgolProductions(),
		syncNonTerminal: syncNonTerminal,
		syncSymbol:      syncSymbol,
		panicMode:       false,
		status:          "accept",
	}
}

func (p *PushdownAutomaton) Move(symbol token.Class) (bool, error) {
	//p.stack.Print()
	if p.panicMode {
		return p.panicMove(symbol)
	}

	currentState, ok := p.stack.Top()
	if !ok {
		panic("empty stack")
	}

	action, ok := p.actions[currentState][symbol]
	if !ok {
		p.status = "reject"
		p.panicMode = true
		return false, errors.New(currentState.rejectMessage(symbol))
	}

	switch action.kind {
	case shiftActionKind:
		state := action.value.(shift).target
		p.stack.Push(state)
		return true, nil
	case reduceActionKind:
		prodId := action.value.(reduce).target
		prod := p.prods[prodId]
		for i := byte(0); i < prod.RightSideSize(); i++ {
			p.stack.Pop()
		}

		nterm := prod.LeftSide()
		currentState, _ = p.stack.Top()
		state := p.detours[currentState][nterm]
		p.stack.Push(state)

		fmt.Printf("%v\n", prod.String())
		return false, nil
	case acceptActionKind:
		return false, errors.New(p.status)
	}

	return false, errors.New("no valid move")
}

func (p *PushdownAutomaton) panicMove(symbol token.Class) (bool, error) {
	currentState, ok := p.stack.Top()
	if !ok {
		panic("empty stack")
	}

	if symbol != p.syncSymbol[currentState] {
		if symbol == token.ClassEOF {
			return false, errors.New(p.status)
		}
		fmt.Printf("Panico: descarte {%s}\n", symbol.String())
		return true, nil
	}

	action, _ := p.actions[currentState][symbol]

	switch action.kind {
	case shiftActionKind:
		state := action.value.(shift).target
		p.stack.Push(state)
		p.panicMode = false
		fmt.Printf("RESYNC {%s}\n", symbol.String())
		return true, nil
	case reduceActionKind:
		prodId := action.value.(reduce).target
		prod := p.prods[prodId]
		for i := byte(0); i < prod.RightSideSize(); i++ {
			p.stack.Pop()
		}

		nterm := prod.LeftSide()
		currentState, _ = p.stack.Top()
		state := p.detours[currentState][nterm]
		p.stack.Push(state)

		fmt.Printf("RESYNC {%s}\n", symbol.String())
		p.panicMode = false
		return false, nil
	case acceptActionKind:
		return false, errors.New("accept")
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
	fmt.Printf("RESYNC {%s}\n", symbol.String())
	return true, nil
}
