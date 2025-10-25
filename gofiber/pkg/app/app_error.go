package app

type AppError struct {
	Code int // bebas, service/repo tentuin
	Msg  string
}

func (e *AppError) Error() string {
	return e.Msg
}

func NewError(code int, msg string) *AppError {
	return &AppError{Code: code, Msg: msg}
}
