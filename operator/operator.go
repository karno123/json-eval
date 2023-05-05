package operator

import (
	"errors"
	"fmt"
)

type Operator struct {
	Symbol     string
	Presedence int
}

func GetAllOperators() []Operator {
	var operators []Operator
	operators = append(operators, Operator{"<", 1})
	operators = append(operators, Operator{"<=", 1})
	operators = append(operators, Operator{">", 1})
	operators = append(operators, Operator{">=", 1})
	operators = append(operators, Operator{"==", 2})
	operators = append(operators, Operator{"!=", 2})
	operators = append(operators, Operator{"&&", 3})
	operators = append(operators, Operator{"||", 4})

	return operators
}

func IsOperator(input string) bool {
	operators := GetAllOperators()
	for _, val := range operators {
		if val.Symbol == input {
			return true
		}
	}

	return false
}

func GetOperator(symbol string) (Operator, error) {
	operators := GetAllOperators()
	for _, val := range operators {
		if val.Symbol == symbol {
			return val, nil
		}
	}

	return Operator{}, errors.New(symbol + " not found")
}

func GetOperatorPrecedence(symbol string) (int, error) {
	operators := GetAllOperators()
	for _, val := range operators {
		if val.Symbol == symbol {
			return val.Presedence, nil
		}
	}

	msg := fmt.Sprintf("%s operator not found", symbol)
	return 0, errors.New(msg)
}
