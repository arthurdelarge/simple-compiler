package automaton

import (
	"errors"
	"fmt"
	"github.com/arthurdelarge/simple-compiler/pkg/data"
	"github.com/arthurdelarge/simple-compiler/pkg/token"
	"slices"
)

type PushdownAutomaton struct {
	stack           *data.Stack[state]
	actions         map[state]map[token.Class]action
	detours         map[state]map[nonTerminal]state
	prods           map[byte]*production
	syncSymbols     map[state][]token.Class
	syncNonTerminal map[state]nonTerminal
	panicMode       bool
	status          string
}

func NewPushdownAutomaton() *PushdownAutomaton {
	stack := data.CreateStack[state]()
	stack.Push(state(0))

	return &PushdownAutomaton{
		stack:           stack,
		actions:         GetMgolActionTable(),
		detours:         GetMgolDetourTable(),
		prods:           GetMgolProductions(),
		syncNonTerminal: GetSyncNonTerminal(),
		syncSymbols:     GetSyncSymbols(),
		panicMode:       false,
		status:          "accept",
	}
}

func GetSyncSymbols() map[state][]token.Class {
	return map[state][]token.Class{
		state(4):  []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassRepita, token.ClassFim},
		state(6):  []token.Class{token.ClassSemicolon},
		state(7):  []token.Class{token.ClassSemicolon},
		state(8):  []token.Class{token.ClassSemicolon},
		state(9):  []token.Class{token.ClassSemicolon},
		state(10): []token.Class{token.ClassSemicolon},
		state(11): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassRepita, token.ClassFim},
		state(12): []token.Class{token.ClassSemicolon},
		state(13): []token.Class{token.ClassSemicolon},
		state(14): []token.Class{token.ClassVarfim},
		state(15): []token.Class{token.ClassSemicolon},
		state(16): []token.Class{token.ClassSemicolon},
		state(17): []token.Class{token.ClassVarfim, token.ClassInteiro, token.ClassReal, token.ClassLiteral},
		state(18): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassRepita, token.ClassFim},
		state(19): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassRepita, token.ClassFim},
		state(20): []token.Class{token.ClassEOF},
		state(21): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassRepita, token.ClassFim},
		state(22): []token.Class{token.ClassEOF},
		state(23): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassRepita, token.ClassFim},
		state(24): []token.Class{token.ClassEOF},
		state(25): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassRepita, token.ClassFim},
		state(26): []token.Class{token.ClassEOF},
		state(27): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassRepita, token.ClassFim},
		state(28): []token.Class{token.ClassEOF},
		state(29): []token.Class{token.ClassEOF},
		state(30): []token.Class{token.ClassSemicolon},
		state(31): []token.Class{token.ClassSemicolon},
		state(32): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassRepita, token.ClassFim, token.ClassFimse, token.ClassFimrepita},
		state(33): []token.Class{token.ClassSemicolon},
		state(34): []token.Class{token.ClassSemicolon},
		state(35): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassRepita, token.ClassFim, token.ClassFimse, token.ClassFimrepita},
		state(36): []token.Class{token.ClassSemicolon},
		state(37): []token.Class{token.ClassSemicolon},
		state(38): []token.Class{token.ClassSemicolon},
		state(39): []token.Class{token.ClassSemicolon},
		state(40): []token.Class{token.ClassSemicolon},
		state(41): []token.Class{token.ClassSemicolon},
		state(42): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassRepita, token.ClassFim, token.ClassFimse, token.ClassFimrepita},
		state(43): []token.Class{token.ClassSemicolon, token.ClassROperator, token.ClassAOperator, token.ClassCloseP},
		state(44): []token.Class{token.ClassSemicolon, token.ClassROperator, token.ClassAOperator, token.ClassCloseP},
		state(45): []token.Class{token.ClassSemicolon},
		state(46): []token.Class{token.ClassSemicolon},
		state(47): []token.Class{token.ClassSemicolon},
		state(48): []token.Class{token.ClassEntao},
		state(49): []token.Class{token.ClassEntao},
		state(50): []token.Class{token.ClassCloseP},
		state(51): []token.Class{token.ClassCloseP},
		state(52): []token.Class{token.ClassCloseP},
		state(53): []token.Class{token.ClassEntao},
		state(54): []token.Class{token.ClassEntao},
		state(55): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassFimse},
		state(56): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassFimse},
		state(57): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassRepita, token.ClassFim, token.ClassFimse, token.ClassFimrepita},
		state(58): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassFimse},
		state(59): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassRepita, token.ClassFim, token.ClassFimse, token.ClassFimrepita},
		state(60): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassFimse},
		state(61): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassRepita, token.ClassFim, token.ClassFimse, token.ClassFimrepita},
		state(62): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassFimse},
		state(63): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassRepita, token.ClassFim, token.ClassFimse, token.ClassFimrepita},
		state(64): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassRepita, token.ClassFim, token.ClassFimse, token.ClassFimrepita},
		state(65): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassFimrepita},
		state(66): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassRepita, token.ClassFim},
		state(67): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassFimrepita},
		state(68): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassRepita, token.ClassFim},
		state(69): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassFimrepita},
		state(70): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassRepita, token.ClassFim},
		state(71): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassFimrepita},
		state(72): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassRepita, token.ClassFim},
		state(73): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassRepita, token.ClassFim},
		state(74): []token.Class{token.ClassCloseP},
		state(75): []token.Class{token.ClassCloseP},
		state(76): []token.Class{token.ClassCloseP},
		state(77): []token.Class{token.ClassLeia, token.ClassEscreva, token.ClassId, token.ClassSe, token.ClassRepita, token.ClassFim},
	}
}

func GetSyncNonTerminal() map[state]nonTerminal {
	return map[state]nonTerminal{
		state(4):  NTermV,
		state(6):  NTermLV,
		state(7):  NTermD,
		state(8):  NTermD,
		state(9):  NTermD,
		state(10): NTermD,
		state(11): NTermLV,
		state(12): NTermD,
		state(13): NTermD,
		state(14): NTermV,
		state(15): NTermD,
		state(16): NTermD,
		state(17): NTermLV,
		state(18): NTermLV,
		state(19): NTermA,
		state(20): NTermP,
		state(21): NTermA,
		state(22): NTermA,
		state(23): NTermA,
		state(24): NTermA,
		state(25): NTermA,
		state(26): NTermA,
		state(27): NTermA,
		state(28): NTermA,
		state(29): NTermA,
		state(30): NTermES,
		state(31): NTermES,
		state(32): NTermES,
		state(33): NTermES,
		state(34): NTermES,
		state(35): NTermES,
		state(36): NTermARG,
		state(37): NTermARG,
		state(38): NTermARG,
		state(39): NTermCMD,
		state(40): NTermCMD,
		state(41): NTermCMD,
		state(42): NTermCMD,
		state(43): NTermOPRD,
		state(44): NTermOPRD,
		state(45): NTermLD,
		state(46): NTermLD,
		state(47): NTermLD,
		state(48): NTermCAB,
		state(49): NTermCAB,
		state(50): NTermEXPR,
		state(51): NTermEXPR,
		state(52): NTermEXPR,
		state(53): NTermCAB,
		state(54): NTermCAB,
		state(55): NTermCAB,
		state(56): NTermCOND,
		state(57): NTermCOND,
		state(58): NTermCP,
		state(59): NTermCP,
		state(60): NTermCP,
		state(61): NTermCP,
		state(62): NTermCP,
		state(63): NTermCP,
		state(64): NTermCP,
		state(65): NTermR,
		state(66): NTermR,
		state(67): NTermCPR,
		state(68): NTermCPR,
		state(69): NTermCPR,
		state(70): NTermCPR,
		state(71): NTermCPR,
		state(72): NTermCPR,
		state(73): NTermCPR,
		state(74): NTermCABR,
		state(75): NTermCABR,
		state(76): NTermCABR,
		state(77): NTermCABR,
	}
}

func (p *PushdownAutomaton) shift(act action) {
	state := act.value.(shiftAction).target
	p.stack.Push(state)
}

func (p *PushdownAutomaton) reduce(act action) *production {
	prodId := act.value.(reduceAction).target
	prod := p.prods[prodId]
	for i := byte(0); i < prod.RightSideSize(); i++ {
		p.stack.Pop()
	}

	nterm := prod.LeftSide()
	currentState, _ := p.stack.Top()
	state := p.detours[currentState][nterm]
	p.stack.Push(state)
	return prod
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

	act, ok := p.actions[currentState][symbol]
	if !ok {
		p.status = "reject"
		p.panicMode = true
		return false, errors.New(currentState.rejectMessage(symbol))
	}

	switch act.kind {
	case shiftActionKind:
		p.shift(act)
		return true, nil
	case reduceActionKind:
		production := p.reduce(act)
		fmt.Printf("%v\n", production.String())
		return false, nil
	case acceptActionKind:
		return false, errors.New(p.status)
	}

	return false, errors.New("no valid move")
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
		production := p.reduce(act)
		fmt.Printf("%v\n", production.String())
		p.panicMode = false
		return false, nil
	case acceptActionKind:
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
