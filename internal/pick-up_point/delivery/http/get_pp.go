package delivery

import (
	"net/http"

	"homework/pkg/response"
)

func (d *PPDelivery) GetPickUpPoints(w http.ResponseWriter, r *http.Request) {
	pickUpPoints, err := d.service.GetPickUpPoints(r.Context())
	if err != nil {
		d.logger.Errorf("internal server error in getting pick-up points: %v", err)
		response.WriteResponse(w, response.Result{Res: response.ErrInternal.Error()}, http.StatusInternalServerError, d.logger)
		return
	}

	response.WriteResponse(w, pickUpPoints, http.StatusOK, d.logger)

}
