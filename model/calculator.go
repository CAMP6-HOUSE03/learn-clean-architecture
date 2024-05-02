package model

type RequestCalculator struct {
	A        int    `json:"a" binding:"required"`
	B        int    `json:"b" binding:"required"` // default empty value int 0
	Operator string `json:"operator" binding:"required"`
}

type ResponseCalculator struct {
	Result int `json:"result"`
}

type ErrorResponse struct {
	// prop / field
	Error string `json:"error"`
}

func ErrResponse(msg string) ErrorResponse {
	// mapping
	return ErrorResponse{
		Error: msg,
	}
}

type HistoryCalculator struct {
	A        int    `json:"a"`
	B        int    `json:"b"`
	Operator string `json:"operator"`
	Result   int    `json:"result"`
	Date     string `json:"date"` // pakcage time
}
