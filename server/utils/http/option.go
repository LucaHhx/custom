package http

import (
	"bytes"
	"io"
)

type Options struct {
	Url     string
	Path    string
	Method  string
	Data    io.Reader
	Headers []Header
	BasicAuth
}

type OptionFunc func(*Options)

func SetUrl(url string) OptionFunc {
	return func(o *Options) {
		o.Url = url
	}
}

func SetPath(path string) OptionFunc {
	return func(o *Options) {
		o.Path = path
	}
}

func SetBasicAuth(username, password string) OptionFunc {
	return func(o *Options) {
		o.NeedBasicAuth = true
		o.Username = username
		o.Password = password
	}
}

func SetData(data []byte) OptionFunc {
	return func(o *Options) {
		o.Data = bytes.NewBuffer(data)
	}
}

func SetMethod(method string) OptionFunc {
	return func(o *Options) {
		o.Method = method
	}
}

func SetHeaders(headers []Header) OptionFunc {
	return func(o *Options) {
		o.Headers = headers
	}
}

func SetHeader(key, value string) OptionFunc {
	return func(o *Options) {
		o.Headers = append(o.Headers, Header{
			Key:   key,
			Value: value,
		})
	}
}
