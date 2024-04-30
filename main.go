package main

import (
	"clean-arc/handler"
	"clean-arc/service"
	"fmt"
	"net/http"
)

// clean arc -> independent
// contract interface -> error juga independent

func main() {
	// wiring process
	// setup service
	svc := service.NewService()

	// setup handler
	handler := handler.NewHandler(svc)

	http.HandleFunc("/calculate", handler.HandleCalculator)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

// factory pattern

// aku mau bikin 1 handler dengan logic sebagai berikut
// calculator, method POST
// ngirim data berupa json
// 3 => a, b, operator
// hasilnya akan dikembalikan ke response
