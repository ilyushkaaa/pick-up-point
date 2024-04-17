package service

import "context"

func (ps *PPService) DeletePickUpPoint(ctx context.Context, id uint64) error {
	return ps.transactionManager.RunReadCommitted(ctx,
		func(ctxTX context.Context) error {
			err := ps.ppStorage.DeletePickUpPoint(ctx, id)
			if err != nil {
				return err
			}
			return ps.orderStorage.DeleteOrdersByPPID(ctx, id)
		})

}
