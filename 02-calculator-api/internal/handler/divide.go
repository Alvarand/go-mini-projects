package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var errorBadDivisor = errors.New("divisor is zero")

type divideRequest struct {
	Dividend *int `json:"dividend"`
	Divisor  *int `json:"divisor"`
}

func (dr divideRequest) checkValues() error {
	if dr.Dividend == nil || dr.Divisor == nil {
		return errorMissingValues
	}
	if *dr.Divisor == 0 {
		return errorBadDivisor
	}
	return nil
}

func (dr divideRequest) divide() int {
	return *dr.Dividend / *dr.Divisor
}

func (h handler) Divide(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var requestBody divideRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request: %s", err), http.StatusBadRequest)
		return
	}

	// check input values
	if err := requestBody.checkValues(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := baseResponse{
		Result: requestBody.divide(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response: %s", err), http.StatusInternalServerError)
	}

}
