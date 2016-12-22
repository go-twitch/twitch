package twitch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	errorResponse := &ErrorResponse{
		Response: r,
	}
	json.Unmarshal(data, errorResponse)
	return errorResponse
}

type ErrorResponse struct {
	Response *http.Response `json:"-"`
	ErrorStr string         `json:"error"`
	Status   int            `json:"status"`
	Message  string         `json:"message"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("twitch: %v", e.Message)
}
