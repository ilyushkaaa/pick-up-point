package prometheus

import "github.com/prometheus/client_golang/prometheus"

type BusinessMetrics struct {
	DeliveredOrdersCount      prometheus.Counter
	DeletedOOrders            prometheus.Counter
	IssuedOrdersCount         prometheus.Counter
	ReturnedOrdersCount       prometheus.Counter
	OrdersCountByPickUpPoints *prometheus.CounterVec
	OrdersByClientCount       *prometheus.CounterVec
}

func newBusinessMetrics() *BusinessMetrics {
	return &BusinessMetrics{
		DeliveredOrdersCount:      createCounter("delivered_orders_count", "num of delivered orders in pick-up-points"),
		DeletedOOrders:            createCounter("deleted_orders_count", "num of deleted orders from pick-up point (returned to courier)"),
		IssuedOrdersCount:         createCounter("issued_orders_count", "num of orders, which were issued by clients"),
		ReturnedOrdersCount:       createCounter("returned_orders_count", "num of orders, which were returned by clients to pick-up point"),
		OrdersCountByPickUpPoints: createCounterVec("orders_count_by_pick_up_point", "num of orders, delivered in each pick-up point", "pick_up_point_id"),
		OrdersByClientCount:       createCounterVec("orders_by_client_count", "num of orders, delivered for each client", "client_id"),
	}
}
