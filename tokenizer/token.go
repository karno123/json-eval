package tokenizer

type Token struct {
	Value string
}

func FromString(val string) Token {
	return Token{val}
}
