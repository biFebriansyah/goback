package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/biFebriansyah/goback/src/helpers"
)

func AuthTest(role ...string) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-type", "application/json")

			var header string
			var valid bool = false

			if header = r.Header.Get("Authorization"); header == "" {
				helpers.New("header not provide", 401, false).Send(w)
				return
			}

			if !strings.Contains(header, "Bearer") {
				helpers.New("invalid header type", 401, false).Send(w)
				return
			}

			token := strings.Replace(header, "Bearer ", "", -1)

			checkTokens, err := helpers.CheckToken(token)
			if err != nil {
				helpers.New(err.Error(), 201, true).Send(w)
				return
			}

			for _, rl := range role {
				if rl == checkTokens.Role {
					valid = true
				}
			}

			if !valid {
				helpers.New("you not have permission to accsess", 401, false).Send(w)
				return
			}

			log.Println("Auth Middleware Pass")

			// share context to controller
			ctx := context.WithValue(r.Context(), "username", checkTokens.Username)

			// Serve the next handler
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AuthWithRole(role ...string) Adapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-type", "application/json")

			var header string
			var valid bool = false

			if header = r.Header.Get("Authorization"); header == "" {
				helpers.New("header not provide", 401, false).Send(w)
				return
			}

			if !strings.Contains(header, "Bearer") {
				helpers.New("invalid header type", 401, false).Send(w)
				return
			}

			token := strings.Replace(header, "Bearer ", "", -1)

			checkTokens, err := helpers.CheckToken(token)
			if err != nil {
				helpers.New(err.Error(), 201, true).Send(w)
				return
			}

			for _, rl := range role {
				if rl == checkTokens.Role {
					valid = true
				}
			}

			if !valid {
				helpers.New("you not have permission to accsess", 401, false).Send(w)
				return
			}

			log.Println("Auth Middleware Pass")

			// share context to controller
			ctx := context.WithValue(r.Context(), "username", checkTokens.Username)

			// Serve the next handler
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
