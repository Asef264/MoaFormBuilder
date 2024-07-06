package errHnadler

import "errors"

type MyError struct {
	MainError      error       `json:"mainError"`
	HttpStatusCode int         `json:"httpStatusCode"`
	Detail         interface{} `json:"detail"`
}

func NewMyError(errorMessage string, httpStatusCode int) MyError {
	return MyError{
		MainError:      errors.New(errorMessage),
		HttpStatusCode: httpStatusCode,
	}
}

func (me MyError) WithDetail(detail interface{}) MyError {
	me.Detail = detail
	return me
}
