package model

import "fmt"

type Response interface {
	IsError() bool
	Error() string
}

type BaseResponse struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

func (r BaseResponse) IsError() bool {
	return r.Code != 0
}

func (r BaseResponse) Error() string {
	return fmt.Sprintf("%d:%s", r.Code, r.Msg)
}
