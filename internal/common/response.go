package common

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Msg    string `json:"message"`
	Data   any    `json:"data,omitempty"`
	Errors any    `json:"errors,omitempty"`
	Status int    `json:"-"`
}

func WriteJSON(w http.ResponseWriter, response *Response) {
	status := http.StatusOK
	if response.Status != 0 {
		status = response.Status
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func WriteError(w http.ResponseWriter, err error) {
	if v, ok := err.(validator.ValidationErrors); ok {
		WriteValidateError(w, v)
	} else {
		WriteJSON(w, &Response{
			Msg:    err.Error(),
			Errors: err.Error(),
			Status: http.StatusBadRequest,
		})
	}
}

func WriteValidateError(w http.ResponseWriter, err validator.ValidationErrors) {
	status := http.StatusBadRequest
	errors := map[string]string{}

	for _, e := range err {
		errMsg := ""
		if e.ActualTag() == "required" {
			errMsg = e.Field() + " is required"
		} else if e.ActualTag() == "min" {
			errMsg = e.Field() + " minimum length not met"
		} else if e.ActualTag() == "oneof" {
			errMsg = e.Field() + " must be a valid option"
		}
		errors[e.Field()] = errMsg
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&Response{
		Msg:    "invalid input",
		Errors: errors,
	})
}
