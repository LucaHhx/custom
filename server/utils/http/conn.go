package http

import (
	"fmt"
	"io"
	"net/http"
)

type SendConn struct {
	*Options
}

func GetHttpConn(opts ...OptionFunc) *SendConn {
	o := &Options{}
	for _, opt := range opts {
		opt(o)
	}
	return &SendConn{Options: o}
}

func (s *SendConn) AddOption(opts ...OptionFunc) {
	for _, opt := range opts {
		opt(s.Options)
	}
}

func (s *SendConn) GetUrl() string {
	return fmt.Sprintf("%s%s", s.Url, s.Path)
}

func (s *SendConn) GetRequest() (req *http.Request, err error) {
	s.AddOption(SetHeader("Content-Type", "application/json"))
	req, err = http.NewRequest(s.Method, s.GetUrl(), s.Data)
	if err != nil {
		return nil, err
	}
	for _, header := range s.Headers {
		req.Header.Set(header.Key, header.Value)
	}
	if s.NeedBasicAuth {
		req.SetBasicAuth(s.Username, s.Password)
	}
	return
}

func (s *SendConn) Send() (ack *SendAck, err error) {
	var (
		req  *http.Request
		resp *http.Response
	)
	req, err = s.GetRequest()
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return &SendAck{
		Code: resp.StatusCode,
		Ok:   resp.StatusCode == http.StatusOK,
		Data: string(body),
	}, nil
}
