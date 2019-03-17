package v3

import "fmt"

const ErrorOverDraw = 10020
type OverDrawError struct {
	msg string
	code int
	desc string
}

func NewError(code int) error {
	switch code {
	case ErrorOverDraw:
		return &OverDrawError{
			msg: "balance over draw!",
			code: 10020,
			desc: "you can not withdraw money more than you have",
		}
	default:
		return &OverDrawError{
			msg: "balance over draw!",
			code: 10020,
			desc: "you can not withdraw money more than you have",
		}
	}
}

func (o OverDrawError) Error() string {
	return fmt.Sprintf("balance over draw error, msg: %s, code: %d ", o.msg, o.code)
}
