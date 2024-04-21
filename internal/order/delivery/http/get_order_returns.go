package delivery

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"homework/internal/order/delivery/dto"
	"homework/internal/order/service"
	"homework/pkg/response"
)

func (d *OrderDelivery) GetOrderReturns(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ordersPerPage, ok := vars["ORDERS_PER_PAGE"]
	if !ok {
		d.logger.Errorf("num of orders per page was not passed")
		response.MarshallAndWriteResponse(w, response.Error{Err: "num of orders per page is not passed"},
			http.StatusBadRequest, d.logger)
		return
	}

	ordersPerPageInt, err := strconv.ParseUint(ordersPerPage, 10, 64)
	if err != nil || ordersPerPageInt < 1 {
		d.logger.Errorf("error in num of orders per page conversion: %s", err)
		response.MarshallAndWriteResponse(w, response.Error{Err: "num of orders per page must be positive integer"},
			http.StatusBadRequest, d.logger)
		return
	}

	params := r.URL.Query()
	pageNum := params.Get("page_num")
	pageNumInt := uint64(1)
	if pageNum != "" {
		pageNumInt, err = strconv.ParseUint(pageNum, 10, 64)
		if err != nil || pageNumInt < 1 {
			d.logger.Errorf("error in page number conversion: %s", err)
			response.MarshallAndWriteResponse(w, response.Error{Err: "page number must positive integer"}, http.StatusBadRequest, d.logger)
			return
		}
	}

	orders, err := d.service.GetOrderReturns(r.Context(), ordersPerPageInt, pageNumInt)
	if err != nil {
		if errors.Is(err, service.ErrNoOrdersOnThisPage) {
			d.logger.Errorf("no orders on page %d when %d orders per page", pageNumInt, ordersPerPageInt)
			response.MarshallAndWriteResponse(w, response.Error{Err: err.Error()}, http.StatusBadRequest, d.logger)
			return
		}
		d.logger.Errorf("internal server error in getting order returns: %v", err)
		response.MarshallAndWriteResponse(w, response.Error{Err: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	ordersOutput := make([]dto.OrderOutput, 0, len(orders))
	for _, order := range orders {
		ordersOutput = append(ordersOutput, dto.NewOrderOutput(order))
	}
	response.MarshallAndWriteResponse(w, ordersOutput, http.StatusOK, d.logger)

}
