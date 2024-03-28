package delivery

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"homework/internal/order/service"
	"homework/internal/order/storage"
	"homework/pkg/response"
)

func (d *OrderDelivery) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, ok := vars["ORDER_ID"]
	if !ok {
		d.logger.Errorf("order id was not passed")
		response.WriteResponse(w, []byte(`{"error": "order id is not passed"}`), http.StatusBadRequest, d.logger)
		return
	}

	orderIDInt, err := strconv.ParseUint(orderID, 10, 64)
	if err != nil {
		d.logger.Errorf("error in order ID conversion: %s", err)
		errText := `{"error": "order ID must be positive integer"}`
		response.WriteResponse(w, []byte(errText), http.StatusBadRequest, d.logger)
		return
	}

	err = d.service.DeleteOrder(r.Context(), orderIDInt)
	errText := fmt.Sprintf(`{"error":"%s"}`, err)
	if errors.Is(err, storage.ErrOrderNotFound) {
		d.logger.Errorf("no orders with id %d", orderIDInt)
		response.WriteResponse(w, []byte(errText), http.StatusNotFound, d.logger)
		return
	}
	if errors.Is(err, service.ErrOrderAlreadyIssued) {
		d.logger.Errorf("order with id %d is issued to user and is not in pick-up point", orderIDInt)
		response.WriteResponse(w, []byte(errText), http.StatusBadRequest, d.logger)
		return
	}
	if errors.Is(err, service.ErrOrderShelfLifeNotExpired) {
		d.logger.Errorf("shelf life for order %d is not expired", orderIDInt)
		response.WriteResponse(w, []byte(errText), http.StatusBadRequest, d.logger)
		return
	}
	if err != nil {
		d.logger.Errorf("internal server error in deleting order: %v", err)
		response.WriteResponse(w, []byte(`{"error":"internal error"}`), http.StatusInternalServerError, d.logger)
		return
	}

	response.WriteResponse(w, []byte(`{"result":"success"}`), http.StatusOK, d.logger)
}
