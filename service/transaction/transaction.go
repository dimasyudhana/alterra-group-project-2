package transaction

import "context"

type TrxServiceInterface interface {
	Create(ctx context.Context, reqs []int, uid int) error
}
