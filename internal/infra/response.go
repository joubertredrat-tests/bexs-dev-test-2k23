package infra

type RequestValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}
