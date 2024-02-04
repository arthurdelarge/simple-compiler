package automaton

import (
	"github.com/arthurdelarge/simple-compiler/pkg/dictionary"
	"github.com/arthurdelarge/simple-compiler/pkg/token"
)

type actionKind byte

const (
	shiftActionKind  actionKind = 1
	reduceActionKind actionKind = 2
	acceptActionKind actionKind = 3
)

type action struct {
	kind  actionKind
	value interface{}
}

type shiftAction struct {
	target state
}

type reduceAction struct {
	target byte
}

func GetMgolActionTable() map[state]map[token.Class]action {
	table := make(map[state]map[token.Class]action)
	reduceTerms := GetReduceTerminal()

	for i := state(0); i <= state(77); i++ {
		table[i] = make(map[token.Class]action)
	}

	table[state(0)][token.ClassInicio] = action{kind: shiftActionKind, value: shiftAction{target: state(2)}}
	table[state(1)][token.ClassEOF] = action{kind: acceptActionKind, value: nil}
	table[state(2)][token.ClassVarinicio] = action{kind: shiftActionKind, value: shiftAction{target: state(3)}}
	table[state(3)][token.ClassInteiro] = action{kind: shiftActionKind, value: shiftAction{target: state(8)}}
	table[state(3)][token.ClassReal] = action{kind: shiftActionKind, value: shiftAction{target: state(9)}}
	table[state(3)][token.ClassLiteral] = action{kind: shiftActionKind, value: shiftAction{target: state(10)}}
	table[state(3)][token.ClassVarfim] = action{kind: shiftActionKind, value: shiftAction{target: state(6)}}
	for _, term := range reduceTerms[state(4)] {
		table[state(4)][term] = action{kind: reduceActionKind, value: reduceAction{target: 3}}
	}
	table[state(6)][token.ClassSemicolon] = action{kind: shiftActionKind, value: shiftAction{target: state(11)}}
	table[state(7)][token.ClassId] = action{kind: shiftActionKind, value: shiftAction{target: state(13)}}
	table[state(8)][token.ClassId] = action{reduceActionKind, reduceAction{target: 9}}
	table[state(9)][token.ClassId] = action{reduceActionKind, reduceAction{target: 10}}
	table[state(10)][token.ClassId] = action{reduceActionKind, reduceAction{target: 11}}
	for _, term := range reduceTerms[state(11)] {
		table[state(11)][term] = action{kind: reduceActionKind, value: reduceAction{target: 5}}
	}
	table[state(12)][token.ClassSemicolon] = action{shiftActionKind, shiftAction{target: state(14)}}
	table[state(13)][token.ClassComma] = action{kind: shiftActionKind, value: shiftAction{target: state(15)}}
	table[state(13)][token.ClassSemicolon] = action{reduceActionKind, reduceAction{target: 8}}
	for _, term := range reduceTerms[state(14)] {
		table[state(14)][term] = action{kind: reduceActionKind, value: reduceAction{target: 6}}
	}
	table[state(15)][token.ClassId] = action{kind: shiftActionKind, value: shiftAction{target: state(13)}}
	table[state(16)][token.ClassSemicolon] = action{reduceActionKind, reduceAction{target: 7}}
	table[state(17)][token.ClassInteiro] = action{kind: shiftActionKind, value: shiftAction{target: state(8)}}
	table[state(17)][token.ClassReal] = action{kind: shiftActionKind, value: shiftAction{target: state(9)}}
	table[state(17)][token.ClassLiteral] = action{kind: shiftActionKind, value: shiftAction{target: state(10)}}
	table[state(17)][token.ClassVarfim] = action{kind: shiftActionKind, value: shiftAction{target: state(6)}}
	for _, term := range reduceTerms[state(18)] {
		table[state(18)][term] = action{kind: reduceActionKind, value: reduceAction{target: 4}}
	}
	for _, t := range []state{19, 21, 23, 25, 27} {
		table[t][token.ClassId] = action{kind: shiftActionKind, value: shiftAction{target: state(39)}}
		table[t][token.ClassLeia] = action{kind: shiftActionKind, value: shiftAction{target: state(30)}}
		table[t][token.ClassEscreva] = action{kind: shiftActionKind, value: shiftAction{target: state(33)}}
		table[t][token.ClassSe] = action{kind: shiftActionKind, value: shiftAction{target: state(48)}}
		table[t][token.ClassRepita] = action{kind: shiftActionKind, value: shiftAction{target: state(74)}}
		table[t][token.ClassFim] = action{kind: shiftActionKind, value: shiftAction{target: state(29)}}
	}
	table[state(20)][token.ClassEOF] = action{reduceActionKind, reduceAction{target: 2}}
	table[state(22)][token.ClassEOF] = action{reduceActionKind, reduceAction{target: 12}}
	table[state(24)][token.ClassEOF] = action{reduceActionKind, reduceAction{target: 18}}
	table[state(26)][token.ClassEOF] = action{reduceActionKind, reduceAction{target: 24}}
	table[state(28)][token.ClassEOF] = action{reduceActionKind, reduceAction{target: 32}}
	table[state(29)][token.ClassEOF] = action{reduceActionKind, reduceAction{target: 39}}
	table[state(30)][token.ClassId] = action{kind: shiftActionKind, value: shiftAction{target: state(31)}}
	table[state(31)][token.ClassSemicolon] = action{kind: shiftActionKind, value: shiftAction{target: state(32)}}
	for _, term := range reduceTerms[state(32)] {
		table[state(32)][term] = action{kind: reduceActionKind, value: reduceAction{target: 13}}
	}
	table[state(33)][token.ClassId] = action{kind: shiftActionKind, value: shiftAction{target: state(38)}}
	table[state(33)][token.ClassLit] = action{kind: shiftActionKind, value: shiftAction{target: state(36)}}
	table[state(33)][token.ClassNum] = action{kind: shiftActionKind, value: shiftAction{target: state(37)}}
	table[state(34)][token.ClassSemicolon] = action{kind: shiftActionKind, value: shiftAction{target: state(35)}}
	for _, term := range reduceTerms[state(35)] {
		table[state(35)][term] = action{kind: reduceActionKind, value: reduceAction{target: 14}}
	}
	table[state(36)][token.ClassSemicolon] = action{kind: reduceActionKind, value: reduceAction{target: 15}}
	table[state(37)][token.ClassSemicolon] = action{kind: reduceActionKind, value: reduceAction{target: 16}}
	table[state(38)][token.ClassSemicolon] = action{kind: reduceActionKind, value: reduceAction{target: 17}}
	table[state(39)][token.ClassReceive] = action{kind: shiftActionKind, value: shiftAction{target: state(40)}}
	table[state(40)][token.ClassId] = action{kind: shiftActionKind, value: shiftAction{target: state(43)}}
	table[state(40)][token.ClassNum] = action{kind: shiftActionKind, value: shiftAction{target: state(44)}}
	table[state(41)][token.ClassSemicolon] = action{kind: shiftActionKind, value: shiftAction{target: state(42)}}
	for _, term := range reduceTerms[state(42)] {
		table[state(42)][term] = action{kind: reduceActionKind, value: reduceAction{target: 19}}
	}
	for _, term := range reduceTerms[state(43)] {
		table[state(43)][term] = action{kind: reduceActionKind, value: reduceAction{target: 22}}
	}
	for _, term := range reduceTerms[state(44)] {
		table[state(44)][term] = action{kind: reduceActionKind, value: reduceAction{target: 23}}
	}
	table[state(45)][token.ClassAOperator] = action{kind: shiftActionKind, value: shiftAction{target: state(46)}}
	table[state(45)][token.ClassSemicolon] = action{kind: reduceActionKind, value: reduceAction{target: 21}}
	table[state(46)][token.ClassId] = action{kind: shiftActionKind, value: shiftAction{target: state(43)}}
	table[state(46)][token.ClassNum] = action{kind: shiftActionKind, value: shiftAction{target: state(44)}}
	table[state(47)][token.ClassSemicolon] = action{kind: reduceActionKind, value: reduceAction{target: 20}}
	table[state(48)][token.ClassOpenP] = action{kind: shiftActionKind, value: shiftAction{target: state(49)}}
	table[state(49)][token.ClassId] = action{kind: shiftActionKind, value: shiftAction{target: state(43)}}
	table[state(49)][token.ClassNum] = action{kind: shiftActionKind, value: shiftAction{target: state(44)}}
	table[state(50)][token.ClassROperator] = action{kind: shiftActionKind, value: shiftAction{target: state(51)}}
	table[state(51)][token.ClassId] = action{kind: shiftActionKind, value: shiftAction{target: state(43)}}
	table[state(51)][token.ClassNum] = action{kind: shiftActionKind, value: shiftAction{target: state(44)}}
	table[state(52)][token.ClassCloseP] = action{kind: reduceActionKind, value: reduceAction{target: 27}}
	table[state(53)][token.ClassCloseP] = action{kind: shiftActionKind, value: shiftAction{target: state(54)}}
	table[state(54)][token.ClassEntao] = action{shiftActionKind, shiftAction{target: state(55)}}
	for _, term := range reduceTerms[state(55)] {
		table[state(55)][term] = action{kind: reduceActionKind, value: reduceAction{target: 26}}
	}
	for _, term := range reduceTerms[state(57)] {
		table[state(57)][term] = action{kind: reduceActionKind, value: reduceAction{target: 25}}
	}
	for _, t := range []state{56, 58, 60, 62} {
		table[t][token.ClassId] = action{kind: shiftActionKind, value: shiftAction{target: state(39)}}
		table[t][token.ClassLeia] = action{kind: shiftActionKind, value: shiftAction{target: state(30)}}
		table[t][token.ClassEscreva] = action{kind: shiftActionKind, value: shiftAction{target: state(33)}}
		table[t][token.ClassSe] = action{kind: shiftActionKind, value: shiftAction{target: state(48)}}
		table[t][token.ClassFimse] = action{kind: shiftActionKind, value: shiftAction{target: state(64)}}
	}
	for _, term := range reduceTerms[state(59)] {
		table[state(59)][term] = action{kind: reduceActionKind, value: reduceAction{target: 28}}
	}
	for _, term := range reduceTerms[state(61)] {
		table[state(61)][term] = action{kind: reduceActionKind, value: reduceAction{target: 29}}
	}
	for _, term := range reduceTerms[state(63)] {
		table[state(63)][term] = action{kind: reduceActionKind, value: reduceAction{target: 30}}
	}
	for _, term := range reduceTerms[state(64)] {
		table[state(64)][term] = action{kind: reduceActionKind, value: reduceAction{target: 31}}
	}
	for _, term := range reduceTerms[state(66)] {
		table[state(66)][term] = action{kind: reduceActionKind, value: reduceAction{target: 33}}
	}
	for _, t := range []state{65, 67, 69, 71} {
		table[t][token.ClassId] = action{kind: shiftActionKind, value: shiftAction{target: state(39)}}
		table[t][token.ClassLeia] = action{kind: shiftActionKind, value: shiftAction{target: state(30)}}
		table[t][token.ClassEscreva] = action{kind: shiftActionKind, value: shiftAction{target: state(33)}}
		table[t][token.ClassSe] = action{kind: shiftActionKind, value: shiftAction{target: state(48)}}
		table[t][token.ClassFimrepita] = action{kind: shiftActionKind, value: shiftAction{target: state(73)}}
	}
	for _, term := range reduceTerms[state(68)] {
		table[state(68)][term] = action{kind: reduceActionKind, value: reduceAction{target: 35}}
	}
	for _, term := range reduceTerms[state(70)] {
		table[state(70)][term] = action{kind: reduceActionKind, value: reduceAction{target: 36}}
	}
	for _, term := range reduceTerms[state(72)] {
		table[state(72)][term] = action{kind: reduceActionKind, value: reduceAction{target: 37}}
	}
	for _, term := range reduceTerms[state(73)] {
		table[state(73)][term] = action{kind: reduceActionKind, value: reduceAction{target: 38}}
	}
	table[state(74)][token.ClassOpenP] = action{kind: shiftActionKind, value: shiftAction{target: state(75)}}
	table[state(75)][token.ClassId] = action{kind: shiftActionKind, value: shiftAction{target: state(43)}}
	table[state(75)][token.ClassLeia] = action{kind: shiftActionKind, value: shiftAction{target: state(44)}}
	table[state(76)][token.ClassCloseP] = action{kind: shiftActionKind, value: shiftAction{target: state(77)}}
	for _, term := range reduceTerms[state(77)] {
		table[state(77)][term] = action{kind: reduceActionKind, value: reduceAction{target: 34}}
	}

	return table
}

func GetMgolDetourTable() map[state]map[dictionary.NonTerminal]state {
	table := make(map[state]map[dictionary.NonTerminal]state)

	for i := state(0); i <= state(77); i++ {
		table[i] = make(map[dictionary.NonTerminal]state)
	}

	table[state(0)][dictionary.NTermP] = state(1)
	table[state(2)][dictionary.NTermV] = state(19)

	table[state(3)][dictionary.NTermLV] = state(4)
	table[state(3)][dictionary.NTermD] = state(17)
	table[state(3)][dictionary.NTermTIPO] = state(7)

	table[state(7)][dictionary.NTermL] = state(12)
	table[state(15)][dictionary.NTermL] = state(16)

	table[state(17)][dictionary.NTermLV] = state(18)
	table[state(17)][dictionary.NTermD] = state(17)
	table[state(17)][dictionary.NTermTIPO] = state(7)

	table[state(19)][dictionary.NTermA] = state(20)
	table[state(19)][dictionary.NTermES] = state(21)
	table[state(19)][dictionary.NTermCMD] = state(23)
	table[state(19)][dictionary.NTermCOND] = state(25)
	table[state(19)][dictionary.NTermR] = state(27)
	table[state(19)][dictionary.NTermCAB] = state(56)
	table[state(19)][dictionary.NTermCABR] = state(65)

	table[state(21)][dictionary.NTermA] = state(22)
	table[state(21)][dictionary.NTermES] = state(21)
	table[state(21)][dictionary.NTermCMD] = state(23)
	table[state(21)][dictionary.NTermCOND] = state(25)
	table[state(21)][dictionary.NTermR] = state(27)
	table[state(21)][dictionary.NTermCAB] = state(56)
	table[state(21)][dictionary.NTermCABR] = state(65)

	table[state(23)][dictionary.NTermA] = state(24)
	table[state(23)][dictionary.NTermES] = state(21)
	table[state(23)][dictionary.NTermCMD] = state(23)
	table[state(23)][dictionary.NTermCOND] = state(25)
	table[state(23)][dictionary.NTermR] = state(27)
	table[state(23)][dictionary.NTermCAB] = state(56)
	table[state(23)][dictionary.NTermCABR] = state(65)

	table[state(25)][dictionary.NTermA] = state(26)
	table[state(25)][dictionary.NTermES] = state(21)
	table[state(25)][dictionary.NTermCMD] = state(23)
	table[state(25)][dictionary.NTermCOND] = state(25)
	table[state(25)][dictionary.NTermR] = state(27)
	table[state(25)][dictionary.NTermCAB] = state(56)
	table[state(25)][dictionary.NTermCABR] = state(65)

	table[state(27)][dictionary.NTermA] = state(28)
	table[state(27)][dictionary.NTermES] = state(21)
	table[state(27)][dictionary.NTermCMD] = state(23)
	table[state(27)][dictionary.NTermCOND] = state(25)
	table[state(27)][dictionary.NTermR] = state(27)
	table[state(27)][dictionary.NTermCAB] = state(56)
	table[state(27)][dictionary.NTermCABR] = state(65)

	table[state(33)][dictionary.NTermARG] = state(34)
	table[state(40)][dictionary.NTermLD] = state(41)
	table[state(40)][dictionary.NTermOPRD] = state(45)
	table[state(46)][dictionary.NTermOPRD] = state(47)
	table[state(49)][dictionary.NTermOPRD] = state(50)
	table[state(49)][dictionary.NTermEXPR] = state(53)
	table[state(51)][dictionary.NTermOPRD] = state(52)

	table[state(56)][dictionary.NTermCAB] = state(56)
	table[state(56)][dictionary.NTermES] = state(58)
	table[state(56)][dictionary.NTermCMD] = state(60)
	table[state(56)][dictionary.NTermCOND] = state(62)
	table[state(56)][dictionary.NTermCP] = state(57)

	table[state(58)][dictionary.NTermES] = state(58)
	table[state(58)][dictionary.NTermCMD] = state(60)
	table[state(58)][dictionary.NTermCOND] = state(62)
	table[state(58)][dictionary.NTermCAB] = state(56)
	table[state(58)][dictionary.NTermCP] = state(59)

	table[state(60)][dictionary.NTermES] = state(58)
	table[state(60)][dictionary.NTermCMD] = state(60)
	table[state(60)][dictionary.NTermCOND] = state(62)
	table[state(60)][dictionary.NTermCAB] = state(56)
	table[state(60)][dictionary.NTermCP] = state(61)

	table[state(62)][dictionary.NTermES] = state(58)
	table[state(62)][dictionary.NTermCMD] = state(60)
	table[state(62)][dictionary.NTermCOND] = state(62)
	table[state(62)][dictionary.NTermCAB] = state(56)
	table[state(62)][dictionary.NTermCP] = state(63)

	table[state(65)][dictionary.NTermES] = state(67)
	table[state(65)][dictionary.NTermCMD] = state(69)
	table[state(65)][dictionary.NTermCOND] = state(71)
	table[state(65)][dictionary.NTermCAB] = state(56)
	table[state(65)][dictionary.NTermCPR] = state(66)

	table[state(67)][dictionary.NTermES] = state(67)
	table[state(67)][dictionary.NTermCMD] = state(69)
	table[state(67)][dictionary.NTermCOND] = state(71)
	table[state(67)][dictionary.NTermCAB] = state(56)
	table[state(67)][dictionary.NTermCPR] = state(68)

	table[state(69)][dictionary.NTermES] = state(67)
	table[state(69)][dictionary.NTermCMD] = state(69)
	table[state(69)][dictionary.NTermCOND] = state(71)
	table[state(69)][dictionary.NTermCAB] = state(56)
	table[state(69)][dictionary.NTermCPR] = state(70)

	table[state(71)][dictionary.NTermES] = state(67)
	table[state(71)][dictionary.NTermCMD] = state(69)
	table[state(71)][dictionary.NTermCOND] = state(71)
	table[state(71)][dictionary.NTermCAB] = state(56)
	table[state(71)][dictionary.NTermCPR] = state(72)

	table[state(75)][dictionary.NTermOPRD] = state(50)
	table[state(75)][dictionary.NTermEXPR] = state(76)

	return table
}
