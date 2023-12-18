package automaton

import "github.com/arthurdelarge/simple-compiler/pkg/token"

func GetShiftTerminal() map[state][]token.Class {
	varTerminals := []token.Class{token.ClassInteiro, token.ClassReal, token.ClassLiteral, token.ClassVarfim}
	ptvTerminal := []token.Class{token.ClassSemicolon}
	idTerminal := []token.Class{token.ClassId}
	commandTerminals := []token.Class{
		token.ClassId,
		token.ClassLeia,
		token.ClassEscreva,
		token.ClassSe,
		token.ClassRepita,
		token.ClassFim,
	}

	return map[state][]token.Class{
		state(0):  []token.Class{token.ClassInicio},
		state(2):  []token.Class{token.ClassVarinicio},
		state(3):  varTerminals,
		state(6):  ptvTerminal,
		state(7):  idTerminal,
		state(12): ptvTerminal,
		state(13): []token.Class{token.ClassComma},
		state(15): idTerminal,
		state(17): varTerminals,
		state(19): commandTerminals,
		state(21): commandTerminals,
		state(23): commandTerminals,
		state(25): commandTerminals,
		state(27): commandTerminals,
		state(30): idTerminal,
		state(31): ptvTerminal,
		state(33): []token.Class{token.ClassId, token.ClassLit, token.ClassNum},
		state(34): ptvTerminal,
		state(39): []token.Class{token.ClassReceive},
		state(40): []token.Class{token.ClassId, token.ClassNum},
		state(41): ptvTerminal,
		state(45): []token.Class{token.ClassAOperator},
		state(46): []token.Class{token.ClassId, token.ClassNum},
		state(48): []token.Class{token.ClassOpenP},
		state(49): []token.Class{token.ClassId, token.ClassNum},
		state(50): []token.Class{token.ClassROperator},
		state(51): []token.Class{token.ClassId, token.ClassNum},
		state(53): []token.Class{token.ClassCloseP},
		state(54): []token.Class{token.ClassEntao},
		state(56): []token.Class{token.ClassId, token.ClassLeia, token.ClassEscreva, token.ClassSe, token.ClassFimse},
		state(58): []token.Class{token.ClassId, token.ClassLeia, token.ClassEscreva, token.ClassSe, token.ClassFimse},
		state(60): []token.Class{token.ClassId, token.ClassLeia, token.ClassEscreva, token.ClassSe, token.ClassFimse},
		state(62): []token.Class{token.ClassId, token.ClassLeia, token.ClassEscreva, token.ClassSe, token.ClassFimse},
		state(65): []token.Class{token.ClassId, token.ClassLeia, token.ClassEscreva, token.ClassSe, token.ClassFimrepita},
		state(67): []token.Class{token.ClassId, token.ClassLeia, token.ClassEscreva, token.ClassSe, token.ClassFimrepita},
		state(69): []token.Class{token.ClassId, token.ClassLeia, token.ClassEscreva, token.ClassSe, token.ClassFimrepita},
		state(71): []token.Class{token.ClassId, token.ClassLeia, token.ClassEscreva, token.ClassSe, token.ClassFimrepita},
		state(74): []token.Class{token.ClassOpenP},
		state(75): []token.Class{token.ClassId, token.ClassLeia},
		state(76): []token.Class{token.ClassCloseP},
	}
}

func GetReduceTerminal() map[state][]token.Class {
	eofTerminals := []token.Class{token.ClassEOF}
	varTerminals := []token.Class{token.ClassInteiro, token.ClassReal, token.ClassLiteral, token.ClassVarfim}
	ptvTerminal := []token.Class{token.ClassSemicolon}
	idTerminal := []token.Class{token.ClassId}
	commandTerminals := []token.Class{
		token.ClassId,
		token.ClassLeia,
		token.ClassEscreva,
		token.ClassSe,
		token.ClassRepita,
		token.ClassFim,
	}
	endCommandTerminals := []token.Class{
		token.ClassId,
		token.ClassLeia,
		token.ClassEscreva,
		token.ClassSe,
		token.ClassRepita,
		token.ClassFim,
		token.ClassFimse,
		token.ClassFimrepita,
	}

	return map[state][]token.Class{
		state(4):  commandTerminals,
		state(8):  idTerminal,
		state(9):  idTerminal,
		state(10): idTerminal,
		state(11): commandTerminals,
		state(13): []token.Class{token.ClassSemicolon},
		state(14): varTerminals,
		state(16): ptvTerminal,
		state(18): commandTerminals,
		state(20): eofTerminals,
		state(22): eofTerminals,
		state(24): eofTerminals,
		state(26): eofTerminals,
		state(28): eofTerminals,
		state(29): eofTerminals,
		state(32): endCommandTerminals,
		state(35): endCommandTerminals,
		state(36): ptvTerminal,
		state(37): ptvTerminal,
		state(38): ptvTerminal,
		state(42): endCommandTerminals,
		state(43): []token.Class{token.ClassSemicolon, token.ClassAOperator, token.ClassROperator, token.ClassCloseP},
		state(44): []token.Class{token.ClassSemicolon, token.ClassAOperator, token.ClassROperator, token.ClassCloseP},
		state(45): []token.Class{token.ClassSemicolon},
		state(47): ptvTerminal,
		state(52): []token.Class{token.ClassCloseP},
		state(55): []token.Class{token.ClassId, token.ClassLeia, token.ClassEscreva, token.ClassSe, token.ClassFimse},
		state(57): endCommandTerminals,
		state(59): endCommandTerminals,
		state(61): endCommandTerminals,
		state(63): endCommandTerminals,
		state(64): endCommandTerminals,
		state(66): commandTerminals,
		state(68): commandTerminals,
		state(70): commandTerminals,
		state(72): commandTerminals,
		state(73): commandTerminals,
		state(77): []token.Class{token.ClassId, token.ClassLeia, token.ClassEscreva, token.ClassSe, token.ClassFimrepita},
	}
}

func GetAcceptTerminal() map[state][]token.Class {
	eofTerminals := []token.Class{token.ClassEOF}

	return map[state][]token.Class{
		state(1): eofTerminals,
	}
}
