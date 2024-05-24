package conversion

import (
	"calculator-svc/container"
	"calculator-svc/operation"
	"errors"
	"strconv"
	"strings"
)

var (
	ErrInvalidExpression     = errors.New("invalid expression")
	ErrInvalidNegativeNumber = errors.New("negative number not allowed")
)

type PostFix struct {
	expression []string
}

func (p *PostFix) String() string {
	return strings.Join(p.expression, "")
}

func (p *PostFix) Add(s string) {
	p.expression = append(p.expression, s)
}

func InfixToPostfix(s string) (PostFix, error) {
	//temArray := []string{"1", "2.5", "3", "/", "4", "*", "+"}
	//postfix := PostFix{expression: temArray}

	expr, err := tokenization(s)
	if err != nil {
		return PostFix{}, err
	}

	var stack container.Stack[operation.Operator]
	factory := operation.NewFactory()
	postfix := PostFix{}

	for _, v := range expr {
		if isOperand(v) {
			postfix.Add(v)
			continue
		}
		currentOperator, err := factory.CreateOperator(v)
		if err != nil {
			return PostFix{}, err
		}
		if stack.IsEmpty() {
			stack.Push(currentOperator)
			continue
		}
		for {
			top, err := stack.Peek()
			if err != nil {
				stack.Push(currentOperator)
				break
			}
			if top.IsGreaterOrEqual(currentOperator) {
				operator, err := stack.Pop()
				if err != nil {
					stack.Push(currentOperator)
					break
				}
				postfix.Add(operator.Symbol())
			} else {
				stack.Push(currentOperator)
				break
			}
		}
	}

	for {
		operator, err := stack.Pop()
		if err != nil {
			break
		}
		postfix.Add(operator.Symbol())
	}

	return postfix, nil
}

func isOperand(value string) bool {
	if _, err := strconv.ParseFloat(value, 64); err != nil {
		return false
	}
	return true
}

func tokenization(s string) ([]string, error) {

	if s == "" {
		return nil, ErrInvalidExpression
	}
	expression := strings.ReplaceAll(s, " ", "")

	var digit string
	var infixTokenized []string
	for _, v := range expression {
		if v >= '0' && v <= '9' || v == '.' {
			digit += string(v)
		} else {
			if digit != "" {
				infixTokenized = append(infixTokenized, digit)
				digit = ""
			}
			infixTokenized = append(infixTokenized, string(v))
		}
	}

	if digit != "" {
		infixTokenized = append(infixTokenized, digit)
	}

	return infixTokenized, nil
}
