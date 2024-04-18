package service

import (
	"context"
	"fmt"
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
	if err != nil {
		ps.cache.GoDeleteFromCache(context.Background(), fmt.Sprintf("pp_%d", id))
	}
	return err

}
