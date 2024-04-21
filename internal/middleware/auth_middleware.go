package middleware

import (
	"net/http"
	"os"

	"homework/pkg/hash"
	"homework/pkg/response"
)

func (mw *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			mw.logger.Errorf("password or/and username are not passed or passed incorrectly")
			response.MarshallAndWriteResponse(w, response.Error{Err: "auth data was not passed/passed incorrectly"}, http.StatusUnauthorized, mw.logger)
			return
		}

		isSuccessAuth, err := isAuthDataCorrect(username, password)
		if err != nil {
			mw.logger.Errorf("error in getting hash password: %v", err)
			response.MarshallAndWriteResponse(w, response.Error{Err: response.ErrInternal.Error()}, http.StatusInternalServerError, mw.logger)
			return
		}
		if !isSuccessAuth {
			mw.logger.Errorf("wrong username/password passed")
			response.MarshallAndWriteResponse(w, response.Error{Err: "username/password is wrong"}, http.StatusUnauthorized, mw.logger)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func isAuthDataCorrect(username, password string) (bool, error) {
	hashedPassword, err := hash.GetHash(password)
	if err != nil {
		return false, err
	}
	return username == os.Getenv("USERNAME") && hashedPassword == os.Getenv("PASSWORD"), nil
}
