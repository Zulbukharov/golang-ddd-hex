package middleware

import (
	"context"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/http/rest/auth"
	"net/http"
)

type Rules interface {
	LoggedIn(next http.Handler) http.Handler
}

type rules struct {
	auth auth.Authenticator
}

func NewRules(auth auth.Authenticator) Rules {
	return &rules{auth}
}

// LoggedIn simple middleware to push value to the context
func (m rules) LoggedIn(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		credentials, err := r.Cookie("credentials")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		t, err := m.auth.ParseToken(credentials.Value)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), "credentials", t)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
