package storage

import (
	"go.opentelemetry.io/otel/trace"
	"homework/pkg/infrastructure/database/postgres"
)

type PPStorageDB struct {
	db     database.Database
	tracer trace.Tracer
}

func New(db database.Database, tracer trace.Tracer) *PPStorageDB {
	return &PPStorageDB{
		db:     db,
		tracer: tracer,
	}
}
