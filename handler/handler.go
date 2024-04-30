package handler

import (
	"clean-arc/model"
	"clean-arc/service"
	"encoding/json"
	"net/http"
)

type Handler struct {
	Service service.IService
}

type IHandler interface {
	HandleCalculator(w http.ResponseWriter, r *http.Request)
}

// init factory
func NewHandler(svc service.IService) IHandler {
	return &Handler{
		Service: svc,
	}
}

// controller / handler
func (h *Handler) HandleCalculator(w http.ResponseWriter, r *http.Request) {
	// cuma berfokus ke kebutuhan external (request dan response)
	// logic services
	// get request body

	switch r.Method {
	case "POST":
		var reqBody model.RequestCalculator

		// decode request body
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := h.Service.Calculator(reqBody.A, reqBody.B, reqBody.Operator)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// response json response
		var resData = model.ResponseCalculator{
			Result: res,
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resData)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
