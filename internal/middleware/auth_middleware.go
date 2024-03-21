package middleware

import (
	"net/http"
	"os"

	"homework/pkg/response"
)

func (mw *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			mw.logger.Errorf("password or/and username are not passed or passed incorrectly")
			err := response.WriteResponse(w, []byte(`{"error":"auth data was not passed/passed incorrectly"}`), http.StatusUnauthorized)
			if err != nil {
				mw.logger.Errorf("error in writing response: %v", err)
			}
			return
		}

		if !isAuthDataCorrect(username, password) {
			mw.logger.Errorf("wrong username/password passed")
			err := response.WriteResponse(w, []byte(`{"error":"username/password is wrong"}`), http.StatusUnauthorized)
			if err != nil {
				mw.logger.Errorf("error in writing response: %v", err)
			}
			return
		}
		next.ServeHTTP(w, r)
	})
}

func isAuthDataCorrect(username, password string) bool {
	return username == os.Getenv("username") && password == os.Getenv("password")
}
