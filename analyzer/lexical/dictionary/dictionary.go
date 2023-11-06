package dictionary

const (
	FinalStatesNum1      SMState = 1
	FinalStatesNum2      SMState = 2
	FinalStatesNum3      SMState = 5
	FinalStatesNum4      SMState = 6
	FinalStatesNum5      SMState = 9
	FinalStateLit        SMState = 12
	FinalStateId         SMState = 13
	FinalStateComment    SMState = 16
	FinalStateEOF        SMState = 17
	FinalStateROperator1 SMState = 18
	FinalStateROperator2 SMState = 19
	FinalStateROperator3 SMState = 20
	FinalStateAOperator  SMState = 22
	FinalStateReceive    SMState = 21
	FinalStateOpenP      SMState = 23
	FinalStateCloseP     SMState = 24
	FinalStateSemicolon  SMState = 25
	FinalStateComma      SMState = 26
	FinalStateIgnore     SMState = 27
)

type transitionKey struct {
	from   SMState
	symbol byte
}

type SMState byte

type MgolStateMachine struct {
	alphabet      *Alphabet
	initialState  SMState
	finalStates   map[SMState]bool
	currentState  SMState
	currentSymbol byte
	transitionMap map[transitionKey]SMState
}

func CreateMgolStateMachine() *MgolStateMachine {
	finalStateNumbers := []SMState{
		FinalStatesNum1, FinalStatesNum2, FinalStatesNum3, FinalStatesNum4, FinalStatesNum5,
		FinalStateLit, FinalStateId, FinalStateComment, FinalStateEOF, FinalStateROperator1,
		FinalStateROperator2, FinalStateROperator3, FinalStateAOperator, FinalStateReceive,
		FinalStateOpenP, FinalStateCloseP, FinalStateSemicolon, FinalStateComma, FinalStateIgnore,
	}

	finalStates := make(map[SMState]bool, 28)
	for _, state := range finalStateNumbers {
		finalStates[state] = true
	}

	stateMachine := &MgolStateMachine{
		alphabet:     CreateAlphabet(),
		initialState: SMState(0),
		currentState: SMState(0),
		finalStates:  finalStates,
	}

	stateMachine.loadTransitionMap()

	return stateMachine
}

func (sm *MgolStateMachine) UpdateState(symbol byte) (SMState, bool) {
	key := transitionKey{from: sm.currentState, symbol: symbol}
	nextState, found := sm.transitionMap[key]

	if !found {
		stopState := sm.currentState
		sm.currentState = 0

		return stopState, true
	}

	sm.currentState = nextState
	return sm.currentState, false
}

func (sm *MgolStateMachine) IsFinalState(s SMState) bool {
	return sm.finalStates[s]
}

func (sm *MgolStateMachine) loadTransitionMap() {
	transitionMap := map[transitionKey]SMState{}

	for _, digit := range sm.alphabet.digits {
		transitionMap[transitionKey{from: 0, symbol: digit}] = 1
		transitionMap[transitionKey{from: 1, symbol: digit}] = 2
		transitionMap[transitionKey{from: 2, symbol: digit}] = 2
		transitionMap[transitionKey{from: 3, symbol: digit}] = 5
		transitionMap[transitionKey{from: 5, symbol: digit}] = 5
		transitionMap[transitionKey{from: 4, symbol: digit}] = 6
		transitionMap[transitionKey{from: 6, symbol: digit}] = 6
		transitionMap[transitionKey{from: 7, symbol: digit}] = 9
		transitionMap[transitionKey{from: 8, symbol: digit}] = 9
		transitionMap[transitionKey{from: 9, symbol: digit}] = 9
		transitionMap[transitionKey{from: 13, symbol: digit}] = 13
	}

	transitionMap[transitionKey{from: 1, symbol: '.'}] = 3
	transitionMap[transitionKey{from: 2, symbol: '.'}] = 4

	for _, e := range []byte{'e', 'E'} {
		transitionMap[transitionKey{from: 1, symbol: e}] = 7
		transitionMap[transitionKey{from: 5, symbol: e}] = 7
	}

	for _, opr := range []byte{'+', '-'} {
		transitionMap[transitionKey{from: 7, symbol: opr}] = 8
	}

	transitionMap[transitionKey{from: 0, symbol: '"'}] = 10
	for _, symbol := range sm.alphabet.symbols {
		if symbol == '"' {
			continue
		}

		transitionMap[transitionKey{from: 10, symbol: symbol}] = 11
		transitionMap[transitionKey{from: 11, symbol: symbol}] = 11
	}
	transitionMap[transitionKey{from: 10, symbol: '"'}] = 12
	transitionMap[transitionKey{from: 11, symbol: '"'}] = 12

	for _, letter := range sm.alphabet.letters {
		transitionMap[transitionKey{from: 0, symbol: letter}] = 13
		transitionMap[transitionKey{from: 13, symbol: letter}] = 13
	}
	transitionMap[transitionKey{from: 13, symbol: '_'}] = 13

	transitionMap[transitionKey{from: 0, symbol: '{'}] = 14
	for _, symbol := range sm.alphabet.symbols {
		if symbol == '}' {
			continue
		}

		transitionMap[transitionKey{from: 14, symbol: symbol}] = 15
		transitionMap[transitionKey{from: 15, symbol: symbol}] = 15
	}
	transitionMap[transitionKey{from: 14, symbol: '}'}] = 16
	transitionMap[transitionKey{from: 15, symbol: '}'}] = 16

	transitionMap[transitionKey{from: 0, symbol: 0}] = 17

	transitionMap[transitionKey{from: 0, symbol: '<'}] = 18
	transitionMap[transitionKey{from: 0, symbol: '>'}] = 19
	transitionMap[transitionKey{from: 0, symbol: '='}] = 20
	transitionMap[transitionKey{from: 18, symbol: '>'}] = 20
	transitionMap[transitionKey{from: 18, symbol: '='}] = 20
	transitionMap[transitionKey{from: 19, symbol: '='}] = 20
	transitionMap[transitionKey{from: 18, symbol: '-'}] = 21

	transitionMap[transitionKey{from: 0, symbol: '+'}] = 22
	transitionMap[transitionKey{from: 0, symbol: '-'}] = 22
	transitionMap[transitionKey{from: 0, symbol: '*'}] = 22
	transitionMap[transitionKey{from: 0, symbol: '/'}] = 22

	transitionMap[transitionKey{from: 0, symbol: '('}] = 23
	transitionMap[transitionKey{from: 0, symbol: ')'}] = 24
	transitionMap[transitionKey{from: 0, symbol: ';'}] = 25
	transitionMap[transitionKey{from: 0, symbol: ','}] = 26

	transitionMap[transitionKey{from: 0, symbol: ' '}] = 27
	transitionMap[transitionKey{from: 0, symbol: '\r'}] = 27
	transitionMap[transitionKey{from: 0, symbol: '\n'}] = 27
	transitionMap[transitionKey{from: 0, symbol: '\t'}] = 27

	sm.transitionMap = transitionMap
}
