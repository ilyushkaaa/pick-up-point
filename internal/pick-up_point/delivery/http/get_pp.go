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
		response.WriteResponse(w, []byte(`{"error":"internal error"}`), http.StatusInternalServerError, d.logger)
		return
	}

	ppJSON, err := json.Marshal(pickUpPoints)
	response.WriteMarshalledResponse(w, ppJSON, err, d.logger)

}
