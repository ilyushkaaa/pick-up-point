package middleware

import (
	"net/http"

	"homework/internal/events/model"
)

func (mw *Middleware) AccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newEvent := model.NewEvent(r.RemoteAddr, r.URL.Path, r.Method)
		err := mw.producer.SendMessage(newEvent)
		if err != nil {
			mw.logger.Errorf("error in writing new event into kafka: %s", err)
		}
		next.ServeHTTP(w, r)
	})
}
