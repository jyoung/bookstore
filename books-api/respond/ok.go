package respond

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

// APIResponse is the interface to bridge APIData and APIDatum creating a single response object
type APIResponse interface {
	*APIData | *APIDatum
	// SetRequestID sets the Request Id on the response object
	SetRequestID(requestID string)
	// GetStatus returns the status on the response object
	GetStatus() int
}

// APIDatum is a single piece of data being returned from the API, i.e. a single object
type APIDatum struct {
	Data      any    `json:"data"`
	Status    int    `json:"status"`
	RequestID string `json:"requestId"`
}

func (d *APIDatum) SetRequestID(requestID string) {
	d.RequestID = requestID
}

func (d *APIDatum) GetStatus() int {
	return d.Status
}

// APIData is an array of data being returned from the API, i.e. a list of objects
type APIData struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"pagination"`
	Status     int        `json:"status"`
	RequestID  string     `json:"requestId"`
}

func (d *APIData) SetRequestID(requestID string) {
	d.RequestID = requestID
}

func (d *APIData) GetStatus() int {
	return d.Status
}

type Pagination struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

// Created returns an APIDatum with the passed in data and a status of 201
func Created(data any) *APIDatum {
	return &APIDatum{
		Data:   data,
		Status: http.StatusCreated,
	}
}

// Accepted returns an APIDatum with the passed in data and a status of 202
func Accepted(data any) *APIDatum {
	return &APIDatum{
		Data:   data,
		Status: http.StatusAccepted,
	}
}

// Single returns an APIDatum with the passed in data and a status of 200
func Single(data any) *APIDatum {
	return &APIDatum{
		Data:   data,
		Status: http.StatusOK,
	}
}

// List returns an APIData with the passed in data a status of 200 and pagination
func List(data any, total, offset, limit int) *APIData {
	return &APIData{
		Data: data,
		Pagination: Pagination{
			Total:  total,
			Offset: offset,
			Limit:  limit,
		},
		Status: http.StatusOK,
	}
}

// / Ok writes the response object to the response stream as JSON
func Ok[T APIResponse](w http.ResponseWriter, r *http.Request, data T) {
	if data == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	data.SetRequestID(middleware.GetReqID(r.Context()))

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(data.GetStatus())
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return
	}
}
