package view

type ErrorView struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func ParseError(message string, code int) *ErrorView {
	return &ErrorView{message, code}
}
