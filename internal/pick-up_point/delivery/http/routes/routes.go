package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"homework/internal/middleware"
	delivery "homework/internal/pick-up_point/delivery/http"
)

func GetRouter(handlers *delivery.PPDelivery, mw *middleware.Middleware) *mux.Router {
	router := mux.NewRouter()
	assignRoutes(router, handlers)
	assignMiddleware(router, mw)
	return router
}

func assignRoutes(router *mux.Router, handlers *delivery.PPDelivery) {
	router.HandleFunc("/pick-up-points", handlers.GetPickUpPoints).Methods(http.MethodGet)
	router.HandleFunc("/pick-up-point/{PP_ID}", handlers.GetPickUpPointByID).Methods(http.MethodGet)
	router.HandleFunc("/pick-up-point/{PP_ID}", handlers.DeletePickUpPoint).Methods(http.MethodDelete)
	router.HandleFunc("/pick-up-point", handlers.AddPickUpPoint).Methods(http.MethodPost)
	router.HandleFunc("/pick-up-point", handlers.UpdatePickUpPoint).Methods(http.MethodPut)
}

func assignMiddleware(router *mux.Router, mw *middleware.Middleware) {
	router.Use(mw.AccessLog)
	router.Use(mw.Auth)
}
