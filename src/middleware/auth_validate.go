package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/biFebriansyah/goback/src/helpers"
)

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		headerToken := r.Header.Get("Authorization")

		if !strings.Contains(headerToken, "Bearer") {
			helpers.New("invalid header type", 401, false).Send(w)
			return
		}

		token := strings.Replace(headerToken, "Bearer ", "", -1)
		checkToken, err := helpers.CheckToken(token)
		if err != nil {
			helpers.New(err.Error(), 201, true).Send(w)
			return
		}

		if checkToken.Username == "" {
			helpers.New("Silahkan login kembali", 401, false).Send(w)
			return
		}

		ctx := context.WithValue(r.Context(), "username", checkToken)

		next.ServeHTTP(w, r.WithContext(ctx))
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

			// share context to controller
			ctx := context.WithValue(r.Context(), "username", checkTokens.Username)

			// Serve the next handler
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
