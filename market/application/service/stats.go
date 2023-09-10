package service

import (
	"context"
	"time"
)

type Stats interface {
	Profit(ctx context.Context, startTime, endTime time.Time) (int, error)
}
