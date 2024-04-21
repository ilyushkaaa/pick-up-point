package delivery

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"homework/internal/pick-up_point/storage"
	"homework/pkg/response"
)

func (d *PPDelivery) DeletePickUpPoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ppID, ok := vars["PP_ID"]
	if !ok {
		d.logger.Errorf("pick-up point id was not passed")
		response.MarshallAndWriteResponse(w, response.Error{Err: "pick-up point id was not passed"},
			http.StatusBadRequest, d.logger)
		return
	}

	ppIDInt, err := strconv.ParseUint(ppID, 10, 64)
	if err != nil {
		d.logger.Errorf("error in pick-up point ID conversion: %s", err)
		response.MarshallAndWriteResponse(w, response.Error{Err: "pick-up point ID must be positive integer"},
			http.StatusBadRequest, d.logger)
		return
	}

	err = d.service.DeletePickUpPoint(r.Context(), ppIDInt)
	if err != nil {
		if errors.Is(err, storage.ErrPickUpPointNotFound) {
			d.logger.Errorf("no pick-up points with id %d", ppIDInt)
			response.MarshallAndWriteResponse(w, response.Error{Err: err.Error()}, http.StatusNotFound, d.logger)
			return
		}
		d.logger.Errorf("internal server error in deleting pick-up point: %v", err)
		response.MarshallAndWriteResponse(w, response.Error{Err: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	d.cache.GoDeleteFromCache(context.Background(), fmt.Sprintf("pp_%s", ppID))

	response.MarshallAndWriteResponse(w, response.Result{Res: "success"}, http.StatusOK, d.logger)
}
