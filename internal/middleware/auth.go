package middleware

import (
	"context"
	"go-web/internal"
	"go-web/pkg/security"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := security.CurrentUser(r)

		if err != nil {
			internal.JsonResponse("Undauthorized.", http.StatusUnauthorized, w)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
