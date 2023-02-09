package helper

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type ErrorJson struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"any"`
}

func ReadJSON(r *http.Request, data any) error {
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return err
	}

	err = json.NewDecoder(r.Body).Decode(&struct{}{})

	if err != io.EOF {
		return errors.New("body must have only a single JSON Value")
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	return err
}

func WriteErrorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload ErrorJson
	payload.Error = true
	payload.Message = err.Error()

	return WriteJSON(w, statusCode, payload)
}
