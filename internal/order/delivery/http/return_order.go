package delivery

import (
	"encoding/json"
	"errors"
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
		response.MarshallAndWriteResponse(w, response.Result{Res: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
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
			response.MarshallAndWriteResponse(w, response.Result{Res: response.ErrInvalidJSON.Error()}, http.StatusBadRequest, d.logger)
			return
		}
		d.logger.Errorf("error in response body unmarshalling: %v", err)
		response.MarshallAndWriteResponse(w, response.Result{Res: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	err = returnOrderData.Validate()
	if err != nil {
		d.logger.Errorf("validation errors in return order input data: %v", err)
		response.MarshallAndWriteResponse(w, response.Result{Res: err.Error()}, http.StatusBadRequest, d.logger)
		return
	}

	err = d.service.ReturnOrder(r.Context(), returnOrderData.ClientID, returnOrderData.OrderID)
	if err != nil {
		if errors.Is(err, service.ErrClientOrderNotFound) || errors.Is(err, storage.ErrClientOrderNotFound) {
			d.logger.Errorf("client with id %d has not got order with id %d", returnOrderData.ClientID, returnOrderData.OrderID)
			response.MarshallAndWriteResponse(w, response.Result{Res: err.Error()}, http.StatusBadRequest, d.logger)
			return
		}
		if errors.Is(err, service.ErrOrderIsNotIssued) {
			d.logger.Errorf("order with id %d is not issued", returnOrderData.OrderID)
			response.MarshallAndWriteResponse(w, response.Result{Res: err.Error()}, http.StatusBadRequest, d.logger)
			return
		}
		if errors.Is(err, service.ErrOrderIsReturned) {
			d.logger.Errorf("order with id %d is already returned", returnOrderData.OrderID)
			response.MarshallAndWriteResponse(w, response.Result{Res: err.Error()}, http.StatusBadRequest, d.logger)
			return
		}
		if errors.Is(err, service.ErrReturnTimeExpired) {
			d.logger.Errorf("return time for order %d has expired", returnOrderData.OrderID)
			response.MarshallAndWriteResponse(w, response.Result{Res: err.Error()}, http.StatusBadRequest, d.logger)
			return
		}
		d.logger.Errorf("internal server error in returning order: %v", err)
		response.MarshallAndWriteResponse(w, response.Result{Res: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	response.MarshallAndWriteResponse(w, response.Result{Res: "success"}, http.StatusOK, d.logger)
}
