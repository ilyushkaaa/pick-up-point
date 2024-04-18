package delivery

import (
	"errors"
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
		response.MarshallAndWriteResponse(w, response.Result{Res: "order id is not passed"},
			http.StatusBadRequest, d.logger)
		return
	}

	orderIDInt, err := strconv.ParseUint(orderID, 10, 64)
	if err != nil {
		d.logger.Errorf("error in order ID conversion: %s", err)
		response.MarshallAndWriteResponse(w, response.Result{Res: "order ID must be positive integer"},
			http.StatusBadRequest, d.logger)
		return
	}

	err = d.service.DeleteOrder(r.Context(), orderIDInt)
	if errors.Is(err, storage.ErrOrderNotFound) {
		d.logger.Errorf("no orders with id %d", orderIDInt)
		response.MarshallAndWriteResponse(w, response.Result{Res: err.Error()}, http.StatusNotFound, d.logger)
		return
	}
	if errors.Is(err, service.ErrOrderAlreadyIssued) {
		d.logger.Errorf("order with id %d is issued to user and is not in pick-up point", orderIDInt)
		response.MarshallAndWriteResponse(w, response.Result{Res: err.Error()}, http.StatusBadRequest, d.logger)
		return
	}
	if errors.Is(err, service.ErrOrderShelfLifeNotExpired) {
		d.logger.Errorf("shelf life for order %d is not expired", orderIDInt)
		response.MarshallAndWriteResponse(w, response.Result{Res: err.Error()}, http.StatusBadRequest, d.logger)
		return
	}
	if err != nil {
		d.logger.Errorf("internal server error in deleting order: %v", err)
		response.MarshallAndWriteResponse(w, response.Result{Res: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	response.MarshallAndWriteResponse(w, response.Result{Res: "success"}, http.StatusOK, d.logger)
}
