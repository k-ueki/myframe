package myframe

import (
	"net/http"
	"net/url"
	"sync"
)

type (
	Context interface {
	}

	context struct {
		request     *http.Request
		response    Response
		path        string
		paramNames  []string
		paramValues []string
		query       url.Values
		logger      Logger
		lock        sync.RWMutex
	}
)

func (c *context) Request() *http.Request {
	return c.request
}

func (c *context) Response() *Response {
	return &c.response
}

func (c *context) Path() string {
	return c.path
}

func (c *context) Param(name string) string {
	for i, p := range c.paramNames {
		if p == name {
			return c.paramValues[i]
		}
		if i >= len(c.paramValues) {
			break
		}
	}
	return ""
}
