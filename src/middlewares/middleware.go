package middlewares

import (
	"net/http"
	"time"

	"Liup.Gateway.Golang/src/utils"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)

		utils.LogRequest(r, start)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				utils.WriteErrorResponse(w, http.StatusInternalServerError, err.(string))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
