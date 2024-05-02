package repository

import "clean-arc/model"

// kebutuhan mengolah data
// komunikasi langsung dengan database
// simpan di array

// nyimpen data
// update data
// hapus data

// factory

type IRepostory interface {
	InsertData(a, b int, operator string, result int, date string) error
	GetListData() []model.HistoryCalculator
}

type repository struct {
	// data
	data []model.HistoryCalculator
}

func NewRepository() IRepostory {
	return &repository{}
}

func (r *repository) InsertData(a, b int, operator string, result int, date string) error {
	r.data = append(r.data, model.HistoryCalculator{
		A:        a,
		B:        b,
		Operator: operator,
		Result:   result,
		Date:     date,
	})

	return nil
}

func (r *repository) GetListData() []model.HistoryCalculator {
	return r.data
}

// method
// insert nambah data
// get list data
