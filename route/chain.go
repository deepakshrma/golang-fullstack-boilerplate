package route

import "net/http"

type Middleware func(http.Handler) http.Handler
type Chain []Middleware

func Combine(middlewares ...Middleware) Chain {
	var slice Chain
	return append(slice, middlewares...)
}
func (c Chain) Then(originalHandler http.Handler) http.Handler {
	if originalHandler == nil {
		originalHandler = http.DefaultServeMux
	}
	for i := range c {
		// Same as to m1(m2(m3(originalHandler)))
		originalHandler = c[len(c)-1-i](originalHandler)
	}
	return originalHandler
}
