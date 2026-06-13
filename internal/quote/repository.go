package quote

import "context"

type Repository interface {
	Random(ctx context.Context) (Quote, error)
}
