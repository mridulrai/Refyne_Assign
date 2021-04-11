package rest

import (
	"../storage"
)

// AccountListResponseData represents account list response data
type AccountListResponseData struct {
	Meta  Meta           `json:"meta"`
	Data  []UserResponse `json:"data"`
	Links Links          `json:"links"`
}

// Meta represents pagination meta data
type Meta struct {
	TotalPages int `json:"totalPages"`
}

// Links represents pagination links
type Links struct {
	Self     string `json:"self"`
	First    string `json:"first"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Last     string `json:"last"`
}

// AccountResponseData represents account response data
type UserResponseData struct {
	Data UserResponse `json:"data"`
}

// AccountResponse represents response for account
type UserResponse struct {
	ID     string `json:"id"`
	mobile string `json:"mobile"`
	Status string `json:"status"`
}

// ErrorResponseData represents error response data
type ErrorResponseData struct {
	Data ErrorResponse `json:"data"`
}

// ErrorResponse represents response for error
type ErrorResponse struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Code        string `json:"code"`
}

// mapFromModel maps fields from dao model to response
func (response *UserResponseData) mapFromModel(account storage.User) {
	response.Data.ID = account.ID
	response.Data.mobile = account.Mobile

}
