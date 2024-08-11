package adapter

import (
	"net/http"

	"github.com/urfave/negroni"
)

func ToNegroni(chiMiddleware func(http.Handler) http.Handler) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		chiMiddleware(http.HandlerFunc(next)).ServeHTTP(w, r)
	}
}
