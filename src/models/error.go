package models

import (
	"encoding/json"
	"errors"
)

type ErrorResponse struct {
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	RequestId    string `json:"requestId,omitempty"`
}

var (
	EmptyGuidError       = errors.New("guid is empty")
	ClassificationsError = errors.New("classification is nil")
	NoDataProvidedError  = errors.New("no data provided")
	NoAttributesProvided = errors.New("no attributes provided")
	EmptyTypeNameError   = errors.New("typeName is empty")
)

func (e ErrorResponse) UnmarshalJson(data []byte) (*ErrorResponse, error) {
	err := json.Unmarshal(data, &e)
	return &e, err
}
