package automaton

type nonTerminal byte

const (
	NTermP    nonTerminal = 0
	NTermV    nonTerminal = 1
	NTermL    nonTerminal = 2
	NTermLV   nonTerminal = 3
	NTermD    nonTerminal = 4
	NTermTIPO nonTerminal = 5
	NTermA    nonTerminal = 6
	NTermES   nonTerminal = 7
	NTermCMD  nonTerminal = 8
	NTermCOND nonTerminal = 9
	NTermR    nonTerminal = 10
	NTermCAB  nonTerminal = 11
	NTermCABR nonTerminal = 12
	NTermARG  nonTerminal = 13
	NTermLD   nonTerminal = 14
	NTermOPRD nonTerminal = 15
	NTermEXPR nonTerminal = 16
	NTermCP   nonTerminal = 17
	NTermCPR  nonTerminal = 18
	NTermPI   nonTerminal = 19
)
