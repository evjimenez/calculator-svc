package operation

import "errors"

var ErrInvalidOperator = errors.New("ups it seems that you tried an unknown operator")

type Factory struct {
}

func NewFactory() *Factory {
	return &Factory{}
}

func (f Factory) CreateOperator(symbol string) (Operator, error) {

	switch symbol {
	case "+":
		return NewAddition(), nil
	case "-":
		return NewSubtraction(), nil
	case "*":
		return NewMultiplication(), nil
	case "/":
		return NewDivision(), nil
	default:
		return nil, ErrInvalidOperator
	}
}
