package delivery

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"homework/internal/pick-up_point/service"
	"homework/pkg/response"
)

func (d *PPDelivery) GetPickUpPointByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ppID, ok := vars["PP_ID"]
	if !ok {
		d.logger.Errorf("pick-up point id was not passed")
		err := response.WriteResponse(w, []byte(`{"error": "pick-up point id is not passed"}`), http.StatusBadRequest)
		if err != nil {
			d.logger.Errorf("error in writing response: %v", err)
		}
		return
	}
	ppIDInt, err := strconv.ParseUint(ppID, 10, 64)
	if err != nil {
		d.logger.Errorf("error in pick-up point ID conversion: %s", err)
		errText := `{"error": "pick-up point ID must be positive integer"}`
		err = response.WriteResponse(w, []byte(errText), http.StatusBadRequest)
		if err != nil {
			d.logger.Errorf("error in writing response: %s", err)
		}
		return
	}

	pickUpPoint, err := d.service.GetPickUpPointByID(r.Context(), ppIDInt)
	if errors.Is(err, service.ErrPickUpPointNotFound) {
		d.logger.Errorf("no pick-up points with id %d", ppIDInt)
		err = response.WriteResponse(w, []byte(`{"error": "pick-up with such id does not exist"}`), http.StatusNotFound)
		if err != nil {
			d.logger.Errorf("error in writing response: %v", err)
		}
		return
	}
	if err != nil {
		d.logger.Errorf("internal server error in getting pick-up point: %v", err)
		err = response.WriteResponse(w, []byte(`{"error":"internal error"}`), http.StatusInternalServerError)
		if err != nil {
			d.logger.Errorf("error in writing response: %v", err)
		}
		return
	}

	ppJSON, err := json.Marshal(pickUpPoint)
	if err != nil {
		d.logger.Errorf("error in marshalling pick-up point: %v", err)
		err = response.WriteResponse(w, []byte(`{"error":"internal error"}`), http.StatusInternalServerError)
		if err != nil {
			d.logger.Errorf("error in writing response: %v", err)
		}
		return
	}

	err = response.WriteResponse(w, ppJSON, http.StatusOK)
	if err != nil {
		d.logger.Errorf("error in writing response: %v", err)
	}
}
