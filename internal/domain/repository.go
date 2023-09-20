package domain

import "context"

type PartnerRepository interface {
	GetByID(ctx context.Context, ID string) (Partner, error)
	GetByDocument(ctx context.Context, document string) (Partner, error)
	Create(ctx context.Context, partner Partner) (Partner, error)
}
