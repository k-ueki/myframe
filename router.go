package myframe

type (
	HandlerFunc func(*Context)

	Route struct {
		Method string
		Path   string
		Name   string
		Func   HandlerFunc
	}
)
