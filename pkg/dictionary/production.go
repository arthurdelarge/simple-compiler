package dictionary

type Production struct {
	id            byte
	rule          string
	left          NonTerminal
	rightSideSize byte
}

func (p Production) String() string {
	return p.rule
}

func (p Production) ID() byte {
	return p.id
}

func (p Production) RightSideSize() int {
	return int(p.rightSideSize)
}

func (p Production) LeftSide() NonTerminal {
	return p.left
}

func GetMgolProductions() map[byte]*Production {
	return map[byte]*Production{
		1:  {1, "P' -> P", NTermPI, 1},
		2:  {2, "P -> inicio V A", NTermP, 3},
		3:  {3, "V -> varinicio LV", NTermV, 2},
		4:  {4, "LV -> D LV", NTermLV, 2},
		5:  {5, "LV -> varfim ;", NTermLV, 2},
		6:  {6, "D -> TIPO L ;", NTermD, 3},
		7:  {7, "L -> id , L", NTermL, 3},
		8:  {8, "L -> id", NTermL, 1},
		9:  {9, "TIPO -> inteiro", NTermTIPO, 1},
		10: {10, "TIPO -> real", NTermTIPO, 1},
		11: {11, "TIPO -> literal", NTermTIPO, 1},
		12: {12, "A -> ES A", NTermA, 2},
		13: {13, "ES -> leia id ;", NTermES, 3},
		14: {14, "ES -> escreva ARG ;", NTermES, 3},
		15: {15, "ARG -> lit", NTermARG, 1},
		16: {16, "ARG -> num", NTermARG, 1},
		17: {17, "ARG -> id", NTermARG, 1},
		18: {18, "A -> CMD A", NTermA, 2},
		19: {19, "CMD -> id rcb LD ;", NTermCMD, 4},
		20: {20, "LD -> OPRD opm OPRD", NTermLD, 3},
		21: {21, "LD -> OPRD", NTermLD, 1},
		22: {22, "OPRD -> id", NTermOPRD, 1},
		23: {23, "OPRD -> num", NTermOPRD, 1},
		24: {24, "A -> COND A", NTermA, 2},
		25: {25, "COND -> CAB CP", NTermCOND, 2},
		26: {26, "CAB -> se ( EXP_R ) entao", NTermCAB, 5},
		27: {27, "EXP_R -> OPRD opr OPRD", NTermEXPR, 3},
		28: {28, "CP -> ES CP", NTermCP, 2},
		29: {29, "CP -> CMD CP", NTermCP, 2},
		30: {30, "CP -> COND CP", NTermCP, 2},
		31: {31, "CP -> fimse", NTermCP, 1},
		32: {32, "A -> R A", NTermA, 2},
		33: {33, "R -> CABR CPR", NTermR, 2},
		34: {34, "CABR -> repita ( EXP_R )", NTermCABR, 4},
		35: {35, "CPR -> ES CPR", NTermCPR, 2},
		36: {36, "CPR -> CMD CPR", NTermCPR, 2},
		37: {37, "CPR -> COND CPR", NTermCPR, 2},
		38: {38, "CPR -> fimrepita", NTermCPR, 1},
		39: {39, "A -> fim", NTermA, 1},
	}
}
