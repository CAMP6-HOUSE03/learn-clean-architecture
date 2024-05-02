package service

import (
	"clean-arc/model"
	"clean-arc/repository"
	"time"
)

type IService interface {
	Calculator(a, b int, operator string) (int, error)
	GetHistoryCalculator() []model.HistoryCalculator
}

// service /usecase
type service struct {
	repository repository.IRepostory
}

func NewService(repository repository.IRepostory) IService {
	return &service{
		repository: repository,
	}
}

// logic process calculator
func (s *service) Calculator(a, b int, operator string) (int, error) {
	var result int

	// logic
	switch operator {
	case "+": // add
		result = a + b
	case "-": // minus
		result = a - b
	case "*": // multiple
		result = a * b
	case "/": // divide
		result = a / b
	case "%": // modulus
		result = a % b
	}

	// insert data
	err := s.repository.InsertData(a, b, operator, result, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (s *service) GetHistoryCalculator() []model.HistoryCalculator {
	return s.repository.GetListData()
}
