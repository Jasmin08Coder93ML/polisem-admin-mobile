package middleware

import (
	"net/http"
	"strings"
)

// AuthMiddleware - JWT barlagy
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		// Tokeniň bardygyny we "Bearer " bilen başlaýandygyny barla
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			http.Error(w, "Forbidden: Invalid token", http.StatusForbidden)
			return
		}

		// Bu ýerde JWT-ni parse edip, ulanyjynyň IDsini almak bolar
		// next - indiki funksiýany çagyrýar
		next.ServeHTTP(w, r)
	})
}
