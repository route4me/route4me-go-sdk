package utils

import "errors"

// ErrOperationFailed is returned as an error if a specific operation fails with unknown error
var ErrOperationFailed = errors.New("Operation has not succedeed with an unknown error.")

type StatusResponse struct {
	Status bool `json:"status,omitempty"`
}
