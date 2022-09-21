package common

import "strings"

const defaultCode = 1001

type CodeError struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type CodeErrorResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewCodeError(code int, msg string, data interface{}) error {
	return &CodeError{Code: code, Msg: msg, Data: data}
}

func NewDefaultError(msg string) error {
	var errMsg string
	errorMsgs := strings.Split(msg, "=")
	if len(errorMsgs) > 2 {
		errMsg = strings.TrimSpace(errorMsgs[len(errorMsgs)-1])
	} else {
		errMsg = msg
	}
	return NewCodeError(defaultCode, errMsg, "")
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Result() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
		Data: e.Data,
	}
}
