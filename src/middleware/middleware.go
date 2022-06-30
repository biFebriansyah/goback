package middleware

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc
type Adapter func(http.Handler) http.Handler

func Do(hf http.HandlerFunc, middle ...Middleware) http.HandlerFunc {
	for _, m := range middle {
		hf = m(hf)
	}

	return hf
}

func Adapt(f func(http.ResponseWriter, *http.Request), adapters ...Adapter) http.Handler {
	var handler http.Handler
	for i := len(adapters); i > 0; i-- {
		handlers := http.HandlerFunc(f)
		handler = adapters[i-1](handlers)
	}
	return handler
}
