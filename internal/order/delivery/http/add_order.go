package delivery

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"homework/internal/order/delivery/dto"
	"homework/internal/order/service"
	"homework/internal/order/service/packages"
	"homework/pkg/response"
)

func (d *OrderDelivery) AddOrder(w http.ResponseWriter, r *http.Request) {
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

	var orderToAdd dto.OrderFromCourierInputData
	err = json.Unmarshal(rBody, &orderToAdd)
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

	err = orderToAdd.Validate()
	if err != nil {
		d.logger.Errorf("validation errors in adding order input data: %v", err)
		errText := fmt.Sprintf(`{"error":"%s"}`, err)
		response.WriteResponse(w, []byte(errText), http.StatusBadRequest, d.logger)
		return
	}

	order := orderToAdd.ConvertToOrder()
	err = d.service.AddOrder(r.Context(), order, orderToAdd.PackageType)
	errText := fmt.Sprintf(`{"error":"%s"}`, err)
	if errors.Is(err, service.ErrOrderAlreadyInPickUpPoint) {
		d.logger.Errorf("order with id %d already in pick-up point", order.ID)
		response.WriteResponse(w, []byte(errText), http.StatusBadRequest, d.logger)
		return
	}
	if errors.Is(err, service.ErrShelfTimeExpired) {
		d.logger.Errorf("shelf time for this order has expired: %v", order.StorageExpirationDate)
		response.WriteResponse(w, []byte(errText), http.StatusBadRequest, d.logger)
		return
	}
	if errors.Is(err, service.ErrUnknownPackage) {
		d.logger.Errorf("unknown package type %s", orderToAdd.PackageType)
		response.WriteResponse(w, []byte(errText), http.StatusBadRequest, d.logger)
		return
	}
	if errors.Is(err, packages.ErrPackageCanNotBeApplied) {
		d.logger.Errorf("%s can not be applied for order %v", orderToAdd.PackageType, order)
		response.WriteResponse(w, []byte(errText), http.StatusBadRequest, d.logger)
		return
	}
	if err != nil {
		d.logger.Errorf("internal server error in adding order: %v", err)
		response.WriteResponse(w, []byte(`{"error":"internal error"}`), http.StatusInternalServerError, d.logger)
		return
	}

	orderJSON, err := json.Marshal(order)
	response.WriteMarshalledResponse(w, orderJSON, err, d.logger)
}
