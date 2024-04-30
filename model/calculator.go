package model

type RequestCalculator struct {
	A        int    `json:"a"`
	B        int    `json:"b"`
	Operator string `json:"operator"`
}

type ResponseCalculator struct {
	Result int `json:"result"`
}
