package datatype

import (
	"errors"
	"fmt"
)

type DataType string

const (
	BOOLEAN DataType = "BOOLEAN"
	STRING  DataType = "STRING"
	LONG    DataType = "LONG"
	FLOAT   DataType = "FLOAT"
	MAP_KEY DataType = "MAP_KEY"
)

func GetDataType(input string) DataType {
	//check is input string
	resp := IsStr(input)
	if resp {
		return STRING
	}

	//check is input boolean
	_, err := IsBool(input)
	if err == nil {
		return BOOLEAN
	}

	//check if input long
	_, err = IsLong(input)
	if err == nil {
		return LONG
	}

	//heck if input float
	_, err = IsFloat(input)
	if err == nil {
		return FLOAT
	}

	//if there are non of above then set to map key
	return MAP_KEY
}

func ValidateDataType(input1 string, input2 string) (DataType, error) {
	input1DataType := GetDataType(input1)
	input2DataType := GetDataType(input2)

	if input1DataType == MAP_KEY {
		errMsg := fmt.Sprintf("%s data type not found", input1)
		return "", errors.New(errMsg)
	}

	if input2DataType == MAP_KEY {
		errMsg := fmt.Sprintf("%s data type not found", input2)
		return "", errors.New(errMsg)
	}

	//compare input data type
	if input1DataType != input2DataType {
		errMsg := fmt.Sprintf("%s and %s is not compatible", input1, input2)
		return "", errors.New(errMsg)
	}

	return input1DataType, nil
}
