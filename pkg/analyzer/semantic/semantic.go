package semantic

import (
	"fmt"
	"strings"

	"github.com/arthurdelarge/simple-compiler/pkg/data"
	"github.com/arthurdelarge/simple-compiler/pkg/dictionary"
	"github.com/arthurdelarge/simple-compiler/pkg/token"
)

const (
	HEADER = true
	BODY   = false
)

type Semantic struct {
	stack          *data.Stack[*token.Token]
	conditionStack *data.Stack[string]
	header         strings.Builder
	body           strings.Builder
	tknTipo        *token.Token
	tabCounter     int
	tmpCounter     int
}

func NewSemantic() *Semantic {
	sem := &Semantic{
		stack:          data.CreateStack[*token.Token](),
		conditionStack: data.CreateStack[string](),
		tmpCounter:     0,
		tabCounter:     2,
	}

	sem.header.WriteString("#include<stdio.h>\ntypedef char literal[256];\nvoid main(void) {\n")

	return sem
}

func (s *Semantic) Tab(header bool) {
	for i := 0; i < s.tabCounter; i++ {
		if header {
			s.header.WriteString(" ")
		} else {
			s.body.WriteString(" ")
		}
	}
}

func (s *Semantic) IncreaseTab() {
	s.tabCounter += 2
}

func (s *Semantic) DecreaseTab() {
	s.tabCounter -= 2
}

func (s *Semantic) PrintStack() {
	for _, v := range s.stack.Array() {
		fmt.Printf("%+v, ", *v)
	}
	fmt.Printf("\n")
}

func (s *Semantic) StackPopCount(count int) {
	for i := 0; i < count; i++ {
		s.stack.Pop()
	}
}

func (s *Semantic) Evaluate(prod dictionary.Production) error {

	switch prod.ID() {
	case 2:
		s.body.WriteString("}\n")
	case 6:
		s.StackPopCount(prod.RightSideSize())
		s.stack.Push(token.NewToken(token.ClassNonTerminal, "D", token.TypeNull))
	case 7:
		ntL, _ := s.stack.Pop()
		s.stack.Pop()
		id, _ := s.stack.Pop()

		if id.GetType() != token.TypeNull {
			s.PushToken(token.NewToken(token.ClassNonTerminal, "L", ntL.GetType()))
			return fmt.Errorf("variável %s já foi declarada anteriormente", id.GetLexeme())
		}
		id.SetType(ntL.GetType())
		s.PushToken(token.NewToken(token.ClassNonTerminal, "L", ntL.GetType()))

		s.Tab(BODY)
		s.body.WriteString(fmt.Sprintf("%s %s;\n", id.GetType().String(), id.GetLexeme()))
	case 8:
		id, _ := s.stack.Pop()

		if id.GetType() != token.TypeNull {
			s.PushToken(token.NewToken(token.ClassNonTerminal, "L", s.tknTipo.GetType()))
			return fmt.Errorf("variável %s já foi declarada anteriormente", id.GetLexeme())
		}
		id.SetType(s.tknTipo.GetType())
		s.PushToken(token.NewToken(token.ClassNonTerminal, "L", id.GetType()))

		s.body.WriteString(fmt.Sprintf(" %s;\n", id.GetLexeme()))
	case 9:
		s.stack.Pop()
		s.tknTipo = token.NewToken(token.ClassNonTerminal, "TIPO", token.TypeInteger)
		s.stack.Push(s.tknTipo)

		s.Tab(BODY)
		s.body.WriteString("int")
	case 10:
		s.stack.Pop()
		s.tknTipo = token.NewToken(token.ClassNonTerminal, "TIPO", token.TypeReal)
		s.stack.Push(s.tknTipo)

		s.Tab(BODY)
		s.body.WriteString("double")
	case 11:
		s.stack.Pop()
		s.tknTipo = token.NewToken(token.ClassNonTerminal, "TIPO", token.TypeLiteral)
		s.stack.Push(s.tknTipo)

		s.Tab(BODY)
		s.body.WriteString("literal")
	case 13:
		s.stack.Pop()
		id, _ := s.stack.Pop()
		s.stack.Pop()

		s.Tab(BODY)
		s.body.WriteString("scanf(\"%")
		if id.GetType() == token.TypeLiteral {
			s.body.WriteString(fmt.Sprintf("s\", %s);\n", id.GetLexeme()))
		} else if id.GetType() == token.TypeInteger {
			s.body.WriteString(fmt.Sprintf("d\", &%s);\n", id.GetLexeme()))
		} else if id.GetType() == token.TypeReal {
			s.body.WriteString(fmt.Sprintf("lf\", &%s);\n", id.GetLexeme()))
		} else {
			return fmt.Errorf("variável %s não declarada", id.GetLexeme())
		}

	case 14:
		s.stack.Pop()
		arg, _ := s.stack.Pop()
		s.stack.Pop()

		s.Tab(BODY)
		s.body.WriteString(fmt.Sprintf("printf(%s);\n", arg.GetLexeme()))
	case 15:
		lit, _ := s.stack.Pop()
		s.stack.Push(token.NewToken(token.ClassNonTerminal, lit.GetLexeme(), lit.GetType()))
	case 16:
		num, _ := s.stack.Pop()
		arg := fmt.Sprintf("\"%s\"", num.GetLexeme())
		s.stack.Push(token.NewToken(token.ClassNonTerminal, arg, num.GetType()))
	case 17:
		id, _ := s.stack.Pop()

		arg := "\"%"
		if id.GetType() == token.TypeLiteral {
			arg += fmt.Sprintf("s\", %s", id.GetLexeme())
			s.stack.Push(token.NewToken(token.ClassNonTerminal, arg, id.GetType()))
		} else if id.GetType() == token.TypeInteger {
			arg += fmt.Sprintf("d\", %s", id.GetLexeme())
			s.stack.Push(token.NewToken(token.ClassNonTerminal, arg, id.GetType()))
		} else if id.GetType() == token.TypeReal {
			arg += fmt.Sprintf("lf\", %s", id.GetLexeme())
			s.stack.Push(token.NewToken(token.ClassNonTerminal, arg, id.GetType()))
		} else {
			s.PushToken(token.NewToken(token.ClassNonTerminal, "\"ARG\"", token.TypeLiteral))
			return fmt.Errorf("variável %s não declarada", id.GetLexeme())
		}

	case 19:
		s.stack.Pop()
		ld, _ := s.stack.Pop()
		s.stack.Pop()
		id, _ := s.stack.Pop()

		if id.GetType() == token.TypeNull {
			return fmt.Errorf("variável %s não declarada", id.GetLexeme())
		} else if id.GetType() != ld.GetType() {
			return fmt.Errorf("tipos diferentes para atribuição")
		}

		s.Tab(BODY)
		s.body.WriteString(fmt.Sprintf("%s = %s;\n", id.GetLexeme(), ld.GetLexeme()))
	case 20:
		oprd1, _ := s.stack.Pop()
		opm, _ := s.stack.Pop()
		oprd2, _ := s.stack.Pop()

		if oprd1.GetType() != oprd2.GetType() || oprd2.GetType() == token.TypeLiteral {
			s.stack.Push(oprd1)
			return fmt.Errorf("Operandos com tipos incompatíveis")
		}

		tmp := fmt.Sprintf("T%d", s.tmpCounter)
		s.tmpCounter++
		s.stack.Push(token.NewToken(token.ClassId, tmp, oprd1.GetType()))

		s.header.WriteString(fmt.Sprintf("  %s %s;\n", oprd1.GetType().String(), tmp))
		s.Tab(BODY)
		s.body.WriteString(fmt.Sprintf("%s = %s %s %s;\n", tmp, oprd2.GetLexeme(), opm.GetLexeme(), oprd1.GetLexeme()))
	case 21:
		oprd, _ := s.stack.Pop()
		s.stack.Push(token.NewToken(token.ClassNonTerminal, oprd.GetLexeme(), oprd.GetType()))
	case 22:
		id, _ := s.stack.Pop()
		s.stack.Push(token.NewToken(token.ClassNonTerminal, id.GetLexeme(), id.GetType()))
		if id.GetType() == token.TypeNull {
			return fmt.Errorf("variável %s não declarada", id.GetLexeme())
		}
	case 23:
		num, _ := s.stack.Pop()
		s.stack.Push(token.NewToken(token.ClassNonTerminal, num.GetLexeme(), num.GetType()))
	case 25:
		s.stack.Pop()
		s.stack.Pop()
		s.conditionStack.Pop()

		s.DecreaseTab()
		s.Tab(BODY)
		s.body.WriteString("}\n")
	case 26:
		s.stack.Pop()
		s.stack.Pop()
		exp_r, _ := s.stack.Pop()
		s.stack.Pop()
		s.stack.Pop()

		s.Tab(BODY)
		s.body.WriteString(fmt.Sprintf("if (%s) {\n", exp_r.GetLexeme()))
		s.IncreaseTab()
	case 27:
		oprd1, _ := s.stack.Pop()
		opr, _ := s.stack.Pop()
		oprd2, _ := s.stack.Pop()

		if oprd1.GetType() != oprd2.GetType() || oprd2.GetType() == token.TypeLiteral {
			s.stack.Push(oprd1)
			return fmt.Errorf("Operandos com tipos incompatíveis")
		}

		tmp := fmt.Sprintf("T%d", s.tmpCounter)
		s.tmpCounter++
		s.stack.Push(token.NewToken(token.ClassId, tmp, token.TypeInteger))

		s.header.WriteString(fmt.Sprintf("  int %s;\n", tmp))
		condition := fmt.Sprintf("%s = %s %s %s;\n", tmp, oprd2.GetLexeme(), opr.GetLexeme(), oprd1.GetLexeme())
		s.Tab(BODY)
		s.body.WriteString(condition)
		s.conditionStack.Push(condition)

	case 33:
		condition, _ := s.conditionStack.Pop()
		s.Tab(BODY)
		s.body.WriteString(condition)
		s.DecreaseTab()
		s.Tab(BODY)
		s.body.WriteString("}\n")
	case 34:
		s.stack.Pop()
		exp_r, _ := s.stack.Pop()
		s.stack.Pop()
		s.stack.Pop()
		s.Tab(BODY)
		s.body.WriteString(fmt.Sprintf("while (%s) {\n", exp_r.GetLexeme()))
		s.IncreaseTab()
	}

	return nil
}

func (s *Semantic) PushToken(t *token.Token) {
	s.stack.Push(t)
}

func (s *Semantic) Generate() string {
	return s.header.String() + s.body.String()
}
