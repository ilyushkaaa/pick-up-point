package delivery

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"homework/internal/pick-up_point/delivery/dto"
	"homework/internal/pick-up_point/service"
	"homework/internal/pick-up_point/storage"
	"homework/pkg/response"
)

func (d *PPDelivery) UpdatePickUpPoint(w http.ResponseWriter, r *http.Request) {
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

	var pickUpPointDTO dto.PickUpPointUpdate
	err = json.Unmarshal(rBody, &pickUpPointDTO)
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

	err = pickUpPointDTO.Validate()
	if err != nil {
		d.logger.Errorf("validation errors in adding pick-up point: %v", err)
		errText := fmt.Sprintf(`{"error":"%s"}`, err)
		response.WriteResponse(w, []byte(errText), http.StatusBadRequest, d.logger)
		return
	}

	pickUpPointToUpdate := pickUpPointDTO.ConvertToPickUpPoint()
	err = d.service.UpdatePickUpPoint(r.Context(), pickUpPointToUpdate)
	if errors.Is(err, storage.ErrPickUpPointNotFound) {
		d.logger.Errorf("pick-up point with id %d does not exist", pickUpPointDTO.ID)
		errText := fmt.Sprintf(`{"error":"no pick-up points with id %d"}`, pickUpPointDTO.ID)
		response.WriteResponse(w, []byte(errText), http.StatusNotFound, d.logger)
		return
	}
	if errors.Is(err, service.ErrPickUpPointAlreadyExists) {
		d.logger.Errorf("pick-up point with name %s already exists", pickUpPointDTO.Name)
		errText := fmt.Sprintf(`{"error":"pick-up point with name %s already exists"}`, pickUpPointDTO.Name)
		response.WriteResponse(w, []byte(errText), http.StatusBadRequest, d.logger)
		return
	}
	if err != nil {
		d.logger.Errorf("internal server error in updating pick-up point: %v", err)
		response.WriteResponse(w, []byte(`{"error":"internal error"}`), http.StatusInternalServerError, d.logger)
		return
	}

	ppJSON, err := json.Marshal(pickUpPointToUpdate)
	response.WriteMarshalledResponse(w, ppJSON, err, d.logger)

}
