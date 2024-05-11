package middleware

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"time"
	"webapp/pkg/helpers/str"
)

type ContextKey string

var ContextKeyApp = ContextKey("applicationContext")

type ApplicationContext struct {
	Log       *slog.Logger
	requestId string
}

func (rc *ApplicationContext) GetRequestId() string {
	return rc.requestId
}

func NewRequestContext(r *http.Request) *ApplicationContext {
	requestId := r.Header.Get("X-Request-ID")
	if str.IsStringEmpty(requestId) {
		requestId = uuid.New().String()
	}
	return &ApplicationContext{
		requestId: requestId,
		Log:       slog.With("request_id", requestId), // change the key to request id or trace id, based on your need
	}
}

func UseAppContext(r *http.Request) *ApplicationContext {
	reqContext := r.Context().Value(ContextKeyApp).(*ApplicationContext)
	return reqContext
}

func ApplicationContextM(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ContextKeyApp, NewRequestContext(r))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := UseAppContext(r)
		startTime := time.Now()
		ctx.Log.Info("Start of request", slog.String("method", r.Method), slog.String("url", r.URL.String()))
		defer func() {
			ctx.Log.Info("End of request", slog.String("method", r.Method), slog.String("url", r.URL.String()), slog.Int64("time", time.Since(startTime).Milliseconds()))
		}()
		next.ServeHTTP(w, r)
	})
}
