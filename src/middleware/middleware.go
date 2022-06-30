package middleware

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc
type Adapter func(http.Handler) http.Handler

func Handle(hf http.HandlerFunc, middle ...Middleware) http.HandlerFunc {
	for i := len(middle); i > 0; i-- {
		hf = middle[i-1](hf)
	}
	return hf
}

func Adapts(handler http.Handler, adapters ...Adapter) http.Handler {
	for i := len(adapters); i > 0; i-- {
		handler = adapters[i-1](handler)
	}
	return handler
}
