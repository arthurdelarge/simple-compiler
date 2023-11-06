package lexical

type SymbolTable struct {
	table map[string]Token
}

func CreateSymbolTable() *SymbolTable {
	st := &SymbolTable{
		table: make(map[string]Token, 0),
	}

	st.loadReservedWords()

	return st
}

func (st *SymbolTable) GetTokenFor(lexeme string) Token {
	token, found := st.table[lexeme]
	if !found {
		token = *NewToken(TokenClassId, lexeme, TokenTypeNull)
		st.table[lexeme] = token
	}

	return token
}

func (st *SymbolTable) loadReservedWords() {
	st.table["inicio"] = *NewToken(TokenClassInicio, "inicio", TokenTypeNull)
	st.table["varinicio"] = *NewToken(TokenClassVarinicio, "varinicio", TokenTypeNull)
	st.table["varfim"] = *NewToken(TokenClassVarfim, "varfim", TokenTypeNull)
	st.table["escreva"] = *NewToken(TokenClassEscreva, "escreva", TokenTypeNull)
	st.table["leia"] = *NewToken(TokenClassLeia, "leia", TokenTypeNull)
	st.table["se"] = *NewToken(TokenClassSe, "se", TokenTypeNull)
	st.table["entao"] = *NewToken(TokenClassEntao, "entao", TokenTypeNull)
	st.table["fimse"] = *NewToken(TokenClassFimse, "fimse", TokenTypeNull)
	st.table["repita"] = *NewToken(TokenClassRepita, "repita", TokenTypeNull)
	st.table["fimrepita"] = *NewToken(TokenClassFimrepita, "fimrepita", TokenTypeNull)
	st.table["fim"] = *NewToken(TokenClassFim, "fim", TokenTypeNull)
	st.table["inteiro"] = *NewToken(TokenClassInteiro, "inteiro", TokenTypeNull)
	st.table["literal"] = *NewToken(TokenClassLiteral, "literal", TokenTypeNull)
	st.table["real"] = *NewToken(TokenClassReal, "real", TokenTypeNull)
}
