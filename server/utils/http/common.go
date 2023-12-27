package http

import (
	"fmt"
)

type BasicAuth struct {
	NeedBasicAuth bool
	Username      string
	Password      string
}

type Header struct {
	Key   string
	Value string
}

type SendAck struct {
	Code int    `json:"code"`
	Ok   bool   `json:"ok"`
	Data string `json:"msg"`
}

func (s SendAck) String() string {
	return fmt.Sprintf("Code: %d, Ok: %t, Data: %s", s.Code, s.Ok, s.Data)
}

func (s SendAck) Error() string {
	//TODO implement me
	panic("implement me")
}
