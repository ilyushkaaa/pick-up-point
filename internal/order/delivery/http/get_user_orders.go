package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"homework/internal/filters/model"
	"homework/pkg/response"
)

func (d *OrderDelivery) GetUserOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientID, ok := vars["CLIENT_ID"]
	if !ok {
		d.logger.Errorf("client id was not passed")
		response.WriteResponse(w, []byte(`{"error": "client id is not passed"}`), http.StatusBadRequest, d.logger)
		return
	}

	clientIDInt, err := strconv.ParseUint(clientID, 10, 64)
	if err != nil {
		d.logger.Errorf("error in client ID conversion: %s", err)
		errText := `{"error": "client ID must be positive integer"}`
		response.WriteResponse(w, []byte(errText), http.StatusBadRequest, d.logger)
		return
	}

	params := r.URL.Query()
	numOfLastOrders := params.Get("num_of_last_orders")
	numOfLastOrdersInt := 0
	if numOfLastOrders != "" {
		numOfLastOrdersInt, err = strconv.Atoi(numOfLastOrders)
		if err != nil || numOfLastOrdersInt < 1 {
			d.logger.Errorf("error in number of last orders conversion: %s", err)
			errText := `{"error": "number of last orders must positive integer"}`
			response.WriteResponse(w, []byte(errText), http.StatusBadRequest, d.logger)
			return
		}
	}
	ppOnly := params.Get("pp-only")
	ppOnlyBool := false
	if ppOnly != "" {
		ppOnlyBool, err = strconv.ParseBool(ppOnly)
		if err != nil {
			d.logger.Errorf("error in pp-only flag conversion: %s", err)
			errText := `{"error": "pp-only must be false or true"}`
			response.WriteResponse(w, []byte(errText), http.StatusBadRequest, d.logger)
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
		response.WriteResponse(w, []byte(`{"error":"internal error"}`), http.StatusInternalServerError, d.logger)
		return
	}

	ordersJSON, err := json.Marshal(orders)
	response.WriteMarshalledResponse(w, ordersJSON, err, d.logger)
}
