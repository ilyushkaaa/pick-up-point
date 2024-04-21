package service

import (
	"context"
	"strconv"
)

func (ps *PPService) DeletePickUpPoint(ctx context.Context, id uint64) error {
	err := ps.transactionManager.RunReadCommitted(ctx,
		func(ctxTX context.Context) error {
			err := ps.ppStorage.DeletePickUpPoint(ctx, id)
			if err != nil {
				return err
			}
			return ps.orderStorage.DeleteOrdersByPPID(ctx, id)
		})
	if err == nil {
		ps.cache.GoDeleteFromCache(context.Background(), strconv.FormatUint(id, 10))
	}
	return err

}
