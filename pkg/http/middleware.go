package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/stovenn/dataimpact/pkg/token"
)

type CtxKey string

const (
	authHeaderKey         = "authorization"
	authTypeBearer        = "bearer"
	authPayloadKey CtxKey = "auth_payload"
)

func authMiddleware(maker token.Maker) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get(authHeaderKey)
			if len(authHeader) == 0 {
				err := errors.New("authorization header is not provided")
				handleError(w, http.StatusUnauthorized, err)
				return
			}

			fields := strings.Fields(authHeader)
			if len(fields) < 2 {
				_ = errors.New("invalid authorization header format")
				return
			}

			authType := strings.ToLower(fields[0])
			if authType != authTypeBearer {
				err := fmt.Errorf("unsupported authorization type %s", authType)
				handleError(w, http.StatusUnauthorized, err)
				return
			}

			accessToken := fields[1]
			payload, err := maker.VerifyToken(accessToken)
			if err != nil {
				handleError(w, http.StatusUnauthorized, err)
				return
			}

			ctx := context.WithValue(r.Context(), authPayloadKey, payload)
			req := r.Clone(ctx)
			next.ServeHTTP(w, req)
		})
	}
}
