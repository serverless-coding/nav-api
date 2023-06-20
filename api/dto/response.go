package dto

type BaseRespose struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
