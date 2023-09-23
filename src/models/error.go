package models

import (
	"encoding/json"
	"errors"
)

type ErrorResponse struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	RequestId    string `json:"requestId"`
}

var (
	EmptyGuidError       = errors.New("guid is empty")
	ClassificationsError = errors.New("classification is nil")
	NoDataProvidedError  = errors.New("no data provided")
	NoAttributesProvided = errors.New("no attributes provided")
)

func (e ErrorResponse) UnmarshalJson(data []byte) (*ErrorResponse, error) {
	err := json.Unmarshal(data, &e)
	return &e, err
}
