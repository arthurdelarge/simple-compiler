package symboltable

import (
	"fmt"

	tkn "github.com/arthurdelarge/simple-compiler/pkg/token"
)

type SymbolTable struct {
	table map[string]*tkn.Token
}

func (st *SymbolTable) Print() {
	for key, val := range st.table {
		fmt.Printf("%s: Lex: %s | Class: %s | Type: %s\n", key, val.GetLexeme(), val.GetClass(), val.GetType())
	}
}

func CreateSymbolTable() *SymbolTable {
	st := &SymbolTable{
		table: make(map[string]*tkn.Token),
	}

	st.loadReservedWords()

	return st
}

func (st *SymbolTable) GetTokenFor(lexeme string) *tkn.Token {
	token, found := st.table[lexeme]
	if !found {
		token = tkn.NewToken(tkn.ClassId, lexeme, tkn.TypeNull)
		st.table[lexeme] = token
	}

	return token
}

func (st *SymbolTable) loadReservedWords() {
	st.table["inicio"] = tkn.NewToken(tkn.ClassInicio, "inicio", tkn.TypeNull)
	st.table["varinicio"] = tkn.NewToken(tkn.ClassVarinicio, "varinicio", tkn.TypeNull)
	st.table["varfim"] = tkn.NewToken(tkn.ClassVarfim, "varfim", tkn.TypeNull)
	st.table["escreva"] = tkn.NewToken(tkn.ClassEscreva, "escreva", tkn.TypeNull)
	st.table["leia"] = tkn.NewToken(tkn.ClassLeia, "leia", tkn.TypeNull)
	st.table["se"] = tkn.NewToken(tkn.ClassSe, "se", tkn.TypeNull)
	st.table["entao"] = tkn.NewToken(tkn.ClassEntao, "entao", tkn.TypeNull)
	st.table["fimse"] = tkn.NewToken(tkn.ClassFimse, "fimse", tkn.TypeNull)
	st.table["repita"] = tkn.NewToken(tkn.ClassRepita, "repita", tkn.TypeNull)
	st.table["fimrepita"] = tkn.NewToken(tkn.ClassFimrepita, "fimrepita", tkn.TypeNull)
	st.table["fim"] = tkn.NewToken(tkn.ClassFim, "fim", tkn.TypeNull)
	st.table["inteiro"] = tkn.NewToken(tkn.ClassInteiro, "inteiro", tkn.TypeNull)
	st.table["literal"] = tkn.NewToken(tkn.ClassLiteral, "literal", tkn.TypeNull)
	st.table["real"] = tkn.NewToken(tkn.ClassReal, "real", tkn.TypeNull)
}
