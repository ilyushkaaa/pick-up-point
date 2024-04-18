package delivery

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"homework/internal/filters/model"
	"homework/internal/order/delivery/dto"
	"homework/pkg/response"
)

func (d *OrderDelivery) GetUserOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientID, ok := vars["CLIENT_ID"]
	if !ok {
		d.logger.Errorf("client id was not passed")
		response.MarshallAndWriteResponse(w, response.Result{Res: "client id was not passed"},
			http.StatusBadRequest, d.logger)
		return
	}

	clientIDInt, err := strconv.ParseUint(clientID, 10, 64)
	if err != nil {
		d.logger.Errorf("error in client ID conversion: %s", err)
		response.MarshallAndWriteResponse(w, response.Result{Res: "client ID must be positive integer"},
			http.StatusBadRequest, d.logger)
		return
	}

	params := r.URL.Query()
	numOfLastOrders := params.Get("num_of_last_orders")
	numOfLastOrdersInt := 0
	if numOfLastOrders != "" {
		numOfLastOrdersInt, err = strconv.Atoi(numOfLastOrders)
		if err != nil || numOfLastOrdersInt < 1 {
			d.logger.Errorf("error in number of last orders conversion: %s", err)
			response.MarshallAndWriteResponse(w, response.Result{Res: "number of last orders must positive integer"},
				http.StatusBadRequest, d.logger)
			return
		}
	}
	ppOnly := params.Get("pp-only")
	ppOnlyBool := false
	if ppOnly != "" {
		ppOnlyBool, err = strconv.ParseBool(ppOnly)
		if err != nil {
			d.logger.Errorf("error in pp-only flag conversion: %s", err)
			response.MarshallAndWriteResponse(w, response.Result{Res: "pp-only must be false or true"},
				http.StatusBadRequest, d.logger)
			return
		}
	}

	filters := model.Filters{
		Limit:  numOfLastOrdersInt,
		PPOnly: ppOnlyBool,
	}

	orders, err := d.service.GetUserOrders(r.Context(), clientIDInt, filters)
	if err != nil {
		d.logger.Errorf("internal server error in getting user orders: %v", err)
		response.MarshallAndWriteResponse(w, response.Result{Res: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	ordersOutput := make([]dto.OrderOutput, 0, len(orders))
	for _, order := range orders {
		ordersOutput = append(ordersOutput, dto.NewOrderOutput(order))
	}
	response.MarshallAndWriteResponse(w, ordersOutput, http.StatusOK, d.logger)
}
