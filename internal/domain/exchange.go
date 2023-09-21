package domain

import "context"

type Exchange interface {
	Convert(ctx context.Context, amount Amount, currency Currency) (Amount, error)
}
