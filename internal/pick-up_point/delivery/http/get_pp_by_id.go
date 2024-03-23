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
		response.WriteResponse(w, []byte(`{"error": "pick-up point id is not passed"}`), http.StatusBadRequest, d.logger)
		return
	}
	ppIDInt, err := strconv.ParseUint(ppID, 10, 64)
	if err != nil {
		d.logger.Errorf("error in pick-up point ID conversion: %s", err)
		errText := `{"error": "pick-up point ID must be positive integer"}`
		response.WriteResponse(w, []byte(errText), http.StatusBadRequest, d.logger)
		return
	}

	pickUpPoint, err := d.service.GetPickUpPointByID(r.Context(), ppIDInt)
	if errors.Is(err, service.ErrPickUpPointNotFound) {
		d.logger.Errorf("no pick-up points with id %d", ppIDInt)
		response.WriteResponse(w, []byte(`{"error": "pick-up with such id does not exist"}`), http.StatusNotFound, d.logger)
		return
	}
	if err != nil {
		d.logger.Errorf("internal server error in getting pick-up point: %v", err)
		response.WriteResponse(w, []byte(`{"error":"internal error"}`), http.StatusInternalServerError, d.logger)
		return
	}

	ppJSON, err := json.Marshal(pickUpPoint)
	response.WriteMarshalledResponse(w, ppJSON, err, d.logger)

}
