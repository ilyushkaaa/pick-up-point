package storage

import (
	"go.opentelemetry.io/otel/trace"
	"homework/pkg/infrastructure/database/postgres"
)

type OrderStoragePG struct {
	db     database.Database
	tracer trace.Tracer
}

func New(db database.Database, tracer trace.Tracer) *OrderStoragePG {
	return &OrderStoragePG{db: db, tracer: tracer}
}
