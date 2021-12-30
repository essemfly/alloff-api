package internal

import (
	"context"
)

type Repository interface {
	GetById(ctx context.Context, id string)
	GetByKey(ctx context.Context, key string, value string)
	List(ctx context.Context)
	UpdateById(ctx context.Context, id string, items interface{})
	Create() error
}
