package middleware

import (
	"context"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/http/rest/auth"
	"net/http"
)

// LoggedIn simple middleware to push value to the context
func LoggedIn(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		credentials, err := r.Cookie("credentials")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		auth := auth.NewAuthenticator("ok")
		t, err := auth.ParseToken(credentials.Value)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), "credentials", t)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
