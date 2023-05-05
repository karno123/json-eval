package datatype

import (
	"errors"
	"strconv"
)

func IsLong(var1 string) (bool, error) {
	_, err := strconv.ParseInt(var1, 10, 64)
	if err != nil {
		return false, err
	}

	return true, err
}

func CompareLong(var1 int64, operator string, var2 int64) (bool, error) {
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

func CompareLongStr(var1 string, operator string, var2 string) (bool, error) {
	var1Val, err := strconv.ParseInt(var1, 10, 64)
	if err != nil {
		return false, err
	}

	var2Val, err := strconv.ParseInt(var2, 10, 64)
	if err != nil {
		return false, err
	}

	return CompareFloat(float64(var1Val), operator, float64(var2Val))
}
