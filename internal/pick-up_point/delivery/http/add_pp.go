package delivery

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"homework/internal/pick-up_point/delivery/dto"
	"homework/internal/pick-up_point/service"
	"homework/pkg/response"
)

func (d *PPDelivery) AddPickUpPoint(w http.ResponseWriter, r *http.Request) {
	rBody, err := io.ReadAll(r.Body)
	if err != nil {
		d.logger.Errorf("error in reading request body: %v", err)
		response.WriteResponse(w, response.Result{Res: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	defer func() {
		err = r.Body.Close()
		if err != nil {
			d.logger.Errorf("error in closing request body")
		}
	}()

	var pickUpPointDTO dto.PickUpPointAdd
	err = json.Unmarshal(rBody, &pickUpPointDTO)
	if err != nil {
		var jsonErr *json.SyntaxError
		if errors.As(err, &jsonErr) {
			d.logger.Errorf("invalid json: %s", string(rBody))
			response.WriteResponse(w, response.Result{Res: response.ErrInvalidJSON.Error()}, http.StatusBadRequest, d.logger)
			return
		}
		d.logger.Errorf("error in response body unmarshalling: %v", err)
		response.WriteResponse(w, response.Result{Res: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	err = pickUpPointDTO.Validate()
	if err != nil {
		d.logger.Errorf("validation errors in adding pick-up point: %v", err)
		response.WriteResponse(w, response.Result{Res: err.Error()}, http.StatusBadRequest, d.logger)
		return
	}

	pickUpPointToAdd := dto.ConvertToPickUpPoint(pickUpPointDTO)
	addedPickUpPoint, err := d.service.AddPickUpPoint(r.Context(), pickUpPointToAdd)
	if err != nil {
		if errors.Is(err, service.ErrPickUpPointAlreadyExists) {
			d.logger.Errorf("pick-up point with name %s already exists", pickUpPointDTO.Name)
			response.WriteResponse(w, response.Result{Res: err.Error()}, http.StatusBadRequest, d.logger)
			return
		}
		d.logger.Errorf("internal server error in adding pick-up point: %v", err)
		response.WriteResponse(w, response.Result{Res: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	response.WriteResponse(w, addedPickUpPoint, http.StatusOK, d.logger)
}
