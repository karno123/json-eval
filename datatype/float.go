package datatype

import (
	"errors"
	"strconv"
)

func IsFloat(input string) (bool, error) {
	_, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return false, err
	}

	return true, nil
}

func CompareFloat(var1 float64, operator string, var2 float64) (bool, error) {
	if operator == "==" {
		return var1 == var2, nil
	}

	if operator == ">" {
		return var1 > var2, nil
	}

	if operator == ">=" {
		return var1 >= var2, nil
	}

	if operator == "<" {
		return var1 < var2, nil
	}

	if operator == "<=" {
		return var1 <= var2, nil
	}

	if operator == "!=" {
		return var1 != var2, nil
	}

	return false, errors.New("invalid operator")
}

func CompareFloatStr(var1 string, operator string, var2 string) (bool, error) {
	var1Val, err := strconv.ParseFloat(var1, 64)
	if err != nil {
		return false, err
	}

	var2Val, err := strconv.ParseFloat(var2, 64)
	if err != nil {
		return false, err
	}

	return CompareFloat(var1Val, operator, var2Val)
}
