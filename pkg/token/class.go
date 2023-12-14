package token

type Class uint8

const (
	ClassError     Class = 0
	ClassNum       Class = 1
	ClassLit       Class = 2
	ClassId        Class = 3
	ClassComment   Class = 4
	ClassEOF       Class = 5
	ClassROperator Class = 6
	ClassAOperator Class = 7
	ClassReceive   Class = 8
	ClassOpenP     Class = 9
	ClassCloseP    Class = 10
	ClassSemicolon Class = 11
	ClassComma     Class = 12
	ClassIgnore    Class = 13
	ClassInicio    Class = 14
	ClassVarinicio Class = 15
	ClassVarfim    Class = 16
	ClassEscreva   Class = 17
	ClassLeia      Class = 18
	ClassSe        Class = 19
	ClassEntao     Class = 20
	ClassFimse     Class = 21
	ClassRepita    Class = 22
	ClassFimrepita Class = 23
	ClassFim       Class = 24
	ClassInteiro   Class = 25
	ClassLiteral   Class = 26
	ClassReal      Class = 27
)

func (t Class) String() string {
	switch t {
	case ClassError:
		return "ERRO"
	case ClassNum:
		return "Num"
	case ClassLit:
		return "Lit"
	case ClassId:
		return "id"
	case ClassComment:
		return "Comentário"
	case ClassEOF:
		return "EOF"
	case ClassROperator:
		return "OPR"
	case ClassAOperator:
		return "OPM"
	case ClassReceive:
		return "RCB"
	case ClassOpenP:
		return "AB_P"
	case ClassCloseP:
		return "FC_P"
	case ClassSemicolon:
		return "PT_V"
	case ClassComma:
		return "VIR"
	case ClassIgnore:
		return "ignorar"
	case ClassInicio:
		return "inicio"
	case ClassVarinicio:
		return "varinicio"
	case ClassVarfim:
		return "varfim"
	case ClassEscreva:
		return "escreva"
	case ClassLeia:
		return "leia"
	case ClassSe:
		return "se"
	case ClassEntao:
		return "entao"
	case ClassFimse:
		return "fimse"
	case ClassRepita:
		return "repita"
	case ClassFimrepita:
		return "fimrepita"
	case ClassFim:
		return "fim"
	case ClassInteiro:
		return "inteiro"
	case ClassLiteral:
		return "literal"
	case ClassReal:
		return "real"
	}

	return "UNKNOWN"
}
