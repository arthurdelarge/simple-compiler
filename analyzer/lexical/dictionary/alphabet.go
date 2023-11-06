package dictionary

type Alphabet struct {
	digits  []byte
	letters []byte
	symbols []byte
}

func CreateAlphabet() *Alphabet {
	alphabet := &Alphabet{
		digits:  make([]byte, 0),
		letters: make([]byte, 0),
		symbols: make([]byte, 0),
	}

	for digit := byte('0'); digit <= byte('9'); digit++ {
		alphabet.digits = append(alphabet.digits, digit)
		alphabet.symbols = append(alphabet.symbols, digit)
	}

	for letter := byte('A'); letter <= byte('z'); letter++ {
		if letter == byte('Z')+1 {
			letter = byte('a')
		}

		alphabet.letters = append(alphabet.letters, letter)
		alphabet.symbols = append(alphabet.symbols, letter)
	}

	alphabet.symbols = append(alphabet.symbols, ',')
	alphabet.symbols = append(alphabet.symbols, ';')
	alphabet.symbols = append(alphabet.symbols, ':')
	alphabet.symbols = append(alphabet.symbols, '!')
	alphabet.symbols = append(alphabet.symbols, '?')
	alphabet.symbols = append(alphabet.symbols, '\\')
	alphabet.symbols = append(alphabet.symbols, '*')
	alphabet.symbols = append(alphabet.symbols, '+')
	alphabet.symbols = append(alphabet.symbols, '-')
	alphabet.symbols = append(alphabet.symbols, '/')
	alphabet.symbols = append(alphabet.symbols, '(')
	alphabet.symbols = append(alphabet.symbols, ')')
	alphabet.symbols = append(alphabet.symbols, '{')
	alphabet.symbols = append(alphabet.symbols, '}')
	alphabet.symbols = append(alphabet.symbols, '[')
	alphabet.symbols = append(alphabet.symbols, ']')
	alphabet.symbols = append(alphabet.symbols, '<')
	alphabet.symbols = append(alphabet.symbols, '>')
	alphabet.symbols = append(alphabet.symbols, '=')
	alphabet.symbols = append(alphabet.symbols, '.')
	alphabet.symbols = append(alphabet.symbols, '\'')
	alphabet.symbols = append(alphabet.symbols, '"')
	alphabet.symbols = append(alphabet.symbols, '_')
	alphabet.symbols = append(alphabet.symbols, ' ')

	return alphabet
}
