package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type multiplyRequest struct {
	baseRequest
}

func (mr multiplyRequest) multiply() int {
	return *mr.Number1 * *mr.Number2
}

func (h handler) Multiply(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var requestBody multiplyRequest
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
		Result: requestBody.multiply(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response: %s", err), http.StatusInternalServerError)
	}

}
