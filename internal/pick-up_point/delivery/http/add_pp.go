package delivery

import (
	"encoding/json"
	"errors"
	"fmt"
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
		err = response.WriteResponse(w, []byte(`{"error":"internal error"}`), http.StatusInternalServerError)
		if err != nil {
			d.logger.Errorf("error in writing response: %v", err)
		}
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
			err = response.WriteResponse(w, []byte(`{"error":"invalid json passed"}`), http.StatusBadRequest)
			if err != nil {
				d.logger.Errorf("error in writing response: %v", err)
			}
			return
		}
		d.logger.Errorf("error in response body unmarshalling: %v", err)
		err = response.WriteResponse(w, []byte(`{"error":"internal error"}`), http.StatusInternalServerError)
		if err != nil {
			d.logger.Errorf("error in writing response: %v", err)
		}
		return
	}

	err = pickUpPointDTO.Validate()
	if err != nil {
		d.logger.Errorf("validation errors in adding pick-up point: %v", err)
		errText := fmt.Sprintf(`{"error":"%s"}`, err)
		err = response.WriteResponse(w, []byte(errText), http.StatusBadRequest)
		if err != nil {
			d.logger.Errorf("error in writing response: %v", err)
		}
		return
	}

	pickUpPointToAdd := pickUpPointDTO.ConvertToPickUpPoint()
	addedPickUpPoint, err := d.service.AddPickUpPoint(r.Context(), pickUpPointToAdd)
	if errors.Is(err, service.ErrPickUpPointAlreadyExists) {
		d.logger.Errorf("pick-up point with name %s already exists", pickUpPointDTO.Name)
		errText := fmt.Sprintf(`{"error":"pick-up point with name %s already exists"}`, pickUpPointDTO.Name)
		err = response.WriteResponse(w, []byte(errText), http.StatusBadRequest)
		if err != nil {
			d.logger.Errorf("error in writing response: %v", err)
		}
		return
	}
	if err != nil {
		d.logger.Errorf("internal server error in adding pick-up point: %v", err)
		err = response.WriteResponse(w, []byte(`{"error":"internal error"}`), http.StatusInternalServerError)
		if err != nil {
			d.logger.Errorf("error in writing response: %v", err)
		}
		return
	}

	ppJSON, err := json.Marshal(addedPickUpPoint)
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
