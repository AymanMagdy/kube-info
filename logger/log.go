package logger

type ServiceLogging struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}