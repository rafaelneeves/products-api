package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"
)

// BasicAuthMiddleware verifica o cabeçalho Authorization
func BasicAuthMiddleware(username, password string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Decodifica as credenciais em Base64
			payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(authHeader, "Basic "))
			if err != nil || !validateCredentials(string(payload), username, password) {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Passa a requisição adiante
			next.ServeHTTP(w, r)
		})
	}
}

func validateCredentials(encoded string, expectedUser, expectedPass string) bool {
	parts := strings.SplitN(encoded, ":", 2)
	if len(parts) != 2 {
		return false
	}
	user, pass := parts[0], parts[1]
	return user == expectedUser && pass == expectedPass
}
