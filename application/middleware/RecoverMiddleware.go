package middleware

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"runtime/debug"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecoverMiddleware struct{}

func NewRecoverMiddleware() *RecoverMiddleware {
	return &RecoverMiddleware{}
}

func (m *RecoverMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				logx.WithContext(r.Context()).Errorf("[PANIC] %v\nStack: %s", rec, debug.Stack())
				httpx.WriteJson(w, http.StatusInternalServerError, nil)
			}
		}()
		next(w, r)
	}
}
