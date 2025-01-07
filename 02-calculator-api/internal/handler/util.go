package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var errorMissingValues = errors.New("one parameter is missing")

type baseRequest struct {
	Number1 *int `json:"number1"`
	Number2 *int `json:"number2"`
}

func (br baseRequest) checkValues() error {
	if br.Number1 == nil || br.Number2 == nil {
		return errorMissingValues
	}
	return nil
}

type baseResponse struct {
	Result int `json:"result"`
}

type request interface {
	checkValues() error
	execute() int
}

func handleRequest(w http.ResponseWriter, r *http.Request, requestBody request) {
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(requestBody); err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request: %s", err), http.StatusBadRequest)
		return
	}

	// check input values
	if err := requestBody.checkValues(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := baseResponse{
		Result: requestBody.execute(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response: %s", err), http.StatusInternalServerError)
	}
}
