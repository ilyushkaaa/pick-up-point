package delivery

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"homework/internal/pick-up_point/storage"
	"homework/pkg/response"
)

func (d *PPDelivery) GetPickUpPointByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ppID, ok := vars["PP_ID"]
	if !ok {
		response.WriteResponse(w, response.Result{Res: "pick-up point id was not passed"},
			http.StatusBadRequest, d.logger)
		return
	}
	ppIDInt, err := strconv.ParseUint(ppID, 10, 64)
	if err != nil {
		d.logger.Errorf("error in pick-up point ID conversion: %s", err)
		response.WriteResponse(w, response.Result{Res: "pick-up point ID must be positive integer"},
			http.StatusBadRequest, d.logger)
		return
	}

	pickUpPoint, err := d.service.GetPickUpPointByID(r.Context(), ppIDInt)
	if err != nil {
		if errors.Is(err, storage.ErrPickUpPointNotFound) {
			d.logger.Errorf("no pick-up points with id %d", ppIDInt)
			response.WriteResponse(w, response.Result{Res: "pick-up with such id does not exist"},
				http.StatusNotFound, d.logger)
			return
		}
		d.logger.Errorf("internal server error in getting pick-up point: %v", err)
		response.WriteResponse(w, response.Result{Res: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	response.WriteResponse(w, pickUpPoint, http.StatusOK, d.logger)

}
