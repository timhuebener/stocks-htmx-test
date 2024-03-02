package psql

import (
	"context"
	"htmx/pkg/db/psql"
	"htmx/pkg/otel"
)

var tracer = otel.NewTracer("psql")

// Create - Create a new entry
func Create[T interface{}](ctx context.Context, value T) (*T, error) {
	_, span := tracer.Start(ctx, "POST")
	defer span.End()
	return psql.Create(value)
}

// GetByID - Get an entry by ID
func GetByID[T interface{}, C interface{}](ctx context.Context, dest T, conds ...C) (*T, error) {
	_, span := tracer.Start(ctx, "GET")
	defer span.End()
	return psql.GetByID(dest, conds)
}

// Get - Get all entries
func Get[T interface{}](ctx context.Context, dest []T) ([]T, error) {
	_, span := tracer.Start(ctx, "GET")
	defer span.End()
	return psql.Get(dest)
}

// Update - Update an existing entry
func Update[T interface{}](ctx context.Context, value T) (*T, error) {
	_, span := tracer.Start(ctx, "PUT")
	defer span.End()
	return psql.Update(value)
}

// Delete - Delete an entry
func Delete[T interface{}, C interface{}](ctx context.Context, value T, conds ...C) error {
	_, span := tracer.Start(ctx, "DELETE")
	defer span.End()
	return psql.Delete(value, conds)
}
