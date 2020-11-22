package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/gufranmirza/imdb-api/models"
	"github.com/gufranmirza/imdb-api/models/authmodel"
	"github.com/gufranmirza/imdb-api/web/renderers"
)

type ctxKey int

const (
	ctxClaims ctxKey = iota
	ctxRefreshToken
)

// Logging middleware logs the incoming requests
func Logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t1 := time.Now().UTC()
			defer func() {
				requestID, ok := r.Context().Value(models.HdrRequestID).(string)
				if !ok {
					requestID = "unknown"
				}
				logger.Printf("%s: %s  Method: %s URL: %s RemoteAddr: %s UserAgent: %s Latency: %v ", models.HdrRequestID, requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent(), time.Since(t1))
			}()
			next.ServeHTTP(w, r)
		})
	}
}

// Tracing adds a TracingID to each requests
func Tracing(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(models.HdrRequestID)
			if requestID == "" {
				requestID = fmt.Sprintf("%d", time.Now().UnixNano())
			}
			ctx := context.WithValue(r.Context(), models.HdrRequestID, requestID)
			r.Header.Set(models.HdrRequestID, requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Authenticator is a default authentication middleware to enforce access from the
// Verifier middleware request context values. The Authenticator sends a 401 Unauthorized
// response for any unverified tokens and passes the good ones through.
func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			renderers.ErrorUnauthorized(w, r, authmodel.ErrInsufficientRights)
			return
		}

		if !token.Valid {
			renderers.ErrorUnauthorized(w, r, authmodel.ErrInsufficientRights)
			return
		}

		// Token is authenticated, parse claims
		var c authmodel.AppClaims
		err = c.ParseClaims(claims)
		if err != nil {
			renderers.ErrorUnauthorized(w, r, authmodel.ErrInsufficientRights)
			return
		}

		// Set AppClaims on context
		ctx := context.WithValue(r.Context(), ctxClaims, c)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequiresRole middleware restricts access to accounts having role parameter in their jwt claims.
func RequiresRole(role authmodel.Role) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			claims := ClaimsFromCtx(r.Context())
			if !hasRole(role, claims.Roles) {
				renderers.ErrorForbidden(w, r, authmodel.ErrInsufficientRights)
				return
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(hfn)
	}
}

func hasRole(role authmodel.Role, roles []authmodel.Role) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}

// ClaimsFromCtx retrieves the parsed AppClaims from request context.
func ClaimsFromCtx(ctx context.Context) authmodel.AppClaims {
	return ctx.Value(ctxClaims).(authmodel.AppClaims)
}
