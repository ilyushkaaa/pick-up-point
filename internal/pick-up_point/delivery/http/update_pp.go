package delivery

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"homework/internal/pick-up_point/delivery/dto"
	"homework/internal/pick-up_point/storage"
	"homework/pkg/response"
)

func (d *PPDelivery) UpdatePickUpPoint(w http.ResponseWriter, r *http.Request) {
	rBody, err := io.ReadAll(r.Body)
	if err != nil {
		d.logger.Errorf("error in reading request body: %v", err)
		response.MarshallAndWriteResponse(w, response.Result{Res: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	defer func() {
		err = r.Body.Close()
		if err != nil {
			d.logger.Errorf("error in closing request body")
		}
	}()

	var pickUpPointDTO dto.PickUpPointUpdate
	err = json.Unmarshal(rBody, &pickUpPointDTO)
	if err != nil {
		var jsonErr *json.SyntaxError
		if errors.As(err, &jsonErr) {
			d.logger.Errorf("invalid json: %s", string(rBody))
			response.MarshallAndWriteResponse(w, response.Result{Res: response.ErrInvalidJSON.Error()}, http.StatusBadRequest, d.logger)
			return
		}
		d.logger.Errorf("error in response body unmarshalling: %v", err)
		response.MarshallAndWriteResponse(w, response.Result{Res: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	err = pickUpPointDTO.Validate()
	if err != nil {
		d.logger.Errorf("validation errors in adding pick-up point: %v", err)
		response.MarshallAndWriteResponse(w, response.Result{Res: err.Error()}, http.StatusBadRequest, d.logger)
		return
	}

	pickUpPointToUpdate := dto.ConvertPPUpdateToPickUpPoint(pickUpPointDTO)
	err = d.service.UpdatePickUpPoint(r.Context(), pickUpPointToUpdate)
	if err != nil {
		if errors.Is(err, storage.ErrPickUpPointNotFound) {
			d.logger.Errorf("pick-up point with id %d does not exist", pickUpPointDTO.ID)
			response.MarshallAndWriteResponse(w, response.Result{Res: "no pick-up points with such id"},
				http.StatusNotFound, d.logger)
			return
		}
		if errors.Is(err, storage.ErrPickUpPointNameExists) {
			d.logger.Errorf("pick-up point with name %s already exists", pickUpPointDTO.Name)
			response.MarshallAndWriteResponse(w, response.Result{Res: "pick-up point with such name already exists"},
				http.StatusBadRequest, d.logger)
			return
		}
		d.logger.Errorf("internal server error in updating pick-up point: %v", err)
		response.MarshallAndWriteResponse(w, response.Result{Res: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	d.cache.GoDeleteFromCache(context.Background(), fmt.Sprintf("pp_%d", pickUpPointToUpdate.ID))

	response.MarshallAndWriteResponse(w, pickUpPointToUpdate, http.StatusOK, d.logger)

}
