package automaton

import (
	"fmt"
	"github.com/arthurdelarge/simple-compiler/pkg/token"
)

type state byte

func (s state) rejectMessage(tclass token.Class) string {

	var expected string
	switch s {
	case state(0):
		expected = fmt.Sprintf("%s", token.ClassInicio.String())
	case state(1), state(20), state(22), state(24), state(26), state(28), state(29):
		expected = fmt.Sprintf("%s", token.ClassEOF.String())
	case state(2):
		expected = fmt.Sprintf("%s", token.ClassVarinicio.String())
	case state(3), state(14), state(17):
		expected = fmt.Sprintf(
			"%s, %s, %s, %s",
			token.ClassInteiro.String(),
			token.ClassReal.String(),
			token.ClassLiteral.String(),
			token.ClassVarfim.String(),
		)
	case state(4), state(11), state(18), state(19), state(21), state(23), state(25), state(27), state(59), state(61):
		expected = fmt.Sprintf(
			"%s, %s, %s, %s, %s, %s",
			token.ClassId.String(),
			token.ClassLeia.String(),
			token.ClassEscreva.String(),
			token.ClassSe.String(),
			token.ClassRepita.String(),
			token.ClassFim.String(),
		)
	case state(63), state(64), state(66), state(68), state(70), state(72), state(73):
		expected = fmt.Sprintf(
			"%s, %s, %s, %s, %s, %s",
			token.ClassId.String(),
			token.ClassLeia.String(),
			token.ClassEscreva.String(),
			token.ClassSe.String(),
			token.ClassRepita.String(),
			token.ClassFim.String(),
		)
	case state(6), state(12), state(16), state(31), state(34), state(36), state(37), state(38), state(41), state(47):
		expected = fmt.Sprintf("%s", token.ClassSemicolon.String())
	case state(7), state(8), state(9), state(10), state(15), state(30):
		expected = fmt.Sprintf("%s", token.ClassId.String())
	case state(13):
		expected = fmt.Sprintf("%s, %s", token.ClassSemicolon.String(), token.ClassComma.String())
	case state(32), state(35), state(42), state(57):
		expected = fmt.Sprintf(
			"%s, %s, %s, %s, %s, %s, %s, %s",
			token.ClassId.String(),
			token.ClassLeia.String(),
			token.ClassEscreva.String(),
			token.ClassSe.String(),
			token.ClassRepita.String(),
			token.ClassFim.String(),
			token.ClassFimse.String(),
			token.ClassFimrepita.String(),
		)
	case state(33):
		expected = fmt.Sprintf("%s, %s, %s", token.ClassId.String(), token.ClassLit.String(), token.ClassNum.String())
	case state(39):
		expected = fmt.Sprintf("%s", token.ClassReceive.String())
	case state(40), state(46), state(49), state(51):
		expected = fmt.Sprintf("%s, %s", token.ClassId.String(), token.ClassNum.String())
	case state(43), state(44):
		expected = fmt.Sprintf("%s, %s, %s, %s", token.ClassSemicolon.String(), token.ClassAOperator.String(), token.ClassROperator.String(), token.ClassCloseP.String())
	case state(45):
		expected = fmt.Sprintf("%s, %s", token.ClassSemicolon.String(), token.ClassAOperator.String())
	case state(48), state(74):
		expected = fmt.Sprintf("%s", token.ClassOpenP.String())
	case state(50):
		expected = fmt.Sprintf("%s", token.ClassROperator.String())
	case state(52), state(53), state(76):
		expected = fmt.Sprintf("%s", token.ClassCloseP.String())
	case state(54):
		expected = fmt.Sprintf("%s", token.ClassEntao.String())
	case state(55), state(58), state(60), state(62):
		expected = fmt.Sprintf("%s, %s, %s, %s, %s",
			token.ClassId.String(),
			token.ClassLeia.String(),
			token.ClassEscreva.String(),
			token.ClassSe.String(),
			token.ClassFimse.String())
	case state(56):
		expected = fmt.Sprintf("%s", token.ClassFimse.String())
	case state(65):
		expected = fmt.Sprintf("%s", token.ClassFimrepita.String())
	case state(67), state(69), state(71), state(77):
		expected = fmt.Sprintf("%s, %s, %s, %s, %s",
			token.ClassId.String(),
			token.ClassLeia.String(),
			token.ClassEscreva.String(),
			token.ClassSe.String(),
			token.ClassFimrepita.String())
	case state(75):
		expected = fmt.Sprintf("%s, %s", token.ClassId.String(), token.ClassLeia.String())
	}

	return fmt.Sprintf("Esperando {%s} recebeu {%s}", expected, tclass.String())
}
