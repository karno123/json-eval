package datatype

import "errors"

func IsStr(var1 string) bool {
	runes := []rune(var1)
	if runes[0] == '"' && runes[len(runes)-1] == '"' {
		return true
	}

	return false
}

func CompareStr(var1 string, operator string, var2 string) (bool, error) {
	if operator == "==" {
		return var1 == var2, nil
	}

	if operator == "<" {
		return var1 < var2, nil
	}

	if operator == "<=" {
		return var1 <= var2, nil
	}

	if operator == ">" {
		return var1 > var2, nil
	}

	if operator == ">=" {
		return var1 >= var2, nil
	}

	if operator == "!=" {
		return var1 != var2, nil
	}

	return false, errors.New("invalid operator")
}
