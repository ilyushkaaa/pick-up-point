package delivery

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"homework/internal/order/delivery/dto"
	"homework/internal/order/service"
	"homework/pkg/response"
)

func (d *OrderDelivery) IssueOrders(w http.ResponseWriter, r *http.Request) {
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

	var ordersToIssue dto.OrdersToIssueInputData
	err = json.Unmarshal(rBody, &ordersToIssue)
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

	err = ordersToIssue.Validate()
	if err != nil {
		d.logger.Errorf("validation errors in issue orders input data: %v", err)
		response.MarshallAndWriteResponse(w, response.Result{Res: err.Error()}, http.StatusBadRequest, d.logger)
		return
	}

	orderIDs := make(map[uint64]struct{}, len(ordersToIssue.OrdersIDs))
	for _, id := range ordersToIssue.OrdersIDs {
		orderIDs[id] = struct{}{}
	}
	err = d.service.IssueOrders(r.Context(), orderIDs)
	if err != nil {
		if errors.Is(err, service.ErrOrdersOfDifferentClients) {
			d.logger.Error("passed orders belong to different clients")
			response.MarshallAndWriteResponse(w, response.Result{Res: err.Error()}, http.StatusBadRequest, d.logger)
			return
		}
		if errors.Is(err, service.ErrOrderAlreadyIssued) {
			d.logger.Error("there are orders which already issued")
			response.MarshallAndWriteResponse(w, response.Result{Res: err.Error()}, http.StatusBadRequest, d.logger)
			return
		}
		if errors.Is(err, service.ErrNotAllOrdersWereFound) {
			d.logger.Error("some of passed orders does not exist")
			response.MarshallAndWriteResponse(w, response.Result{Res: err.Error()}, http.StatusBadRequest, d.logger)
			return
		}
		d.logger.Errorf("internal server error in issuing orders: %v", err)
		response.MarshallAndWriteResponse(w, response.Result{Res: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	response.MarshallAndWriteResponse(w, response.Result{Res: "success"}, http.StatusOK, d.logger)
}
