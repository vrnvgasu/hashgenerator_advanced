package middlewares

import (
	"context"
	"net/http"

	guid "github.com/satori/go.uuid"
)

const RequestID = "requestID"

func ContextRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get("X-Request-ID")
		if reqID == "" {
			uuid := guid.Must(guid.NewV4(), nil)
			reqID = uuid.String()
		}

		ctx := context.WithValue(context.Background(), RequestID, reqID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
