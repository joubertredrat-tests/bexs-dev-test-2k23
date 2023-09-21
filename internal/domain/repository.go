package domain

import "context"

type Pagination struct {
	Offset uint
	Limit  uint
}

func NewPagination(offset, limit uint) Pagination {
	return Pagination{
		Offset: offset,
		Limit:  limit,
	}
}

type PartnerRepository interface {
	GetByID(ctx context.Context, ID string) (Partner, error)
	GetByDocument(ctx context.Context, document string) (Partner, error)
	Create(ctx context.Context, partner Partner) (Partner, error)
}

type PaymentRepository interface {
	GetByID(ctx context.Context, ID string) (Payment, error)
	List(ctx context.Context, paginaton Pagination) ([]Payment, error)
	Create(ctx context.Context, payment Payment) (Payment, error)
}
