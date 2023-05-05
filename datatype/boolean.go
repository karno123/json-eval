package datatype

import (
	"errors"
	"fmt"
	"strconv"
)

func IsBool(var1 string) (bool, error) {
	if var1 == "true" {
		return true, nil
	}

	if var1 == "false" {
		return false, nil
	}
	errMsg := fmt.Sprintf("invalid boolean type : %s", var1)
	return false, errors.New(errMsg)
}

func CompareBool(var1 bool, operator string, var2 bool) (bool, error) {
	if operator == "==" {
		return var1 == var2, nil
	}

	if operator == "||" {
		return var1 || var2, nil
	}

	if operator == "&&" {
		return var1 && var2, nil
	}

	if operator == "!=" {
		return var1 != var2, nil
	}

	return false, errors.New("invalid operator")
}

func CompareBoolStr(var1 string, operator string, var2 string) (bool, error) {
	var1Value, err := strconv.ParseBool(var1)
	if err != nil {
		return false, err
	}

	var2Value, err := strconv.ParseBool(var2)
	if err != nil {
		return false, err
	}

	return CompareBool(var1Value, operator, var2Value)
}
