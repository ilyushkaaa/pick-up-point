package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"homework/internal/middleware"
	deliveryOrder "homework/internal/order/delivery/http"
	deliveryPP "homework/internal/pick-up_point/delivery/http"
)

func GetRouter(handlersPP *deliveryPP.PPDelivery, handlersOrder *deliveryOrder.OrderDelivery, mw *middleware.Middleware) *mux.Router {
	router := mux.NewRouter()
	assignRoutes(router, handlersPP, handlersOrder)
	assignMiddleware(router, mw)
	return router
}

func assignRoutes(router *mux.Router, handlersPP *deliveryPP.PPDelivery, handlersOrder *deliveryOrder.OrderDelivery) {
	router.HandleFunc("/pick-up-points", handlersPP.GetPickUpPoints).Methods(http.MethodGet)
	router.HandleFunc("/pick-up-point/{PP_ID}", handlersPP.GetPickUpPointByID).Methods(http.MethodGet)
	router.HandleFunc("/pick-up-point/{PP_ID}", handlersPP.DeletePickUpPoint).Methods(http.MethodDelete)
	router.HandleFunc("/pick-up-point", handlersPP.AddPickUpPoint).Methods(http.MethodPost)
	router.HandleFunc("/pick-up-point", handlersPP.UpdatePickUpPoint).Methods(http.MethodPut)

	router.HandleFunc("/order", handlersOrder.AddOrder).Methods(http.MethodPost)
	router.HandleFunc("/order/{ORDER_ID}", handlersOrder.DeleteOrder).Methods(http.MethodDelete)
	router.HandleFunc("/orders/{CLIENT_ID}", handlersOrder.GetUserOrders).Methods(http.MethodGet)
	router.HandleFunc("/orders/returns/{ORDERS_PER_PAGE}", handlersOrder.GetOrderReturns).Methods(http.MethodGet)
	router.HandleFunc("/orders/issue", handlersOrder.IssueOrders).Methods(http.MethodPut)
	router.HandleFunc("/orders/return", handlersOrder.ReturnOrder).Methods(http.MethodPut)
}

func assignMiddleware(router *mux.Router, mw *middleware.Middleware) {
	router.Use(mw.AccessLog)
	router.Use(mw.Auth)
}
