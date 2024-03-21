package response

import "net/http"

func WriteResponse(w http.ResponseWriter, body []byte, statusCode int) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Connection", "keep-alive")
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	return err
}
