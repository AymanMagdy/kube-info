package logger

type ServiceError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}