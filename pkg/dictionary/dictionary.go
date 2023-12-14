package dictionary

type transition struct {
	from   State
	symbol byte
}

type MgolStateMachine struct {
	alphabet      *Alphabet
	initialState  State
	currentState  State
	isFinalState  map[State]bool
	currentSymbol byte
	transitionMap map[transition]State
}

func CreateMgolStateMachine() *MgolStateMachine {
	finalStates := GetFinalStates()

	isFinalState := make(map[State]bool, len(finalStates))
	for _, state := range finalStates {
		isFinalState[state] = true
	}

	stateMachine := &MgolStateMachine{
		alphabet:     CreateAlphabet(),
		initialState: InitialState,
		currentState: InitialState,
		isFinalState: isFinalState,
	}

	stateMachine.loadTransitionMap()

	return stateMachine
}

func (sm *MgolStateMachine) UpdateState(symbol byte) (State, bool) {
	key := transition{from: sm.currentState, symbol: symbol}
	nextState, found := sm.transitionMap[key]

	if !found {
		stopState := sm.currentState
		sm.currentState = 0

		return stopState, true
	}

	sm.currentState = nextState
	return sm.currentState, false
}

func (sm *MgolStateMachine) IsFinalState(s State) bool {
	return sm.isFinalState[s]
}

func (sm *MgolStateMachine) loadTransitionMap() {
	transitionMap := map[transition]State{}

	for _, digit := range sm.alphabet.digits {
		transitionMap[transition{from: 0, symbol: digit}] = 1
		transitionMap[transition{from: 1, symbol: digit}] = 2
		transitionMap[transition{from: 2, symbol: digit}] = 2
		transitionMap[transition{from: 3, symbol: digit}] = 5
		transitionMap[transition{from: 5, symbol: digit}] = 5
		transitionMap[transition{from: 4, symbol: digit}] = 6
		transitionMap[transition{from: 6, symbol: digit}] = 6
		transitionMap[transition{from: 7, symbol: digit}] = 9
		transitionMap[transition{from: 8, symbol: digit}] = 9
		transitionMap[transition{from: 9, symbol: digit}] = 9
		transitionMap[transition{from: 13, symbol: digit}] = 13
	}

	transitionMap[transition{from: 1, symbol: '.'}] = 3
	transitionMap[transition{from: 2, symbol: '.'}] = 4

	for _, e := range []byte{'e', 'E'} {
		transitionMap[transition{from: 1, symbol: e}] = 7
		transitionMap[transition{from: 5, symbol: e}] = 7
	}

	for _, opr := range []byte{'+', '-'} {
		transitionMap[transition{from: 7, symbol: opr}] = 8
	}

	transitionMap[transition{from: 0, symbol: '"'}] = 10
	for _, symbol := range sm.alphabet.symbols {
		if symbol == '"' {
			continue
		}

		transitionMap[transition{from: 10, symbol: symbol}] = 11
		transitionMap[transition{from: 11, symbol: symbol}] = 11
	}
	transitionMap[transition{from: 10, symbol: '"'}] = 12
	transitionMap[transition{from: 11, symbol: '"'}] = 12

	for _, letter := range sm.alphabet.letters {
		transitionMap[transition{from: 0, symbol: letter}] = 13
		transitionMap[transition{from: 13, symbol: letter}] = 13
	}
	transitionMap[transition{from: 13, symbol: '_'}] = 13

	transitionMap[transition{from: 0, symbol: '{'}] = 14
	for _, symbol := range sm.alphabet.symbols {
		if symbol == '}' {
			continue
		}

		transitionMap[transition{from: 14, symbol: symbol}] = 15
		transitionMap[transition{from: 15, symbol: symbol}] = 15
	}
	transitionMap[transition{from: 14, symbol: '}'}] = 16
	transitionMap[transition{from: 15, symbol: '}'}] = 16

	transitionMap[transition{from: 0, symbol: 0}] = 17

	transitionMap[transition{from: 0, symbol: '<'}] = 18
	transitionMap[transition{from: 0, symbol: '>'}] = 19
	transitionMap[transition{from: 0, symbol: '='}] = 20
	transitionMap[transition{from: 18, symbol: '>'}] = 20
	transitionMap[transition{from: 18, symbol: '='}] = 20
	transitionMap[transition{from: 19, symbol: '='}] = 20
	transitionMap[transition{from: 18, symbol: '-'}] = 21

	transitionMap[transition{from: 0, symbol: '+'}] = 22
	transitionMap[transition{from: 0, symbol: '-'}] = 22
	transitionMap[transition{from: 0, symbol: '*'}] = 22
	transitionMap[transition{from: 0, symbol: '/'}] = 22

	transitionMap[transition{from: 0, symbol: '('}] = 23
	transitionMap[transition{from: 0, symbol: ')'}] = 24
	transitionMap[transition{from: 0, symbol: ';'}] = 25
	transitionMap[transition{from: 0, symbol: ','}] = 26

	transitionMap[transition{from: 0, symbol: ' '}] = 27
	transitionMap[transition{from: 0, symbol: '\r'}] = 27
	transitionMap[transition{from: 0, symbol: '\n'}] = 27
	transitionMap[transition{from: 0, symbol: '\t'}] = 27

	sm.transitionMap = transitionMap
}
