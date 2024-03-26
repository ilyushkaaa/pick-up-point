package response

import (
	"net/http"

	"go.uber.org/zap"
)

func WriteResponse(w http.ResponseWriter, body []byte, statusCode int, logger *zap.SugaredLogger) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Connection", "keep-alive")
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		logger.Errorf("error in writing response: %v", err)
	}
}

func WriteMarshalledResponse(w http.ResponseWriter, body []byte, err error, logger *zap.SugaredLogger) {
	if err != nil {
		logger.Errorf("error in marshalling response: %v", err)
		WriteResponse(w, []byte(`{"error":"internal error"}`), http.StatusInternalServerError, logger)
		return
	}
	WriteResponse(w, body, http.StatusOK, logger)
}
