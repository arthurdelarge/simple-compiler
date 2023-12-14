package dictionary

type State byte

const (
	InitialState         State = 0
	FinalStatesNum1      State = 1
	FinalStatesNum2      State = 2
	FinalStatesNum3      State = 5
	FinalStatesNum4      State = 6
	FinalStatesNum5      State = 9
	FinalStateLit        State = 12
	FinalStateId         State = 13
	FinalStateComment    State = 16
	FinalStateEOF        State = 17
	FinalStateROperator1 State = 18
	FinalStateROperator2 State = 19
	FinalStateROperator3 State = 20
	FinalStateAOperator  State = 22
	FinalStateReceive    State = 21
	FinalStateOpenP      State = 23
	FinalStateCloseP     State = 24
	FinalStateSemicolon  State = 25
	FinalStateComma      State = 26
	FinalStateIgnore     State = 27
)

func GetFinalStates() []State {
	return []State{
		FinalStatesNum1, FinalStatesNum2, FinalStatesNum3, FinalStatesNum4, FinalStatesNum5,
		FinalStateLit, FinalStateId, FinalStateComment, FinalStateEOF, FinalStateROperator1,
		FinalStateROperator2, FinalStateROperator3, FinalStateAOperator, FinalStateReceive,
		FinalStateOpenP, FinalStateCloseP, FinalStateSemicolon, FinalStateComma, FinalStateIgnore,
	}
}
