package service

import "errors"

type IService interface {
	Calculator(a, b int, operator string) (int, error)
}

// service /usecase
type service struct{}

func NewService() IService {
	return &service{}
}

// logic process calculator
func (s *service) Calculator(a, b int, operator string) (int, error) {
	// logic
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	case "%":
		return a % b, nil
	}

	return 0, errors.New("invalid operator")
}
