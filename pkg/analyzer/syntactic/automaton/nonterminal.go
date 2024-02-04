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
