package middleware

import (
	"boilerplate/config"
	"boilerplate/util"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"time"
)

type ContextKey string

var ContextKeyRequest = ContextKey("requestContext")

type RequestContext struct {
	Log       *slog.Logger
	requestId string
}

func (rc *RequestContext) GetRequestId() string {
	return rc.requestId
}

func NewRequestContext(log *slog.Logger, r *http.Request) *RequestContext {
	requestId := r.Header.Get("X-Request-ID")
	if util.IsStringEmpty(requestId) {
		requestId = uuid.New().String()
	}
	return &RequestContext{
		requestId: requestId,
		Log:       slog.With("request_id", requestId), // change the key to request id or trace id, based on your need
	}
}

func UseContext(r *http.Request) *RequestContext {
	reqContext := r.Context().Value(ContextKeyRequest).(*RequestContext)
	return reqContext
}

func RequestContextID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ContextKeyRequest, NewRequestContext(config.Logger, r))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := UseContext(r)
		startTime := time.Now()
		ctx.Log.Info("Start of request", slog.String("method", r.Method), slog.String("url", r.URL.String()))
		defer func() {
			ctx.Log.Info("End of request", slog.String("method", r.Method), slog.String("url", r.URL.String()), slog.Int64("time", time.Since(startTime).Milliseconds()))
		}()
		next.ServeHTTP(w, r)
	})
}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := UseContext(r)
		defer func() {
			err := recover()
			if err != nil {
				ctx.Log.Error("internal server error", slog.String("error", fmt.Sprintf("%v", err)))

				jsonBody, _ := json.Marshal(map[string]string{
					"error": "There was an internal server error",
				})

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}

		}()

		next.ServeHTTP(w, r)

	})
}
