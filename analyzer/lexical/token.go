package lexical

const (
	TokenTypeNull    tokenType = 0
	TokenTypeInteger tokenType = 1
	TokenTypeReal    tokenType = 2
	TokenTypeLiteral tokenType = 3

	TokenClassError     tokenClass = 0
	TokenClassNum       tokenClass = 1
	TokenClassLit       tokenClass = 2
	TokenClassId        tokenClass = 3
	TokenClassComment   tokenClass = 4
	TokenClassEOF       tokenClass = 5
	TokenClassROperator tokenClass = 6
	TokenClassAOperator tokenClass = 7
	TokenClassReceive   tokenClass = 8
	TokenClassOpenP     tokenClass = 9
	TokenClassCloseP    tokenClass = 10
	TokenClassSemicolon tokenClass = 11
	TokenClassComma     tokenClass = 12
	TokenClassIgnore    tokenClass = 13
	TokenClassInicio    tokenClass = 14
	TokenClassVarinicio tokenClass = 15
	TokenClassVarfim    tokenClass = 16
	TokenClassEscreva   tokenClass = 17
	TokenClassLeia      tokenClass = 18
	TokenClassSe        tokenClass = 19
	TokenClassEntao     tokenClass = 20
	TokenClassFimse     tokenClass = 21
	TokenClassRepita    tokenClass = 22
	TokenClassFimrepita tokenClass = 23
	TokenClassFim       tokenClass = 24
	TokenClassInteiro   tokenClass = 25
	TokenClassLiteral   tokenClass = 26
	TokenClassReal      tokenClass = 27
)

type tokenType uint8
type tokenClass uint8

type Token struct {
	class     tokenClass
	lexeme    string
	tokenType tokenType
}

func NewToken(class tokenClass, lexeme string, tType tokenType) *Token {
	return &Token{
		class:     class,
		lexeme:    lexeme,
		tokenType: tType,
	}
}

func (t *Token) GetClass() tokenClass {
	return t.class
}

func (t *Token) GetLexeme() string {
	if t.class == TokenClassEOF {
		return "EOF"
	}
	return t.lexeme
}

func (t *Token) GetType() tokenType {
	return t.tokenType
}

func (t tokenType) String() string {
	switch t {
	case TokenTypeNull:
		return "Nulo"
	case TokenTypeInteger:
		return "Inteiro"
	case TokenTypeReal:
		return "Real"
	case TokenTypeLiteral:
		return "Literal"
	}

	return "UNKNOWN"
}

func (t tokenClass) String() string {
	switch t {
	case TokenClassError:
		return "ERRO"
	case TokenClassNum:
		return "Num"
	case TokenClassLit:
		return "Lit"
	case TokenClassId:
		return "id"
	case TokenClassComment:
		return "Comentário"
	case TokenClassEOF:
		return "EOF"
	case TokenClassROperator:
		return "OPR"
	case TokenClassAOperator:
		return "OPM"
	case TokenClassReceive:
		return "RCB"
	case TokenClassOpenP:
		return "AB_P"
	case TokenClassCloseP:
		return "FC_P"
	case TokenClassSemicolon:
		return "PT_V"
	case TokenClassComma:
		return "VIR"
	case TokenClassIgnore:
		return "ignorar"
	case TokenClassInicio:
		return "inicio"
	case TokenClassVarinicio:
		return "varinicio"
	case TokenClassVarfim:
		return "varfim"
	case TokenClassEscreva:
		return "escreva"
	case TokenClassLeia:
		return "leia"
	case TokenClassSe:
		return "se"
	case TokenClassEntao:
		return "entao"
	case TokenClassFimse:
		return "fimse"
	case TokenClassRepita:
		return "repita"
	case TokenClassFimrepita:
		return "fimrepita"
	case TokenClassFim:
		return "fim"
	case TokenClassInteiro:
		return "inteiro"
	case TokenClassLiteral:
		return "literal"
	case TokenClassReal:
		return "real"
	}

	return "UNKNOWN"
}
