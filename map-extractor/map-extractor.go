package mapextractor

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

/**
{
	"attr1": {
		"attr2": [
			{
				item1: {
					attr3: ""
				}
			},
			{

			}
		]
	}
}

attr2.attr2[0].attr3
**/
type MapExtractor struct {
}

func NewMapExtractor() *MapExtractor {
	return &MapExtractor{}
}

func (n MapExtractor) GetValue(key string, param map[string]interface{}) (interface{}, error) {
	keys := strings.Split(key, ".")
	return n.GetValueKeys(keys, key, param)
}

func (n MapExtractor) GetValueKeys(keys []string, keyOri string, param map[string]interface{}) (interface{}, error) {
	if len(keys) == 0 {
		return nil, fmt.Errorf("no value found for key: %s", keyOri)
	}

	if len(keys) == 1 {

		isArr, keyParsed, idx, err := n.checkIfKeyArray(keys[0])
		if err != nil {
			return nil, err
		}

		if param[keyParsed] == nil {
			return nil, fmt.Errorf("key %s not found", keyParsed)
		}

		if isArr {
			if reflect.TypeOf(param[keyParsed]).Kind() != reflect.Slice {
				return nil, fmt.Errorf("incompatible type: map value is not an array for key %s", keyParsed)
			}

			value, err := n.getSliceValue(param[keyParsed].([]interface{}), idx)
			if err != nil {
				return nil, err
			}

			return value, nil
		}

		return param[keys[0]], nil
	}

	isArr, keyStr, idx, err := n.checkIfKeyArray(keys[0])
	if err != nil {
		return nil, err
	}

	if isArr {
		//check map type
		if reflect.TypeOf(param[keyStr]).Kind() != reflect.Slice {
			return nil, errors.New("incompatible type, map value is not array")
		}

		arr := param[keyStr].([]interface{})

		if len(arr) <= idx {
			return nil, fmt.Errorf("index %d out of range for key %s", idx, keyStr)
		}

		if reflect.TypeOf(arr[idx]).Kind() == reflect.Map {
			return n.GetValueKeys(keys[1:], keyOri, arr[idx].(map[string]interface{}))
		} else {
			return nil, fmt.Errorf("incompatible type, arr value key %s is not map", keyOri)
		}
	} else {
		if param[keyStr] == nil {
			return nil, fmt.Errorf("no value found for key: %s", keyOri)
		}

		if reflect.TypeOf(param[keyStr]).Kind() != reflect.Map {
			return nil, fmt.Errorf("incompatible type, arr value key %s is not map", keyOri)
		}

		return n.GetValueKeys(keys[1:], keyOri, param[keyStr].(map[string]interface{}))
	}
}

func (n MapExtractor) getSliceValue(arr []interface{}, idx int) (interface{}, error) {
	if len(arr) < idx {
		return nil, fmt.Errorf("index %d out of range", idx)
	}

	return arr[idx], nil
}

//return isArray, key, index, and error if any
//example key ttr2[0] will return true, ttr2, 0, nil
func (n MapExtractor) checkIfKeyArray(key string) (bool, string, int, error) {
	keyRunes := []rune(key)

	var parsedKey []rune
	for i := 0; i < len(keyRunes); i++ {
		if keyRunes[i] == '[' {
			//check if valid array format attr2[0]
			if keyRunes[len(keyRunes)-1] != ']' {
				return false, "", 0, fmt.Errorf("invalid array key format: %s", key)
			}

			idxArraStr := string(keyRunes[i+1 : len(keyRunes)-1])
			idx, err := strconv.Atoi(idxArraStr)
			if err != nil {
				return false, "", 0, fmt.Errorf("invalid array key format: %s", key)
			}

			return true, string(parsedKey), idx, nil
		}
		parsedKey = append(parsedKey, keyRunes[i])
	}

	if len(parsedKey) == 0 {
		return false, "", 0, fmt.Errorf("invalid array key format: %s", key)
	}

	return false, string(parsedKey), 0, nil
}
