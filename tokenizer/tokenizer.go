package tokenizer

import (
	"errors"

	"github.com/karno123/json-eval/operator"
	"github.com/karno123/json-eval/stack"
)

type Tokenizer struct {
}

func NewTokenizer() Tokenizer {
	return Tokenizer{}
}

func (t Tokenizer) Tokenize(expression string) ([]Token, error) {

	var tokens []Token
	expRunes := []rune(expression)

	token := Token{}
	for i := 0; i < len(expRunes); i++ {
		if expRunes[i] == '>' {
			if len(token.Value) > 0 {
				tokenArr, err := t.SanitizeAndAppend(tokens, token)
				tokens = tokenArr
				if err != nil {
					return nil, errors.New("invalid syntax")
				}
				token = Token{}
			}

			if expRunes[i] == '>' && i+1 <= len(expRunes) && expRunes[i+1] == '=' {
				tokens = append(tokens, Token{">="})
				i++
				continue
			}
			tokens = append(tokens, Token{">"})
		} else if expRunes[i] == '<' {
			if len(token.Value) > 0 {
				tokenArr, err := t.SanitizeAndAppend(tokens, token)
				tokens = tokenArr
				if err != nil {
					return nil, errors.New("invalid syntax")
				}
				token = Token{}
			}

			if expRunes[i] == '<' && i+1 < len(expRunes) && expRunes[i+1] == '=' {
				tokens = append(tokens, Token{"<="})
				i++
				continue
			}
			tokens = append(tokens, Token{"<"})
		} else if expRunes[i] == '=' {

			if len(token.Value) > 0 {
				tokenArr, err := t.SanitizeAndAppend(tokens, token)
				tokens = tokenArr
				if err != nil {
					return nil, errors.New("invalid syntax")
				}
				token = Token{}
			}

			if i+1 < len(expRunes) && expRunes[i+1] == '=' {
				tokens = append(tokens, Token{"=="})
				i++
				continue
			}
			return nil, errors.New("invalid operator ==")
		} else if expRunes[i] == '&' {

			if len(token.Value) > 0 {
				tokenArr, err := t.SanitizeAndAppend(tokens, token)
				tokens = tokenArr
				if err != nil {
					return nil, errors.New("invalid syntax")
				}
				token = Token{}
			}

			if i+1 < len(expRunes) && expRunes[i+1] == '&' {
				tokens = append(tokens, Token{"&&"})
				i++
				continue
			}
			return nil, errors.New("invalid operator &&")
		} else if expRunes[i] == '|' {

			if len(token.Value) > 0 {
				tokenArr, err := t.SanitizeAndAppend(tokens, token)
				tokens = tokenArr
				if err != nil {
					return nil, errors.New("invalid syntax")
				}
				token = Token{}
			}

			if i+1 < len(expRunes) && expRunes[i+1] == '|' {
				tokens = append(tokens, Token{"||"})
				i++
				continue
			}
			return nil, errors.New("invalid operator ||")
		} else if expRunes[i] == '(' {

			if len(token.Value) > 0 {
				tokenArr, err := t.SanitizeAndAppend(tokens, token)
				tokens = tokenArr
				if err != nil {
					return nil, errors.New("invalid syntax")
				}
				token = Token{}
			}

			tokens = append(tokens, Token{"("})
			continue
		} else if expRunes[i] == ')' {

			if len(token.Value) > 0 {
				tokenArr, err := t.SanitizeAndAppend(tokens, token)
				tokens = tokenArr
				if err != nil {
					return nil, errors.New("invalid syntax")
				}
				token = Token{}
			}

			tokens = append(tokens, Token{")"})
			continue
		} else if expRunes[i] == '!' {
			if len(token.Value) > 0 {
				tokenArr, err := t.SanitizeAndAppend(tokens, token)
				tokens = tokenArr
				if err != nil {
					return nil, errors.New("invalid syntax")
				}
				token = Token{}
			}

			if i+1 < len(expRunes) && expRunes[i+1] == '=' {
				tokens = append(tokens, Token{"!="})
				i++
				continue
			}
		} else {
			token.Value = token.Value + string(expRunes[i])
		}
	}

	if len(token.Value) > 0 {
		tokenArr, err := t.SanitizeAndAppend(tokens, token)
		tokens = tokenArr
		if err != nil {
			return nil, errors.New("invalid syntax")
		}
		token = Token{}
	}

	return tokens, nil
}

//remove space from start and end
func (t Tokenizer) SanitizeAndAppend(tokens []Token, newToken Token) ([]Token, error) {
	if len(newToken.Value) <= 0 {
		return nil, errors.New("invalid token")
	}

	tokenValueRunes := []rune(newToken.Value)
	i := 0
	for i < len(tokenValueRunes) && tokenValueRunes[i] == ' ' {
		i++
	}

	tokenValueRunes = tokenValueRunes[i:]
	if len(tokenValueRunes) <= 0 {
		return tokens, nil
	}

	i = len(tokenValueRunes) - 1
	for tokenValueRunes[i] == ' ' && i > 0 {
		i--
	}

	tokenValueRunes = tokenValueRunes[0 : i+1]
	if len(tokenValueRunes) <= 0 {
		return tokens, nil
	}

	tokenVal := string(tokenValueRunes)

	tokens = append(tokens, Token{tokenVal})
	return tokens, nil
}

func (t Tokenizer) InFixToPostFix(tokens []Token) ([]Token, error) {
	if len(tokens) == 0 {
		return nil, errors.New("tokens can not be empty")
	}

	var output []Token
	stack := stack.NewStack()
	for _, val := range tokens {

		if val.Value == "(" {
			stack.Push(val.Value)
			continue
		}

		if val.Value == ")" {
			if stack.IsEmpty() {
				return nil, errors.New("invalid syntax")
			}

			top, err := stack.Top()
			if err != nil {
				return nil, errors.New("invalid syntax")
			}

			for !stack.IsEmpty() && top != "(" {

				stackItm, err := stack.Pop()
				if err != nil {
					return nil, err
				}

				output = append(output, FromString(stackItm))
				top, err = stack.Top()
				if err != nil {
					return nil, errors.New("invalid syntax")
				}
			}

			_, err = stack.Pop()
			if err != nil {
				return nil, errors.New("invalid syntax")
			}
			continue
		}

		if operator.IsOperator(val.Value) {

			if stack.IsEmpty() {
				stack.Push(val.Value)
			} else {
				prevOpStr, err := stack.Top()
				if err != nil {
					return nil, errors.New("invalid syntax")
				}

				if prevOpStr == "(" {
					stack.Push(val.Value)
					continue
				}

				prevOp, err := operator.GetOperator(prevOpStr)
				if err != nil {
					return nil, err
				}

				currentOp, err := operator.GetOperator(val.Value)
				if err != nil {
					return nil, err
				}

				if prevOp.Presedence <= currentOp.Presedence {
					output = append(output, FromString(prevOp.Symbol))
					_, err = stack.Pop()
					if err != nil {
						return nil, errors.New("invalid syntax")
					}
					stack.Push(val.Value)
				} else {
					stack.Push(val.Value)
				}
			}
		} else {
			output = append(output, val)
		}
	}

	if !stack.IsEmpty() {
		for !stack.IsEmpty() {
			val, err := stack.Pop()
			if err != nil {
				return nil, errors.New("invalid syntax")
			}
			output = append(output, FromString(val))
		}
	}

	return output, nil
}
