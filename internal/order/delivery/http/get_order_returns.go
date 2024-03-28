package delivery

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"homework/internal/order/service"
	"homework/pkg/response"
)

func (d *OrderDelivery) GetOrderReturns(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ordersPerPage, ok := vars["ORDERS_PER_PAGE"]
	if !ok {
		d.logger.Errorf("num of orders per page was not passed")
		response.WriteResponse(w, []byte(`{"error": "num of orders per page is not passed"}`), http.StatusBadRequest, d.logger)
		return
	}

	ordersPerPageInt, err := strconv.ParseUint(ordersPerPage, 10, 64)
	if err != nil || ordersPerPageInt < 1 {
		d.logger.Errorf("error in num of orders per page conversion: %s", err)
		errText := `{"error": "num of orders per page must be positive integer"}`
		response.WriteResponse(w, []byte(errText), http.StatusBadRequest, d.logger)
		return
	}

	params := r.URL.Query()
	pageNum := params.Get("page_num")
	pageNumInt := uint64(1)
	if pageNum != "" {
		pageNumInt, err = strconv.ParseUint(pageNum, 10, 64)
		if err != nil || pageNumInt < 1 {
			d.logger.Errorf("error in page number conversion: %s", err)
			errText := `{"error": "page number must positive integer"}`
			response.WriteResponse(w, []byte(errText), http.StatusBadRequest, d.logger)
			return
		}
	}

	orders, err := d.service.GetOrderReturns(r.Context(), ordersPerPageInt, pageNumInt)
	if errors.Is(err, service.ErrNoOrdersOnThisPage) {
		d.logger.Errorf("no orders on page %d when %d orders per page", pageNumInt, ordersPerPageInt)
		errText := fmt.Sprintf(`{"error":"no orders on page %d when %d orders per page"}`, pageNumInt, ordersPerPageInt)
		response.WriteResponse(w, []byte(errText), http.StatusBadRequest, d.logger)
		return
	}

	if err != nil {
		d.logger.Errorf("internal server error in getting order returns: %v", err)
		response.WriteResponse(w, []byte(`{"error":"internal error"}`), http.StatusInternalServerError, d.logger)
		return
	}

	ordersJSON, err := json.Marshal(orders)
	response.WriteMarshalledResponse(w, ordersJSON, err, d.logger)

}
