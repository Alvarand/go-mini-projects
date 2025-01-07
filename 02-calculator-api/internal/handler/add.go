package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var errorBadValues = errors.New("one of numbers is invalid")

type addRequest struct {
	Number1 *int `json:"number1"`
	Number2 *int `json:"number2"`
}

func (a addRequest) checkValues() error {
	if a.Number1 == nil || a.Number2 == nil {
		return errorBadValues
	}
	return nil
}

func (a addRequest) sum() int {
	return *a.Number1 + *a.Number2
}

type addResponse struct {
	Result int `json:"result"`
}

func (h handler) Add(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var requestBody addRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, fmt.Sprintf("failed to decode add: %s", err), http.StatusBadRequest)
		return
	}

	// check input values
	if err := requestBody.checkValues(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := addResponse{
		Result: requestBody.sum(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response: %s", err), http.StatusInternalServerError)
	}
}
