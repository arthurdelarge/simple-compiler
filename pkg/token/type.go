package token

type Type uint8

const (
	TypeNull    Type = 0
	TypeInteger Type = 1
	TypeReal    Type = 2
	TypeLiteral Type = 3
)

func (t Type) String() string {
	switch t {
	case TypeNull:
		return "Nulo"
	case TypeInteger:
		return "Inteiro"
	case TypeReal:
		return "Real"
	case TypeLiteral:
		return "Literal"
	}

	return "UNKNOWN"
}
