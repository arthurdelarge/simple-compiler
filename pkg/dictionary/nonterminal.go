package dictionary

type NonTerminal byte

const (
	NTermP    NonTerminal = 0
	NTermV    NonTerminal = 1
	NTermL    NonTerminal = 2
	NTermLV   NonTerminal = 3
	NTermD    NonTerminal = 4
	NTermTIPO NonTerminal = 5
	NTermA    NonTerminal = 6
	NTermES   NonTerminal = 7
	NTermCMD  NonTerminal = 8
	NTermCOND NonTerminal = 9
	NTermR    NonTerminal = 10
	NTermCAB  NonTerminal = 11
	NTermCABR NonTerminal = 12
	NTermARG  NonTerminal = 13
	NTermLD   NonTerminal = 14
	NTermOPRD NonTerminal = 15
	NTermEXPR NonTerminal = 16
	NTermCP   NonTerminal = 17
	NTermCPR  NonTerminal = 18
	NTermPI   NonTerminal = 19
)
