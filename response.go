package http

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data,omitempty"`
	Message    string      `json:"message,omitempty"`
}

func NewResponse(w http.ResponseWriter, response *Response) {
	b, _ := json.Marshal(&response)
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(response.StatusCode)
	w.Write(b)
}
func NewResponseSuccessWithStatus(w http.ResponseWriter, status int, data interface{}) {
	NewResponse(w, &Response{
		StatusCode: status,
		Data:       data,
	})
}
func NewResponseSuccess(w http.ResponseWriter, data interface{}) {
	NewResponseSuccessWithStatus(w, http.StatusOK, data)
}

func NewResponseErrorWithStatus(w http.ResponseWriter, status int, message string) {
	NewResponse(w, &Response{
		StatusCode: status,
		Message:    message,
	})
}
func NewResponseError(w http.ResponseWriter, message string) {
	NewResponseErrorWithStatus(w, http.StatusBadRequest, message)
}
