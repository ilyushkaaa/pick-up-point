package delivery

import (
	"encoding/json"
	"net/http"

	"homework/pkg/response"
)

func (d *PPDelivery) GetPickUpPoints(w http.ResponseWriter, r *http.Request) {
	pickUpPoints, err := d.service.GetPickUpPoints(r.Context())
	if err != nil {
		d.logger.Errorf("internal server error in getting pick-up points: %v", err)
		err = response.WriteResponse(w, []byte(`{"error":"internal error"}`), http.StatusInternalServerError)
		if err != nil {
			d.logger.Errorf("error in writing response: %v", err)
		}
		return
	}

	ppJSON, err := json.Marshal(pickUpPoints)
	if err != nil {
		d.logger.Errorf("error in marshalling pick-up points: %v", err)
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
