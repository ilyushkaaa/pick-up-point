package delivery

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"homework/internal/order/delivery/dto"
	"homework/internal/order/service"
	"homework/internal/order/storage"
	"homework/pkg/response"
)

func (d *OrderDelivery) ReturnOrder(w http.ResponseWriter, r *http.Request) {
	rBody, err := io.ReadAll(r.Body)
	if err != nil {
		d.logger.Errorf("error in reading request body: %v", err)
		response.WriteResponse(w, []byte(`{"error":"internal error"}`), http.StatusInternalServerError, d.logger)
		return
	}

	defer func() {
		err = r.Body.Close()
		if err != nil {
			d.logger.Errorf("error in closing request body")
		}
	}()

	var returnOrderData dto.ReturnOrderInputData
	err = json.Unmarshal(rBody, &returnOrderData)
	if err != nil {
		var jsonErr *json.SyntaxError
		if errors.As(err, &jsonErr) {
			d.logger.Errorf("invalid json: %s", string(rBody))
			response.WriteResponse(w, []byte(`{"error":"invalid json passed"}`), http.StatusBadRequest, d.logger)
			return
		}
		d.logger.Errorf("error in response body unmarshalling: %v", err)
		response.WriteResponse(w, []byte(`{"error":"internal error"}`), http.StatusInternalServerError, d.logger)
		return
	}

	err = returnOrderData.Validate()
	if err != nil {
		d.logger.Errorf("validation errors in return order input data: %v", err)
		errText := fmt.Sprintf(`{"error":"%s"}`, err)
		response.WriteResponse(w, []byte(errText), http.StatusBadRequest, d.logger)
		return
	}

	err = d.service.ReturnOrder(r.Context(), returnOrderData.ClientID, returnOrderData.OrderID)
	errText := fmt.Sprintf(`{"error":"%s"}`, err)
	if errors.Is(err, service.ErrClientOrderNotFound) || errors.Is(err, storage.ErrClientOrderNotFound) {
		d.logger.Errorf("client with id %d has not got order with id %d", returnOrderData.ClientID, returnOrderData.OrderID)
		response.WriteResponse(w, []byte(errText), http.StatusBadRequest, d.logger)
		return
	}
	if errors.Is(err, service.ErrOrderIsNotIssued) {
		d.logger.Errorf("order with id %d is not issued", returnOrderData.OrderID)
		response.WriteResponse(w, []byte(errText), http.StatusBadRequest, d.logger)
		return
	}
	if errors.Is(err, service.ErrOrderIsReturned) {
		d.logger.Errorf("order with id %d is already returned", returnOrderData.OrderID)
		response.WriteResponse(w, []byte(errText), http.StatusBadRequest, d.logger)
		return
	}
	if errors.Is(err, service.ErrReturnTimeExpired) {
		d.logger.Errorf("return time for order %d has expired", returnOrderData.OrderID)
		response.WriteResponse(w, []byte(errText), http.StatusBadRequest, d.logger)
		return
	}
	if err != nil {
		d.logger.Errorf("internal server error in returning order: %v", err)
		response.WriteResponse(w, []byte(`{"error":"internal error"}`), http.StatusInternalServerError, d.logger)
		return
	}

	response.WriteResponse(w, []byte(`{"result":"success"}`), http.StatusOK, d.logger)
}
