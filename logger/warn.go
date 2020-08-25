package logger

type ServiceWarning struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}