package jsoneval

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/karno123/json-eval/datatype"
	mapextractor "github.com/karno123/json-eval/map-extractor"
	"github.com/karno123/json-eval/operator"
	"github.com/karno123/json-eval/stack"
	"github.com/karno123/json-eval/tokenizer"
)

/*
Operator supported:
()
< <=, > >=
== !=
&&
||
for precedence and associativity will follow this list
https://en.cppreference.com/w/c/language/operator_precedence
*/

type JsonEvaluator struct {
	Tokenizer    tokenizer.Tokenizer
	MapExtractor mapextractor.MapExtractor
}

func NewJsonEvaluator() *JsonEvaluator {
	return &JsonEvaluator{Tokenizer: tokenizer.Tokenizer{},
		MapExtractor: *mapextractor.NewMapExtractor()}
}

func (m JsonEvaluator) EvaluateJson(expression string, jsonStr string) (bool, error) {
	var x map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &x)
	if err != nil {
		return false, errors.New("invalid json format")
	}

	return m.Evaluate(expression, x)
}

func (m JsonEvaluator) Evaluate(expression string, parameter map[string]interface{}) (bool, error) {
	tokens, err := m.Tokenizer.Tokenize(expression)
	if err != nil {
		return false, err
	}

	tokens, err = m.Tokenizer.InFixToPostFix(tokens)
	if err != nil {
		return false, err
	}

	return m.evaluateTokens(tokens, parameter)
}

func (m JsonEvaluator) evaluateTokens(tokens []tokenizer.Token, param map[string]interface{}) (bool, error) {
	if len(tokens) <= 0 {
		return false, errors.New("tokens can not be empty")
	}

	stack := stack.NewStack()
	for _, val := range tokens {
		if operator.IsOperator(val.Value) {
			rightOperand, err := stack.Pop()
			if err != nil {
				return false, errors.New("invalid syntax near " + val.Value)
			}

			leftOperand, err := stack.Pop()
			if err != nil {
				return false, errors.New("invalid syntax near " + val.Value)
			}

			result, err := m.Compare(leftOperand, val.Value, rightOperand, param)
			if err != nil {
				return false, err
			}

			stack.Push(strconv.FormatBool(result))
		} else {
			stack.Push(val.Value)
		}
	}

	result, err := stack.Pop()
	if err != nil {
		return false, errors.New("invalid syntax")
	}

	if !stack.IsEmpty() {
		return false, errors.New("invalid syntax")
	}

	boolResult, err := strconv.ParseBool(result)
	if err != nil {
		return false, errors.New("invalid syntax")
	}

	return boolResult, nil
}

func (m JsonEvaluator) Compare(leftOperand string, operator string, rightOperand string, parameter map[string]interface{}) (bool, error) {
	//check data type
	leftDataType := datatype.GetDataType(leftOperand)
	rightDataType := datatype.GetDataType(rightOperand)

	lefData := leftOperand
	if leftDataType == datatype.MAP_KEY {
		val, err := m.MapExtractor.GetValue(leftOperand, parameter)
		if val == nil {
			errMsg := fmt.Sprintf("key %s not found", leftOperand)
			return false, errors.New(errMsg)
		}

		if reflect.TypeOf(val).String() == "string" {
			lefData = fmt.Sprintf("\"%v\"", val)
		} else {
			lefData = fmt.Sprintf("%v", val)
		}

		if err != nil {
			return false, err
		}
	}

	rightData := rightOperand
	if rightDataType == datatype.MAP_KEY {
		val, err := m.MapExtractor.GetValue(rightOperand, parameter)
		if val == nil {
			errMsg := fmt.Sprintf("key %s not found", rightOperand)
			return false, errors.New(errMsg)
		}

		if reflect.TypeOf(val).String() == "string" {
			rightData = fmt.Sprintf("\"%v\"", val)
		} else {
			rightData = fmt.Sprintf("%v", val)
		}
		if err != nil {
			return false, err
		}
	}

	varType, err := datatype.ValidateDataType(lefData, rightData)
	if err != nil {
		return false, err
	}

	switch varType {
	case datatype.BOOLEAN:
		return datatype.CompareBoolStr(lefData, operator, rightData)
	case datatype.STRING:
		return datatype.CompareStr(lefData, operator, rightData)
	case datatype.FLOAT:
		return datatype.CompareFloatStr(lefData, operator, rightData)
	case datatype.LONG:
		return datatype.CompareFloatStr(lefData, operator, rightData)
	}

	errMsg := fmt.Sprintf("%s and %s invalid type", lefData, rightData)
	return false, errors.New(errMsg)
}
