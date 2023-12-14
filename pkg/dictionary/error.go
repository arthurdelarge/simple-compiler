package dictionary

type ErrorCode byte

const (
	Unknown           ErrorCode = 0
	InvalidNumber     ErrorCode = 1
	IncompleteLiteral ErrorCode = 2
	IncompleteComment ErrorCode = 3
)
